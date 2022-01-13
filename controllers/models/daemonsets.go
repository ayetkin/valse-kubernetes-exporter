package models

type DaemonSets struct {
	Name                   string `json:"name" bson:"name"`
	Namespace              string `json:"namespace" bson:"namespace"`
	Age                    string `json:"age" bson:"ages"`
	DesiredNumberScheduled int32  `json:"desired_number_scheduled" bson:"desired_number_scheduled"`
	CurrentNumberScheduled int32  `json:"current_number_scheduled" bson:"current_number_scheduled"`
	NumberReady            int32  `json:"number_ready" bson:"number_ready"`
	NumberAvailable        int32  `json:"number_available" bson:"number_available"`
}
