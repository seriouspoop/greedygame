package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"seriouspoop/greedygame/go-common/db/postgres"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envPrefix = "DL_MS"
)

type Postgres struct {
	postgres.Config
}
type WebServer struct {
	RestPort int `required:"true" split_words:"true"`
	GrpcPort int `required:"true" split_words:"true"`
}

type App struct {
	ServiceName string `required:"true" split_words:"true"`
	LogLevel    string `required:"true" split_words:"true"`
	WebServer   WebServer
	Postgres    Postgres
}

func FromEnv() (*App, error) {
	fromFileToEnv()
	cfg := &App{}
	if err := envconfig.Process(envPrefix, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func fromFileToEnv() {
	cfgFileName := os.Getenv("ENV_FILE")
	if cfgFileName != "" {
		if err := godotenv.Load(cfgFileName); err != nil {
			fmt.Println("error: failure reading ENV_FILE file, ", err)
		} else {
			return
		}
	}

	_, b, _, _ := runtime.Caller(0)
	cfgFileName = filepath.Join(filepath.Dir(b), "../../etc/config.local.env")

	if err := godotenv.Load(cfgFileName); err != nil {
		fmt.Println("error: failure reading config file:, ", err)
	}
}
