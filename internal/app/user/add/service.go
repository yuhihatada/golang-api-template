package user_add

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang-api-template/configs/db_connect"
	"golang-api-template/internal/pkg/common"
)

type request struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func Exec(w http.ResponseWriter, r *http.Request) {
	// preflight
	common.SetPreflightHeader(&w, r)
	if r.Method == http.MethodOptions {
		return
	}

	// get user info
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("failed decode json to struct: ", err)
		w.WriteHeader(400)
		return
	}
	defer r.Body.Close()

	// open db
	db, err := sql.Open("postgres", db_connect.DataSource)
	if err != nil {
		fmt.Println("failed opening db: ", err)
		return
	}
	defer db.Close()

	// insert user
	sqlOrder := "INSERT INTO t_users (name, mail) VALUES($1, $2)"
	_, err = db.Exec(
		sqlOrder,
		req.Name,
		req.Mail,
	)
	if err != nil {
		fmt.Println("failed inserting data: ", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	return
}
