package myProject

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
func GetDataFromMySQLAndConvToJson() {

	db, err := sql.Open("mysql", dns(dbname))
	errstring := "Error %s when opening DB\n"
	CheckErrorsWithPrintStr(err, errstring)
	rows, err := db.Query(SelectCitiesForJson)
	CheckErrors(err)
	connectionPoolsettings()
	cities := []MyCitiesModel{}
	for rows.Next() {
		city := MyCitiesModel{}
		rows.Scan(&city.ID, &city.Name, &city.Population)
		cities = append(cities, city)
	}
	ub, _ := json.MarshalIndent(&cities, "", "  ")
	fmt.Println(string(ub))
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
	CheckErrorsWithRowAffected(res, err)
	CheckErrorsTx(err, *tx)
	tx.Commit()
	id, err := res.LastInsertId()
	CheckErrors(err)
	defer db.Close()
	fmt.Println("Inserted ID=", id)

}
func UpdateCityDataWithTx() {
	connectionPoolsettings()
	db, err := sql.Open(DB_MYSQL, dns(dbname)) // Opening  connections
	CheckErrors(err)                           // Checking for connection errors
	tx, _ := db.Begin()                        // Bigning the transaction
	stmt, err := tx.Prepare(UpdateStatement)   // passing the updatestatement or error in case its caught
	CheckErrors(err)                           // Check for errros
	res, err01 := stmt.Exec("LA", "13")        // collecting te respose or error in the callback
	CheckErrorsWithRowAffected(res, err01)     // returns and print Rows effected in the system
	CheckErrorsTx(err, *tx)                    // Checking errors if any
	tx.Commit()                                // Commiting Transaction (Tx)
	defer db.Close()                           // Defered connection cleanup

}

/**
 * This function is used for set connections to e DB/
**/
func connectionPoolsettings() {
	db, err := sql.Open("mysql", dns(dbname)) // opening up connections
	CheckErrorsWithReturn(err)                // Checking errors if any
	db.SetMaxOpenConns(5)                     // setting max Open Connections
	db.SetMaxIdleConns(3)                     // setting max Idle Connections
	db.SetConnMaxLifetime(time.Minute * 1)    // Setting max life
	db.Stats()                                // Gets the stats of the DB
	defer db.Close()                          // Defered connection cleanup
}
