package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"golang-api-template/configs/db_connect"
	user_add "golang-api-template/internal/app/user/add"
	user_delete "golang-api-template/internal/app/user/delete"
	user_edit "golang-api-template/internal/app/user/edit"
	user_get "golang-api-template/internal/app/user/get"
)

func main() {
	// open db
	db, err := sql.Open("postgres", db_connect.DataSource) // using pos driver
	if err != nil {
		fmt.Println("failed opening db: ", err)
		return
	}
	defer db.Close()

	// connect db
	err = db.Ping()
	if err != nil {
		fmt.Println("failed connecting db: ", err)
		return
	}

	// set routing
	router := setRouting()

	// run server
	err = http.ListenAndServe(fmt.Sprintf(":%d", 8000), router)
	if err != nil {
		log.Fatal("failed running server: ", err)
	}
	fmt.Println("server is running!")

	return
}

// set routing
func setRouting() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true) // using gorilla/mux (gorilla/mux archived. -> will replace gin)

	// example routing
	router.HandleFunc("/user/get/{user_id}", user_get.Exec).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/add", user_add.Exec).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/edit", user_edit.Exec).Methods("PUT", "OPTIONS")
	router.HandleFunc("/user/delete/{user_id}", user_delete.Exec).Methods("DELETE", "OPTIONS")

	return
}
