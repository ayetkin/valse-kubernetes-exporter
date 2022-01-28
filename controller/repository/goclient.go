package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	meta1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"strings"
	"valse/controller/models"
	"valse/pkg/k8s"
	"valse/pkg/utils"
)

type Get interface {
	Version() (*version.Info, error)
	CertExpireDate() (string, error)
	Nodes() ([]models.Nodes, error)
	Namespaces() ([]models.Namespaces, error)
	Deployments() ([]models.Deployments, error)
	DaemonSets() ([]models.DaemonSets, error)
	StatefulSets() ([]models.StatefulSets, error)
	CronJobs() ([]models.CronJobs, error)
	Jobs() ([]models.Jobs, error)
	Pods() ([]models.Pods, error)
	Services() ([]models.Services, error)
}

type get struct {
	namespace          string
	excludedNamespaces string
	client             k8s.Client
	logger             *logrus.Logger
}

func NewGet(namespace string, excludedNamespaces string, clientSet k8s.Client, logger *logrus.Logger) Get {
	return &get{
		namespace:          namespace,
		excludedNamespaces: excludedNamespaces,
		client:             clientSet,
		logger:             logger,
	}
}

func (g *get) Version() (*version.Info, error) {

	g.logger.Info("Getting resource from kubernetes api: Version")

	api := g.client.ApiDiscovery()

	ver, err := api.ServerVersion()
	if err != nil {
		return nil, err
	}

	return ver, nil
}

