package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joepaul-martin/banking-auth/domain"
	"github.com/joepaul-martin/banking-auth/errs"
	"github.com/joepaul-martin/banking-auth/service"
	"gopkg.in/yaml.v3"
)

const configFileName = "config.yaml"

type serverData struct {
	Server struct {
		ServerAddress string `yaml:"serverAddress"`
		ServerPort    string `yaml:"serverPort"`
	} `yaml:"server"`
}

type databaseData struct {
	MySql struct {
		DB       string `yaml:"db"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		DbName   string `yaml:"dbName"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
}

func readServerConfig(fileName string) (*serverData, *errs.AppError) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("Error while reading the config file: %s, error : %s", fileName, err.Error()))
	}
	data := &serverData{}
	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("Error while unmarshalling object: %s", err.Error()))
	}
	return data, nil
}

func readDBConfig(fileName string) (*databaseData, *errs.AppError) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("Error while reading the config file: %s, error : %s", fileName, err.Error()))
	}
	data := &databaseData{}
	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("Error while unmarshalling object: %s", err.Error()))
	}
	return data, nil
}

func getDbClient() *sqlx.DB {
	dbConfig, appErr := readDBConfig(configFileName)
	if appErr != nil {
		log.Fatal(appErr)
	}
	password := os.Getenv("DBPASSWORD")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.MySql.User, password, dbConfig.MySql.Port, dbConfig.MySql.DbName)
	client, err := sqlx.Open(dbConfig.MySql.DB, dataSourceName)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func Start() {
	router := mux.NewRouter()

	authRepositoryDb := domain.NewAuthRepositoryDb(getDbClient())
	lh := LoginHandler{
		service: service.NewDefaultLoginService(authRepositoryDb),
	}

	router.HandleFunc("/login", lh.login).Methods(http.MethodPost)

	// starting server
	serverConfig, err := readServerConfig(configFileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverConfig.Server.ServerAddress, serverConfig.Server.ServerPort), router))

}
