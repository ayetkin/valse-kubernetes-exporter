package models

type Namespaces struct {
	Name  string `json:"name" bson:"name"`
	Phase string `json:"phase" bson:"phase"`
	Age   string `json:"age" bson:"age"`
}
