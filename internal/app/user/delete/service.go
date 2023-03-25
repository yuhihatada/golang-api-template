package user_delete

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"golang-api-template/configs/db_connect"
	"golang-api-template/internal/pkg/common"
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

	// delete user
	sqlOrder := "DELETE FROM t_users WHERE id=$1"
	_, err = db.Exec(sqlOrder, userID)
	if err != nil {
		fmt.Println("failed deleting user: ", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	return
}
