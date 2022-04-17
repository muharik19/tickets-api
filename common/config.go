package common

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	RootURL    string `split_words:"true" default:"/api/v1/ticket"`
	Port       string `split_words:"true" default:"5007"`
	DBHost     string `split_words:"true" default:"localhost"`
	DBPort     string `split_words:"true" default:"3306"`
	DBUser     string `split_words:"true" default:"root"`
	DBPassword string `split_words:"true" default:""`
	DBName     string `split_words:"true" default:"ticket"`
}

var Config Configuration
var ctx = context.Background()

func InitConfig() {
	err := envconfig.Process("ticket", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithField("config", Config).Info("Config successfully loaded")
}
