package database

import (
	"database/sql"
	"log"

	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type DBInstance struct {
	Db  *sql.DB
	Orm orm.Ormer
}

var Database DBInstance

func ConnectDB() {
	// gunakan 5432 (port di docker, karena akan jalan di image docker) --> 5435:5432
	// var dbUrl = "postgresql://postgres:postgres@db:5432/web_microservice?sslmode=disable"
	// orm.RegisterDriver("postgres", orm.DRPostgres)

	// substituted our env vars into our connection string.
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// "default" alias, "postgres" driver name, dbUrl datasource
	if err := orm.RegisterDataBase("default", "postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	}

	if db, err := sql.Open("postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	} else {
		if err := orm.RunSyncdb("default", false, true); err != nil {
			log.Fatal(err.Error())
		} else {
			// We set the debug field to true so we can see the queries in our terminal.
			orm.Debug = true
			o := orm.NewOrm()
			Database = DBInstance{Db: db, Orm: o}
			log.Print("Connected Succcessfully")
		}
	}
}
