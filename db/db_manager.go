package db

import (
	"database/sql"
	"fmt"
	"go-laundry/config"
	"go-laundry/util"
	"log"
)

type DbManager interface {
	ConnectDb() *sql.DB
}

type dbManager struct {
	db *sql.DB
}

func (m *dbManager) ConnectDb() *sql.DB {
	return m.db
}

func NewDbManager(config config.Config) DbManager {
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open("postgres", datasource)

	util.CheckErr(err)

	defer func() {
		if err := recover(); err != nil {
			log.Println("application failed to run")
			db.Close()
		}
	}()

	return &dbManager{
		db,
	}
}
