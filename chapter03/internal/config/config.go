package config

import (
	"fmt"
	"os"
	"taksmanagement/internal/web"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stackus/dotenv"
)

type (
	AppConfig struct {
		Environment string `envconfig:"ENVIRONMENT" default:"development"`
		LogLevel    string `envconfig:"LOG_LEVEL" default:"debug"`
		Rpc RpcConfig
		Web web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
	RpcConfig struct {
		Host string `default:"0.0.0.0"`
		Port string `default:":8085"`
	}
)

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
