package config

import (
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	logLevel string
	port     int
	databaseConfig *databaseConfig
	statsdConfig   *statsdConfig
}

var appConfig *Config


func Load() {
	viper.SetDefault("APP_PORT", "3000")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	appConfig = &Config{
		logLevel: getString("LOG_LEVEL"),
		port:     getIntOrPanic("APP_PORT"),
		databaseConfig: newDatabaseConfig(),
		statsdConfig: newStastdConfig(),
	}

}

func GetAppConfig() *Config {
	return appConfig;
}

func GetAppPort() int {
	return appConfig.port
}

func GetLogLevel() string {
	return appConfig.logLevel
}

func GetDBConfig() *databaseConfig {
	return appConfig.databaseConfig
}

func GetStatsDAdderss() string{
	return fmt.Sprintf("%s:%d", appConfig.statsdConfig.host, appConfig.statsdConfig.port)
}

func IsStatsDEnabled() bool{
	return appConfig.statsdConfig.enabled
}

func GetStatsDAppName() string{
	return appConfig.statsdConfig.appName
}



