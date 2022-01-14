package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joepaul-martin/banking-auth/domain"
	"github.com/joepaul-martin/banking-auth/service"
)

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:acer@dts@tcp(localhost:3306)/banking")
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

	fmt.Printf("Started the application")

	http.ListenAndServe("localhost:8001", router)

}
