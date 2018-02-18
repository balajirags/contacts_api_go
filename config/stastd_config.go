package config

type statsdConfig struct {
	appName string
	port    int
	enabled bool
	host    string
}

func newStastdConfig() *statsdConfig {
	return &statsdConfig{
		appName: fatalGetString("STATSD_APP_NAME"),
		port:    getIntOrPanic("STATSD_PORT"),
		enabled: getBoolWithDefault("STATSD_ENABLED", true),
		host:    fatalGetString("STATSD_HOST"),
	}
}

