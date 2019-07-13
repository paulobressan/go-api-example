package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// Importing for open connection with mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DB : Instancia pool do banco de dados
	DB = initDatabase
)

type configDatabase struct {
	name        string
	host        string
	port        string
	username    string
	password    string
	maxOpenCon  int
	maxIdleCon  int
	maxLifetime int
}

func initDatabase() *gorm.DB {
	configDatabase, err := getEnvsDatabase()
	if err != nil {
		log.Fatal("Invalid environments")
		os.Exit(1)
	}
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDatabase.username,
		configDatabase.password,
		configDatabase.host,
		configDatabase.port,
		configDatabase.name)

	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Fatal("Fail to open database")
		os.Exit(1)
	}

	db.DB().SetMaxIdleConns(configDatabase.maxIdleCon)
	db.DB().SetMaxOpenConns(configDatabase.maxOpenCon)

	if configDatabase.maxLifetime > 0 {
		db.DB().SetConnMaxLifetime(time.Millisecond * time.Duration(configDatabase.maxLifetime))
	} else {
		db.DB().SetConnMaxLifetime(0)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Fail ping in Database", err)
		os.Exit(1)
	}

	return db
}

func getEnvsDatabase() (configDatabase, error) {
	var err error
	var configDatabase configDatabase

	configDatabase.host, err = validateEnv("DATABASE_HOST", "Invalid environment DATABASE_HOST")
	configDatabase.port, err = validateEnv("DATABASE_PORT", "Invalid environment DATABASE_PORT")
	configDatabase.name, err = validateEnv("DATABASE_NAME", "Invalid environment DATABASE_NAME")
	configDatabase.password, err = validateEnv("DATABASE_PASSWORD", "Invalid environment DATABASE_PASSWORD")
	configDatabase.username, err = validateEnv("DATABASE_USERNAME", "Invalid environment DATABASE_USERNAME")
	configDatabase.maxIdleCon, err = strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE"))
	configDatabase.maxLifetime, err = strconv.Atoi(os.Getenv("DATABASE_IDLE_TIMEOUT"))
	configDatabase.maxOpenCon, err = strconv.Atoi(os.Getenv("DATABASE_MAX_CON"))

	if err != nil {
		return configDatabase, err
	}
	return configDatabase, nil
}

func validateEnv(env string, errMessage string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		log.Fatal(errMessage)
		return "", errors.New(errMessage)
	}
	return value, nil
}