func (g *get) Nodes() ([]models.Nodes, error) {

	g.logger.Info("Getting resource from kubernetes api: Nodes")

	api := g.client.ApiCoreV1()

	nodeList, err := api.Nodes().List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nodes []models.Nodes

	for _, item := range nodeList.Items {

		var nodeCondition []*models.Conditions

		for _, condition := range item.Status.Conditions {
			c := models.Conditions{
				Type:    string(condition.Type),
				Status:  string(condition.Status),
				Reason:  condition.Reason,
				Message: condition.Message,
			}

			nodeCondition = append(nodeCondition, &c)
		}

		n := models.Nodes{
			Role:       "worker",
			Age:        utils.Age(item.CreationTimestamp.Time),
			Version:    item.Status.NodeInfo.KubeletVersion,
			Conditions: nodeCondition,
		}

		for _, address := range item.Status.Addresses {
			if address.Type == "Hostname" {
				n.Hostname = address.Address
			}
			if address.Type == "InternalIP" {
				n.Ip = address.Address
			}
		}

		for i, v := range item.Labels {
			if strings.Contains(i, "node-role.kubernetes.io/master") || strings.Contains(i, "node-role.kubernetes.io/control-plane") {
				if v == "" {
					n.Role = "master"
					break
				}
			}
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

func (g *get) Namespaces() ([]models.Namespaces, error) {

	g.logger.Info("Getting resource from kubernetes api: Namespaces")

	api := g.client.ApiCoreV1()

	namespaceList, err := api.Namespaces().List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaces []models.Namespaces

	for _, item := range namespaceList.Items {

		if strings.Contains(g.excludedNamespaces, item.Name) {
			continue
		}

		n := models.Namespaces{
			Name:  item.Name,
			Phase: string(item.Status.Phase),
			Age:   utils.Age(item.CreationTimestamp.Time),
		}

		namespaces = append(namespaces, n)
	}

	return namespaces, err
}

func (g *get) Deployments() ([]models.Deployments, error) {

	g.logger.Info("Getting resource from kubernetes api: Deployments")

	api := g.client.ApiAppsV1()

	deploymentList, err := api.Deployments(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var deployments []models.Deployments

	for _, item := range deploymentList.Items {

		if strings.Contains(g.excludedNamespaces, item.Namespace) {
			continue
		}

		d := models.Deployments{
			Name:                item.Name,
			Namespace:           item.Namespace,
			Age:                 utils.Age(item.CreationTimestamp.Time),
			Replicas:            item.Status.Replicas,
			ReadyReplicas:       item.Status.ReadyReplicas,
			AvailableReplicas:   item.Status.AvailableReplicas,
			UnavailableReplicas: item.Status.UnavailableReplicas,
		}

		deployments = append(deployments, d)
	}

	return deployments, nil
}

func (g *get) DaemonSets() ([]models.DaemonSets, error) {

	g.logger.Info("Getting resource from kubernetes api: DaemonSets")

	api := g.client.ApiAppsV1()

	daemonSetList, err := api.DaemonSets(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var daemonSets []models.DaemonSets

	for _, item := range daemonSetList.Items {

		if strings.Contains(g.excludedNamespaces, item.Namespace) {
			continue
		}

		d := models.DaemonSets{
			Name:                   item.Name,
			Namespace:              item.Namespace,
			Age:                    utils.Age(item.CreationTimestamp.Time),
			DesiredNumberScheduled: item.Status.DesiredNumberScheduled,
			CurrentNumberScheduled: item.Status.CurrentNumberScheduled,
			NumberReady:            item.Status.NumberReady,
			NumberAvailable:        item.Status.NumberAvailable,
		}

		daemonSets = append(daemonSets, d)
	}

	return daemonSets, err
}

func (g *get) StatefulSets() ([]models.StatefulSets, error) {

	g.logger.Info("Getting resource from kubernetes api: StatefulSets")

	api := g.client.ApiAppsV1()

	statefulSetList, err := api.StatefulSets(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var statefulSets []models.StatefulSets

	for _, item := range statefulSetList.Items {

		if strings.Contains(g.excludedNamespaces, item.Namespace) {
			continue
		}

		s := models.StatefulSets{
			Name:          item.Name,
			Namespace:     item.Namespace,
			Age:           utils.Age(item.CreationTimestamp.Time),
			Replicas:      item.Status.Replicas,
			ReadyReplicas: item.Status.ReadyReplicas,
		}

		statefulSets = append(statefulSets, s)
	}
	return statefulSets, err
}

func (g *get) Jobs() ([]models.Jobs, error) {

	g.logger.Info("Getting resource from kubernetes api: Jobs")

	api := g.client.ApiBatchV1()

	jobList, err := api.Jobs(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var jobs []models.Jobs

	for _, item := range jobList.Items {

		var jobsConditions []models.JobsConditions

		for _, condition := range item.Status.Conditions {
			c := models.JobsConditions{
				Type:    string(condition.Type),
				Status:  string(condition.Status),
				Reason:  condition.Reason,
				Message: condition.Message,
			}
			jobsConditions = append(jobsConditions, c)
		}

		j := models.Jobs{
			Name:            item.Name,
			Namespaces:      item.Namespace,
			Age:             utils.Age(item.CreationTimestamp.Time),
			OwnerReferences: &models.JobOwnerReferences{},
			Conditions:      jobsConditions,
		}

		if item.OwnerReferences != nil {
			j.OwnerReferences.Kind = item.OwnerReferences[0].Kind
			j.OwnerReferences.Name = item.OwnerReferences[0].Name
		} else {
			j.OwnerReferences = nil
		}

		jobs = append(jobs, j)
	}

	return jobs, nil

}

func (g *get) CronJobs() ([]models.CronJobs, error) {

	g.logger.Info("Getting resource from kubernetes api: CronJobs")

	api := g.client.ApiBatchV1beta1()

	cronJobsList, err := api.CronJobs(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var cronJobs []models.CronJobs

	for _, item := range cronJobsList.Items {

		var cronJobsActive []models.Active

		for _, v := range item.Status.Active {
			a := models.Active{
				Kind:      v.Kind,
				Name:      v.Name,
				Namespace: v.Namespace,
			}
			cronJobsActive = append(cronJobsActive, a)
		}

		c := models.CronJobs{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Age:               utils.Age(item.CreationTimestamp.Time),
			Schedule:          item.Spec.Schedule,
			Suspended:         item.Spec.Suspend,
			Active:            cronJobsActive,
			LastScheduledTime: nil,
		}

		if item.Status.LastScheduleTime != nil {
			c.LastScheduledTime = &item.Status.LastScheduleTime.Time
		}

		c.Active = cronJobsActive

		cronJobs = append(cronJobs, c)
	}

	return cronJobs, nil
}

func (g *get) Pods() ([]models.Pods, error) {

	g.logger.Info("Getting resource from kubernetes api: Pods")

	api := g.client.ApiCoreV1()

	PodsList, err := api.Pods(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var pods []models.Pods

	for _, item := range PodsList.Items {

		if strings.Contains(g.excludedNamespaces, item.Namespace) {
			continue
		}

		var containerStatuses []models.ContainerStatuses

		for _, v := range item.Status.ContainerStatuses {

			c := models.ContainerStatuses{
				Name:         v.Name,
				Ready:        v.Ready,
				Started:      v.Started,
				RestartCount: v.RestartCount,
			}

			if v.State.Waiting != nil {
				c.State.Waiting = &models.ContainerStatusesStateWaiting{
					Reason:  v.State.Waiting.Reason,
					Message: v.State.Waiting.Message,
				}
			}

			if v.State.Running != nil {
				c.State.Running = &models.ContainerStatusesStateRunning{
					StartedAt: v.State.Running.StartedAt.Time,
				}
			}

			if v.State.Terminated != nil {
				c.State.Terminated = &models.ContainerStatusesTerminated{
					ExitCode: v.State.Terminated.ExitCode,
					Signal:   v.State.Terminated.Signal,
					Reason:   v.State.Terminated.Reason,
					Message:  v.State.Terminated.Message,
				}
			}

			containerStatuses = append(containerStatuses, c)
		}

		p := models.Pods{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Age:               utils.Age(item.CreationTimestamp.Time),
			Phase:             string(item.Status.Phase),
			Reason:            item.Status.Reason,
			Message:           item.Status.Message,
			HostIP:            item.Status.HostIP,
			ContainerStatuses: containerStatuses,
		}

		if item.OwnerReferences != nil {
			p.OwnerReferences = &models.PodOwnerReferences{
				Kind: item.OwnerReferences[0].Kind,
				Name: item.OwnerReferences[0].Name,
			}

			if p.OwnerReferences.Kind == "ReplicaSet" {
				p.OwnerReferences.Kind = "Deployment"
				p.OwnerReferences.Name = utils.SplitKindName(p.OwnerReferences.Name)
			}

			if p.OwnerReferences.Kind == "Node" {
				p.OwnerReferences.Kind = "Pod"
				p.OwnerReferences.Name = item.Name
			}
		} else {
			p.OwnerReferences = nil
		}

		pods = append(pods, p)
	}
	return pods, err
}

func (g *get) Services() ([]models.Services, error) {

	g.logger.Info("Getting resource from kubernetes api: Services")

	api := g.client.ApiCoreV1()

	serviceList, err := api.Services(g.namespace).List(context.Background(), meta1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var services []models.Services

	for _, item := range serviceList.Items {

		if strings.Contains(g.excludedNamespaces, item.Namespace) {
			continue
		}

		var servicePorts []*models.ServicePorts

		for _, port := range item.Spec.Ports {

			p := models.ServicePorts{
				Name:     port.Name,
				Port:     port.Port,
				NodePort: &port.NodePort,
			}

			if port.NodePort == 0 {
				p.NodePort = nil
			}

			servicePorts = append(servicePorts, &p)
		}

		var annotations = &map[string]string{}

		if item.Annotations == nil {
			annotations = nil
		}

		for k, v := range item.Annotations {
			if !strings.Contains(k, "envoy") {
				annotations = nil
				break
			}
			(*annotations)[k] = v
		}

		s := models.Services{
			Name:        item.Name,
			Namespace:   item.Namespace,
			Annotations: annotations,
			Age:         utils.Age(item.CreationTimestamp.Time),
			Type:        string(item.Spec.Type),
			Ports:       servicePorts,
		}

		services = append(services, s)
	}

	return services, nil
}

func (g *get) CertExpireDate() (string, error) {

	g.logger.Info("Getting resource from kubernetes url: Certificate expire date")

	apiServer := g.client.ApiCoreV1().RESTClient().Get().URL().Host

	cnnState, err := utils.TlsDial(apiServer)
	if err != nil {
		return "", err
	}

	expiry := cnnState().PeerCertificates[0].NotAfter

	return expiry.Format("02.01.2006"), nil
}
