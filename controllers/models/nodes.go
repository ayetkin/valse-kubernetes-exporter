package models

type Nodes struct {
	Hostname   string        `json:"hostname" bson:"hostname"`
	Ip         string        `json:"ip" bson:"ip"`
	Role       string        `json:"role" bson:"role"`
	Age        string        `json:"age" bson:"age"`
	Version    string        `json:"version" bson:"version"`
	Conditions []*Conditions `json:"conditions" bson:"conditions"`
}

type Conditions struct {
	Type    string `json:"type" bson:"type"`
	Status  string `json:"status" bson:"status"`
	Reason  string `json:"reason" bson:"reason"`
	Message string `json:"message" bson:"message"`
}
