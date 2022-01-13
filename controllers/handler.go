package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"valse/controllers/entity"
	"valse/controllers/models"
	"valse/controllers/repository"
	"valse/pkg/config"
	"valse/pkg/k8s"
	"valse/pkg/mongodb"
)

func RunScheduledJob(appConfig *config.AppConfig, k8sClient k8s.Client, mongodbClient mongodb.OfficialClient, logger *logrus.Logger) {

	logger.Warningf("Kubernetes resources discovery task triggered (interval: %v)", appConfig.ScheduledTaskIntervalSeconds*time.Second)

	k8sResources, err := DiscoverKubernetesResources(appConfig, k8sClient, logger)
	if err != nil {
		logger.WithField("exception.backtrace", err).Errorf("An error occurred while discovering resources from kubernetes!")
		return
	}

	err = UpdateMongoCollection(appConfig, k8sResources, mongodbClient, logger)
	if err != nil {
		logger.WithField("exception.backtrace", err).Errorf("An error occurred while updating mongodb collection!")
		return
	}

	logger.Infof("Scheduled discovery task succesfully complated (next run: %v)",
		time.Now().UTC().Add(3*time.Hour+appConfig.ScheduledTaskIntervalSeconds*time.Second).Format("2006-01-02 15:04:05"))
}

func DiscoverKubernetesResources(appConfig *config.AppConfig, k8sClient k8s.Client, logger *logrus.Logger) (*entity.KubernetesResources, error) {

	var k8sGet = repository.NewGet("", appConfig.ExcludedNamespaces, k8sClient, logger)

	version, err := k8sGet.Version()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve Kubernetes version. %v", err))
	}

	nodes, err := k8sGet.Nodes()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve nodes. %v", err))
	}

	namespaces, err := k8sGet.Namespaces()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve namespaces. %v", err))
	}

	deployments, err := k8sGet.Deployments()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve deployments. %v", err))
	}

	daemonSets, err := k8sGet.DaemonSets()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve daemomsets. %v", err))
	}

	statefulSets, err := k8sGet.StatefulSets()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve statefulsets. %v", err))
	}

	jobs, err := k8sGet.Jobs()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve jobs. %v", err))
	}

	cronJobs, err := k8sGet.CronJobs()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve cronjobs. %v", err))
	}

	services, err := k8sGet.Services()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve services. %v", err))
	}

	pods, err := k8sGet.Pods()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to retrieve pods. %v", err))
	}

	cluster := models.Cluster{
		Name:     appConfig.Cluster.Name,
		Version:  version.String(),
		Region:   appConfig.Cluster.Region,
		Address:  nodes[0].Ip,
		Hostname: nodes[0].Hostname,
		Statics: &models.Statics{
			NodeCount:        len(nodes),
			NamespaceCount:   len(namespaces),
			DeploymentCount:  len(deployments),
			StatefulSetCount: len(statefulSets),
			DaemonSetCount:   len(daemonSets),
			JobCount:         len(jobs),
			CronJobCount:     len(cronJobs),
			PodCount:         len(pods),
			ServiceCount:     len(services),
		},
	}

	var k8sResources = &entity.KubernetesResources{
		Id:           primitive.NewObjectID(),
		LastUpdate:   time.Now().UTC().Add(3 * time.Hour),
		Cluster:      cluster,
		Nodes:        nodes,
		Namespaces:   namespaces,
		Deployments:  deployments,
		DaemonSets:   daemonSets,
		StatefulSets: statefulSets,
		Jobs:         jobs,
		CronJobs:     cronJobs,
		Pods:         pods,
		Services:     services,
	}

	return k8sResources, nil
}

func UpdateMongoCollection(appConfig *config.AppConfig, k8sResources *entity.KubernetesResources, client mongodb.OfficialClient, logger *logrus.Logger) error {

	valseRepository := repository.NewValseRepository(appConfig.MongoDB.Database, client, logger)

	itemFromMongodb, err := valseRepository.GetItemByClusterAddress(context.Background(), k8sResources.Cluster.Address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.WithField("exception.backtrace", err).Warningf(fmt.Sprintf("Failed to getting item from mongodb!"))
		} else {
			return errors.New(fmt.Sprintf("Failed to getting item from mongodb. %v", err))
		}
	} else {
		k8sResources.Id = itemFromMongodb.Id
	}

	updateResult, err := valseRepository.UpsertItem(context.Background(), *k8sResources)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to upserting item to mongodb. %v", err))
	}

	logger.Infof("Collection successfully updated. (Mached:%v Upserted:%v Modified:%v)",
		updateResult.MatchedCount, updateResult.UpsertedCount, updateResult.ModifiedCount)

	return nil
}
