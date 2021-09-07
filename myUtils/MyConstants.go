package myProject

import (
	"fmt"
)

const (
	username = "app"
	password = "lizapepsi"
	hostname = "127.0.0.1:3306"
	dbname   = "test"
)

func dns(dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
