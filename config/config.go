package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	USERNAME string = "root"
	PASSWORD string = "root"
	HOST     string = "localhost"
	PORT     int    = 3306
	DATABASE string = "online-shop"
)

func SetConfigDeployConnection() {
	USERNAME = os.Getenv("MYSQL_USERNAME")
	PASSWORD = os.Getenv("MYSQL_PASSWORD")
	DATABASE = os.Getenv("MYSQL_DATABASE")
	HOST = os.Getenv("MYSQL_HOST")

	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatalf("not port selected in connection %v", err)
	}
	PORT = port
}

func GetURLConnection() string {
	dev := os.Getenv("PORT")
	if dev != "" {
		SetConfigDeployConnection()
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", USERNAME, PASSWORD, HOST, PORT, DATABASE)
}
