package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"valse/controllers/consts"
	"valse/controllers/entity"
	"valse/pkg/mongodb"
)

type ValseRepository interface {
	GetItemByClusterAddress(ctx context.Context, clusterAddress string) (*entity.KubernetesResources, error)
	UpsertItem(ctx context.Context, item entity.KubernetesResources) (*mongo.UpdateResult, error)
}

type valseRepository struct {
	client       mongodb.OfficialClient
	databaseName string
	collection   string
	logger       *logrus.Logger
}

func NewValseRepository(databaseName string, client mongodb.OfficialClient, logger *logrus.Logger) ValseRepository {

	err := client.EnsureIndex(
		[]string{"cluster.address"},
		true,
		"ClusterAddress",
		databaseName, consts.DefaultCollection,
	)
	if err != nil {
		logger.Fatal(err)
	}

	return &valseRepository{
		client:       client,
		databaseName: databaseName,
		collection:   consts.DefaultCollection,
		logger:       logger,
	}
}

func (v *valseRepository) GetItemByClusterAddress(ctx context.Context, clusterAddress string) (k8sResources *entity.KubernetesResources, err error) {

	var session = v.client.NewSessionWithSecondaryPreferred(v.databaseName)

	err = session.Collection(v.collection).FindOne(ctx, bson.M{"cluster.address": clusterAddress}).Decode(&k8sResources)
	if err != nil {
		return nil, err
	}

	return k8sResources, nil
}

func (v *valseRepository) UpsertItem(ctx context.Context, item entity.KubernetesResources) (*mongo.UpdateResult, error) {

	var session = v.client.NewSession(v.databaseName)

	opts := options.Replace().SetUpsert(true)
	filter := bson.M{"cluster.address": item.Cluster.Address}

	updateResult, err := session.Collection(v.collection).ReplaceOne(ctx, filter, item, opts)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
