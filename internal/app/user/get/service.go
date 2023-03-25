package user_get

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"golang-api-template/configs/db_connect"
	"golang-api-template/internal/pkg/common"
	"golang-api-template/models"
)

func Exec(w http.ResponseWriter, r *http.Request) {
	// preflight
	common.SetPreflightHeader(&w, r)
	if r.Method == http.MethodOptions {
		return
	}

	// get user_id
	userID := mux.Vars(r)["user_id"]

	// open db
	db, err := sql.Open("postgres", db_connect.DataSource)
	if err != nil {
		fmt.Println("failed opening db: ", err)
		return
	}
	defer db.Close()

	// select user
	var tUser models.TUsers
	sqlOrder := "SELECT * FROM t_users WHERE id=$1"
	if err := db.QueryRow(sqlOrder, userID).Scan(&tUser); err != nil {
		fmt.Println("failed selecting data: ", err)
		w.WriteHeader(400)
		return
	}

	// encode struct to json
	json.NewEncoder(w).Encode(tUser)

	w.WriteHeader(200)
	return
}
