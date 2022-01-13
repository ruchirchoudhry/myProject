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
	SelectCities        = "select id, name, population from cities order by name"
	SelectCitiesForJson = "select id, name, population from cities order by id limit 50"
	DeleteCities        = "DELETE FROM cities WHERE id IN (2, 4, 6)"
	InsertStatement     = "INSERT INTO cities(id,name, population) VALUES(?,?,?)"
	UpdateStatement     = "update cities set name=? where id=?"
	CityName            = "Burbank CA"
	Population          = "70000"
	Id                  = "75"
	JsonURL             = "http://api.theysaidso.com/qod.json"
)

func dns(dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
