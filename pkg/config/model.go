package config

import "time"

type AppConfig struct {
	Cluster struct {
		Name   string `json:"name" yaml:"name"`
		Env    string `json:"env" yaml:"env"`
		Region string `json:"region" yaml:"region"`
	} `json:"cluster" yaml:"cluster"`

	ScheduledTaskIntervalSeconds time.Duration `json:"scheduled_task_interval_seconds" yaml:"scheduledTaskIntervalSeconds"`

	ExcludedNamespaces string `json:"excluded_namespaces" yaml:"excludedNamespaces"`

	MongoDB struct {
		Hosts          []string      `json:"hosts" yaml:"hosts"`
		Username       string        `json:"username" yaml:"username"`
		Password       string        `json:"password" yaml:"password"`
		Port           string        `json:"port" yaml:"port"`
		Database       string        `json:"database" yaml:"database"`
		ReplicaSet     string        `json:"replica_set" yaml:"replicaSet"`
		TimeoutSeconds time.Duration `json:"timeout_seconds" yaml:"timeoutSeconds"`
	} `json:"mongo_db" yaml:"mongoDB"`

	Client struct {
		InClusterConfig bool `json:"in_cluster_config" yaml:"inClusterConfig"`
	} `json:"client" yaml:"client"`
}
