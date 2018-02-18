package appcontext

import (
	"github.com/jmoiron/sqlx"
	"github.com/contacts_api_go/config"
	"github.com/contacts_api_go/logger"
	_ "github.com/lib/pq"
	statsdv2 "gopkg.in/alexcesaro/statsd.v2"
	"fmt"
)

var db *sqlx.DB

var statsD *statsdv2.Client

func Initialize() {
	db = initDB()
	statsD = initializeStatsdClient()
}

func initDB() *sqlx.DB {
	db, e := sqlx.Open("postgres", config.GetDBConfig().ConnectionString())
	if e != nil {
		logger.Log.Errorf("Error Connecting DB - %s", e.Error())
		panic(e)
	}
	if err := db.Ping(); err != nil {
		logger.Log.Errorf("Ping to Db failed - %s", err.Error())
		panic(e)
	}
	db.SetMaxOpenConns(config.GetDBConfig().DatabaseMaxPoolSize())
	return db;
}

func initializeStatsdClient() *statsdv2.Client {
	fmt.Println(config.GetStatsDAdderss())
	client, err := statsdv2.New(statsdv2.Address(config.GetStatsDAdderss()), statsdv2.Prefix(config.GetStatsDAppName()))
	if err != nil {
		panic(err)
	}
	return client
}

func GetDB() *sqlx.DB {
	return db
}

func GetStatsDClient() *statsdv2.Client {
	return statsD
}
