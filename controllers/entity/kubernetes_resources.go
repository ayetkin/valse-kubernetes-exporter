package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"valse/controllers/models"
)

type KubernetesResources struct {
	Id           primitive.ObjectID    `json:"id" bson:"_id"`
	LastUpdate   time.Time             `json:"last_update" bson:"last_update"`
	Cluster      models.Cluster        `json:"cluster" bson:"cluster"`
	Nodes        []models.Nodes        `json:"nodes" bson:"nodes"`
	Namespaces   []models.Namespaces   `json:"namespaces" bson:"namespaces"`
	Deployments  []models.Deployments  `json:"deployments" bson:"deployments"`
	DaemonSets   []models.DaemonSets   `json:"daemonsets" bson:"daemonsets"`
	StatefulSets []models.StatefulSets `json:"statefulsets" bson:"statefulsets"`
	Jobs         []models.Jobs         `json:"jobs" bson:"jobs"`
	CronJobs     []models.CronJobs     `json:"cronjobs" bson:"cronjobs"`
	Pods         []models.Pods         `json:"pods" bson:"pods"`
	Services     []models.Services     `json:"services" bson:"services"`
}
