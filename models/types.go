package models

// t_users
type TUsers struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Mail string `json:"mail" db:"mail"`
}
