package db_connect

const (
	host       = "localhost"
	port       = "5432"
	database   = "database_name"
	user       = "user_name"
	password   = "password"
	DataSource = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + database + " sslmode=disable"
)
