package main

import (
	"log"
	"github.com/quasar-man/dockertest-sample/infrastructure"
)

func init() {
	// GORM Auto migration
	mysqlConn := infrastructure.NewMySQLConnection()
	db, err := mysqlConn.DBConnectionOpen()
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	mysqlConn.GormAutoMigrate(db)
}

func main() {
	log.Println("Hello, World!")
}
