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

const (
	SelectCities    = "select id, name, population from cities order by name"
	DeleteCities    = "DELETE FROM cities WHERE id IN (2, 4, 6)"
	InsertStatement = "INSERT INTO cities(id,name, population) VALUES(?,?,?)"
	CityName        = "Bangalore"
	Population      = "1123000"
	Id              = "21"
	JsonURL         = "http://api.theysaidso.com/qod.json"
)

func dns(dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
