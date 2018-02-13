package appcontext

import (
	"github.com/jmoiron/sqlx"
	"github.com/contacts_api_go/config"
	"github.com/contacts_api_go/logger"
	_ "github.com/lib/pq"
)


var db *sqlx.DB

func Initialize() {
	db = initDB()
}

func initDB() *sqlx.DB{
	db, e := sqlx.Open("postgres", config.GetDBConfig().ConnectionString())
	if e != nil {
		logger.Log.Errorf("Error Connecting DB - %s", e.Error())
		panic(e)
	}
	if err := db.Ping(); err != nil{
		logger.Log.Errorf("Ping to Db failed - %s", err.Error())
		panic(e)
	}
	db.SetMaxOpenConns(config.GetDBConfig().DatabaseMaxPoolSize())
	return  db;
}


func GetDB() *sqlx.DB {
	return db
}
