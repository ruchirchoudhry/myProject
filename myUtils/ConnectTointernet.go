package myProject

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

const (
	username = "app"
	password = "lizapepsi"
	hostname = "127.0.0.1:3306"
	dbname   = "test"
)

func dns(dbname string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func GetJasonFromInternet() {
	resp, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))

}
func GetDataFromMySQL() {

	db, err := sql.Open("mysql", dns(dbname))

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	res, err := db.Query("select * from cities order by name")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	for res.Next() {
		var cities City
		err := res.Scan(&cities.Id, &cities.Name, &cities.Population)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%v\n", cities)
	}
	defer db.Close()
}
func DeleteDataFromMySQL() {
	db, err := sql.Open("mysql", dns(dbname))

	if err != nil {
		log.Fatal(err)
		return
	}
	res, err := db.Exec("DELETE FROM cities WHERE id IN (2, 4, 6)")
	if err != nil {
		log.Panic(err)
		return
	}

	//var affectedRow int64
	affectedRow, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	fmt.Printf("The statement affected %d rows\n", affectedRow)
}
