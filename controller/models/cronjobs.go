package models

import "time"

type CronJobs struct {
	Name              string     `json:"name" bson:"name"`
	Namespace         string     `json:"namespace" bson:"namespace"`
	Age               string     `json:"age" bson:"age"`
	Schedule          string     `json:"schedule" bson:"schedule"`
	Suspended         *bool      `json:"suspended" bson:"suspended"`
	Active            []Active   `json:"active" bson:"active"`
	LastScheduledTime *time.Time `json:"last_scheduled_time" bson:"last_scheduled_time"`
}

type Active struct {
	Kind      string `json:"kind" bson:"kind"`
	Name      string `json:"name" bson:"name"`
	Namespace string `json:"namespace" bson:"namespace"`
}
