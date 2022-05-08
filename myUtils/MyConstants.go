package myProject

import (
	"fmt"
)

const (
	username     = "app"             // user name
	password     = "lizapepsi"       // passwaord
	hostname     = "127.0.0.1:3306"  // hostname and port of the DB
	dbname       = "test"            // DB Name
	hostmemcachd = "127.0.0.1:11211" // host name and port of the MemcacheD
)
const (
	SelectCities        = "select id, name, population from cities order by name"        // Select query
	SelectCitiesForJson = "select id, name, population from cities order by id limit 50" // Select query
	DeleteCities        = "DELETE FROM cities WHERE id IN (2, 4, 6)"                     // Delete query
	InsertStatement     = "INSERT INTO cities(id,name, population) VALUES(?,?,?)"        // Insert query
	UpdateStatement     = "update cities set name=? where id=?"                          // Update query
	CityName            = "Santa Rosa"                                                   // City name Value to be inserted
	Population          = "10002"                                                        // Population Value to be inserted
	Id                  = "12"                                                           // ID value to be inserted
	JsonURL             = "http://api.theysaidso.com/qod.json"                           // JSON URL
	DB_MYSQL            = "mysql"                                                        // DB Type
)

func dns(dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname) // Returns a single object with user, pass,hostname and db name
}
