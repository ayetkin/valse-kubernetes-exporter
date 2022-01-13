package models

type Jobs struct {
	Name            string              `json:"name" bson:"name"`
	Namespaces      string              `json:"namespaces" bson:"namespaces"`
	Age             string              `json:"age" bson:"age"`
	OwnerReferences *JobOwnerReferences `json:"owner_references" bson:"owner_references"`
	Conditions      []JobsConditions    `json:"conditions" bson:"conditions"`
}

type JobsConditions struct {
	Type    string `json:"type" bson:"type"`
	Status  string `json:"status" bson:"status"`
	Reason  string `json:"reason" bson:"reason"`
	Message string `json:"message" bson:"message"`
}

type JobOwnerReferences struct {
	Kind string `json:"kind" bson:"kind"`
	Name string `json:"name" bson:"name"`
}
