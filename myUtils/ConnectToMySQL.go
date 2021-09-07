package myProject

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func GetJasonFromInternet() {
	resp, err := http.Get(JsonURL)
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
	res, err := db.Query(SelectCities)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	db.SetMaxOpenConns(5)                  // setting max Open Connections
	db.SetMaxIdleConns(3)                  // setting max Idle Connections
	db.SetConnMaxLifetime(time.Minute * 1) // Setting max life

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
	res, err := db.Exec(DeleteCities)
	if err != nil {
		log.Panic(err)
		return
	}

	affectedRow, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	fmt.Printf("The statement affected %d rows\n", affectedRow)
}
