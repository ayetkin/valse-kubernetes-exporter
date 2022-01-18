package models

type Services struct {
	Name        string             `json:"name" bson:"name"`
	Namespace   string             `json:"namespace" bson:"namespace"`
	Annotations *map[string]string `json:"annotations" bson:"annotations"`
	Age         string             `json:"age" bson:"age"`
	Type        string             `json:"type" bson:"type"`
	Ports       []*ServicePorts    `json:"ports" bson:"ports"`
}
type ServicePorts struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Port     int32  `json:"port" bson:"port"`
	NodePort *int32 `json:"node_port" bson:"node_port"`
}
