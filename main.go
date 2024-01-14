package main

import (
	"fmt"

	"github.com/juandserrano/timb/db"
	"github.com/juandserrano/timb/imports"
	"github.com/juandserrano/timb/model"
)

func main() {
	db, err := db.ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = imports.ImportTransactions(db, "/Users/wizzi/Downloads/pcbanking (1).ascii", "credit")
	if err != nil {
		fmt.Println(err)
	}
	err = imports.ImportTransactions(db, "/Users/wizzi/Downloads/pcbanking.ascii", "debit")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT amount, day, name, uuid, year, month from transaction")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Transaction
		err = rows.Scan(&t.Amount, &t.Day, &t.Name, &t.UUID, &t.Year, &t.Month)
		if err != nil {
			fmt.Println("Scan error:", err)
		}
		fmt.Println(t)

	}

}
