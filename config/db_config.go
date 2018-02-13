package config

import "fmt"

type databaseConfig struct {
	host        string
	port        int
	username    string
	password    string
	name        string
	maxPoolSize int
}

func newDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		host:        fatalGetString("DB_HOST"),
		port:        getIntOrPanic("DB_PORT"),
		name:        fatalGetString("DB_NAME"),
		username:    getString("DB_USER"),
		password:    getString("DB_PASSWORD"),
		maxPoolSize: getIntOrPanic("DB_POOL"),
	}
}

func (dc *databaseConfig) DatabaseMaxPoolSize() int {
	return dc.maxPoolSize
}

func (dc *databaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s sslmode=disable", dc.name, dc.username, dc.password, dc.host)
}

func (dc *databaseConfig) DbName() string {
	return dc.name
}
