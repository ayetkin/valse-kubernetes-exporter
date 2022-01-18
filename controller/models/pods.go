package models

import (
	"time"
)

type Pods struct {
	Name              string              `json:"name" bson:"name"`
	Namespace         string              `json:"namespace" bson:"namespace"`
	Age               string              `json:"age" bson:"age"`
	Phase             string              `json:"phase" bson:"phase"`
	Reason            string              `json:"reason" bson:"reason"`
	Message           string              `json:"message" bson:"message"`
	OwnerReferences   *PodOwnerReferences `json:"owner_references" bson:"owner_references"`
	HostIP            string              `json:"host_ip" bson:"host_ip"`
	ContainerStatuses []ContainerStatuses `json:"container_statuses" bson:"container_statuses"`
}

type PodOwnerReferences struct {
	Kind string `json:"kind" bson:"kind"`
	Name string `json:"name" bson:"name"`
}

type ContainerStatuses struct {
	Name         string                 `json:"name" bson:"name"`
	Ready        bool                   `json:"ready" bson:"ready"`
	Started      *bool                  `json:"started" bson:"started"`
	RestartCount int32                  `json:"restart_count" bson:"restart_count"`
	State        ContainerStatusesState `json:"state" bson:"state"`
}

type ContainerStatusesState struct {
	Waiting    *ContainerStatusesStateWaiting
	Running    *ContainerStatusesStateRunning
	Terminated *ContainerStatusesTerminated
}

type ContainerStatusesStateWaiting struct {
	Reason  string `json:"reason" bson:"reason"`
	Message string `json:"message"  bson:"message"`
}

type ContainerStatusesStateRunning struct {
	StartedAt time.Time `json:"started_at" bson:"started_at"`
}

type ContainerStatusesTerminated struct {
	ExitCode int32  `json:"exitCode" bson:"exit_code"`
	Signal   int32  `json:"signal,omitempty" bson:"signal"`
	Reason   string `json:"reason,omitempty" bson:"reason"`
	Message  string `json:"message,omitempty" bson:"message"`
}
