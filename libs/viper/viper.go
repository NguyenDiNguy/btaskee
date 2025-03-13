package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Endpoint  string `mapstructure:"endpoint"`
	PortHTTP  string `mapstructure:"port_http"`
	PortGRPC  string `mapstructure:"port_grpc"`
	DBuri     string `mapstructure:"db_uri"`
	DBName    string `mapstructure:"db_name"`
	RedisUrl  string `mapstructure:"redis_url"`
	RedisPass string `mapstructure:"redis_pass"`
}

var GlobalConfig Config

func init() {
	dir := os.Getenv("XDIR")
	if len(dir) <= 0 {
		dir = "../.."
	}

	dir += "/"
	os.Setenv("XDIR", dir)
	addConfig(dir, "config", "yml")

	// Parser config to proto
	parseConfigProto(dir)
}

func addConfig(path, configName, configType string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.ReadInConfig()
}

func parseConfigProto(dir string) {
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println("Failed to parse config:", err)
	}
}
