package main

import (
	"fmt"
	"log"
	"os"
)

type database struct {
	user     string
	pass     string
	dbhost   string
	port     string
	protocol string
	dbname   string
}

func (db *database) setDBConnection() {
	db.user = getEnv("DB_USER")
	db.pass = getEnv("DB_PASS")
	db.protocol = getEnv("DB_PROTOCOL")
	db.dbhost = getEnv("DB_HOST")
	db.port = getEnv("DB_PORT")
	db.dbname = getEnv("DB_NAME")
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("%s not found", key)
	return ""
}

func (db *database) connectionString() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		db.user, db.pass, db.protocol, db.dbhost, db.port, db.dbname)
}
