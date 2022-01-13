package controllers

import (
	"github.com/spf13/viper"
	"time"
	"valse/pkg/_errors"
	"valse/pkg/config"
	"valse/pkg/k8s"
	mongodb "valse/pkg/mongodb"
)

func Init() {

	config.ReadConfig(".configs/config.yaml")

	var logger = _errors.NewLogger()

	var appConfig *config.AppConfig

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		logger.Fatalf("Configuration is invalid! %e", err)
	}

	var mongoClient = mongodb.NewClient(appConfig, logger)
	var k8sClient = k8s.NewClient(appConfig, logger)

	for {
		RunScheduledJob(appConfig, k8sClient, mongoClient, logger)
		<-time.After(appConfig.ScheduledTaskIntervalSeconds * time.Second)
	}
}
