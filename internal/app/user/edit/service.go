package user_edit

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

	// get user info
	var tUser models.TUsers
	err := json.NewDecoder(r.Body).Decode(&tUser)
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

	// update user
	sqlOrder := "UPDATE t_users SET (name, mail)=($1, $2) WHERE id=$3"
	_, err = db.Exec(sqlOrder, tUser.Name, tUser.Mail, tUser.Id)
	if err != nil {
		fmt.Println("failed updating user: ", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	return
}
