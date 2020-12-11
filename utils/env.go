package utils

import (
	"os"

	"github.com/fatih/color"
)

// EnvConfig a handy util to load environment variables
type EnvConfig struct {
	ServerPort string
}

// NewConfig returns a new instance of Envconfig with default values
func NewConfig() *EnvConfig {
	defaultport := ":8080"

	return &EnvConfig{defaultport}

}

// LoadEnv loads required environment variables if present
func (ec *EnvConfig) LoadEnv() string {
	sp := os.Getenv("SERVER_PORT")
	var port string

	if sp != " " {
		ec.ServerPort = ":" + sp
		port = *&ec.ServerPort
		return port

	}
	color.Yellow("SERVER_PORT not found using default")
	ec.ServerPort = ":" + "8080"
	port = *&ec.ServerPort
	return port

}
