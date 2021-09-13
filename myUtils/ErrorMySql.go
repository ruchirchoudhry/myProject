package myProject

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CheckErrors(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err.Error())

	}
}
func CheckErrorsTx(err error, tx sql.Tx) {
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}
func CheckErrorsWithReturn(err error) {
	if err != nil {
		log.Fatal(err)
		return

	}
}
func CheckErrorsWithPrintStr(err error, errsting string) {
	if err != nil {
		log.Printf(errsting, err)
		return
	}
}
