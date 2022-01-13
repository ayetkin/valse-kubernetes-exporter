package models

type Cluster struct {
	Name     string   `json:"name" bson:"name"`
	Region   string   `json:"region" bson:"region"`
	Version  string   `json:"version" bson:"version"`
	Address  string   `json:"address" bson:"address"`
	Hostname string   `json:"hostname" bson:"hostname"`
	Statics  *Statics `json:"statics" bson:"statics"`
}

type Statics struct {
	NodeCount        int `json:"node_count" bson:"node_count"`
	NamespaceCount   int `json:"namespace_count" bson:"namespace_count"`
	DeploymentCount  int `json:"deployment_count" bson:"deployment_count"`
	StatefulSetCount int `json:"statefulset_count" bson:"statefulset_count"`
	DaemonSetCount   int `json:"daemonset_count" bson:"daemonset_count"`
	JobCount         int `json:"job_count" bson:"job_count"`
	CronJobCount     int `json:"cronjob_count" bson:"cronjob_count"`
	PodCount         int `json:"pod_count" bson:"pod_count"`
	ServiceCount     int `json:"service_count" bson:"service_count"`
}
