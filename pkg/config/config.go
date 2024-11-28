package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/pelletier/go-toml/v2"
)

type WebServer struct {
	Port        int    `toml:"port"`
	RoutePrefix string `toml:"route_prefix"`
}

type App struct {
	WebServer WebServer
}

func LoadConfig() (*App, error) {
	_, d, _, _ := runtime.Caller(0)
	b, err := os.ReadFile(filepath.Join(filepath.Dir(d), "/../../etc/config.local.toml"))
	if err != nil {
		return nil, err
	}

	app := &App{}
	err = toml.Unmarshal(b, app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
