package models

type Deployments struct {
	Name                string `json:"name" bson:"name"`
	Namespace           string `json:"namespace" bson:"namespace"`
	Age                 string `json:"age" bson:"age"`
	Replicas            int32  `json:"replicas" bson:"replicas"`
	ReadyReplicas       int32  `json:"ready_replicas" bson:"ready_replicas"`
	AvailableReplicas   int32  `json:"available_replicas" bson:"available_replicas"`
	UnavailableReplicas int32  `json:"unavailable_replicas" bson:"unavailable_replicas"`
}
