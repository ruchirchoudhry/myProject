package myProject

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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
	errstring := "Error %s when opening DB\n"
	CheckErrorsWithPrintStr(err, errstring)
	res, err := db.Query(SelectCities)
	CheckErrors(err)
	connectionPoolsettings()
	for res.Next() {
		var cities City
		err := res.Scan(&cities.Id, &cities.Name, &cities.Population)
		CheckErrors(err)
		fmt.Printf("%v\n", cities)
	}
	defer db.Close()
}
func DeleteDataFromMySQL() {
	db, err := sql.Open("mysql", dns(dbname))
	CheckErrorsWithReturn(err)
	res, err := db.Exec(DeleteCities)
	CheckErrors(err)
	affectedRow, err := res.RowsAffected()
	CheckErrorsWithReturn(err)
	defer db.Close()
	fmt.Printf("The statement affected %d rows\n", affectedRow)
}
func InsertIntoCityWithTx() {
	connectionPoolsettings()
	db, err := sql.Open("mysql", dns(dbname))
	CheckErrors(err)
	tx, _ := db.Begin()
	stmt, err := tx.Prepare(InsertStatement)
	fmt.Println("Insert statemetn used", stmt)
	CheckErrors(err)
	res, err := stmt.Exec(Id, CityName, Population)
	println(res)
	CheckErrorsTx(err, *tx)
	tx.Commit()
	id, err := res.LastInsertId()
	CheckErrors(err)
	fmt.Println("Inserted ID=", id)
	defer db.Close()
}
func connectionPoolsettings() {
	db, err := sql.Open("mysql", dns(dbname))
	CheckErrorsWithReturn(err)
	defer db.Close()
	db.SetMaxOpenConns(5)                  // setting max Open Connections
	db.SetMaxIdleConns(3)                  // setting max Idle Connections
	db.SetConnMaxLifetime(time.Minute * 1) // Setting max life
	db.Stats()                             // Gets the stats of the DB
}
