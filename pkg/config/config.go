package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadConfig(cfgFile string) {

	var (
		err           error
		configYaml    []byte
		configuration AppConfig
	)

	if configYaml, err = ioutil.ReadFile(cfgFile); err != nil {
		log.Fatalf("Configuration file could not be read!")
	}

	if err = yaml.Unmarshal(configYaml, &configuration); err != nil {
		panic(err)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(cfgFile)

	if err = viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
