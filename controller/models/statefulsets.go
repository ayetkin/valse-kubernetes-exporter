package models

type StatefulSets struct {
	Name          string `json:"name" bson:"name"`
	Namespace     string `json:"namespace" bson:"namespace"`
	Age           string `json:"age" bson:"age"`
	Replicas      int32  `json:"replicas" bson:"replicas"`
	ReadyReplicas int32  `json:"ready_replicas" bson:"ready_replicas"`
}
