package config

import "github.com/spf13/viper"

type Config struct {
	logLevel string
	port     int
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
