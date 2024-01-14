package imports

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juandserrano/timb/model"
)

func importFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ImportTransactions(db *sql.DB, path, t string) error {
	file, err := importFile(path)
	if err != nil {
		return err
	}
	if t == "debit" {
		err = parseFileToDebitStruct(db, file)
		if err != nil {
			return err
		}
		return nil

	} else if t == "credit" {
		err = parseFileToCreditStruct(db, file)
		if err != nil {
			return err
		}
		return nil

	}
	return fmt.Errorf("Not a debit or credit transaction\n")
}

func parseFileToCreditStruct(db *sql.DB, file []byte) error {
	transactions := strings.Split(string(file), "\n")
	for _, t := range transactions[:len(transactions)-1] {
		data := strings.Split(t, ",")
		date := strings.Split(data[0], "/")
		month, err := strconv.Atoi(date[0])
		if err != nil {
			return err
		}
		year, err := strconv.Atoi(date[2])
		if err != nil {
			return err
		}
		day, err := strconv.Atoi(date[1])
		if err != nil {
			return err
		}
		amount, err := strconv.ParseFloat(data[2], 32)
		if err != nil {
			return err
		}
		name := strings.TrimSpace(strings.Trim(data[1], "\""))
		uuid := fmt.Sprintf("===%d%d%d-%s-%f===", year, month, day, name, amount)
		transaction := model.Transaction{
			UUID:   uuid,
			Year:   year,
			Month:  month,
			Day:    day,
			Amount: float32(amount),
			Name:   name,
		}
		err = saveTransaction(db, transaction)
		if err != nil {
			return err
		}
		// fmt.Printf("%v\n", transaction)

	}
	return nil

}

func parseFileToDebitStruct(db *sql.DB, file []byte) error {
	transactions := strings.Split(string(file), "\n")
	for _, t := range transactions[:len(transactions)-1] {
		data := strings.Split(t, ",")
		date := strings.Split(data[0], "/")
		month, err := strconv.Atoi(date[0])
		if err != nil {
			return err
		}
		year, err := strconv.Atoi(date[2])
		if err != nil {
			return err
		}
		day, err := strconv.Atoi(date[1])
		if err != nil {
			return err
		}
		amount, err := strconv.ParseFloat(data[1], 32)
		if err != nil {
			return err
		}
		name := strings.TrimSpace(strings.Trim(data[3], "\"") + " / " + strings.Trim(data[4], "\""))
		uuid := fmt.Sprintf("===%d%d%d-%s-%f===", year, month, day, name, amount)
		transaction := model.Transaction{
			UUID:   uuid,
			Year:   year,
			Month:  month,
			Day:    day,
			Amount: float32(amount),
			Name:   name,
		}
		saveTransaction(db, transaction)
		// fmt.Printf("%v\n", transaction)

	}
	return nil

}

func saveTransaction(db *sql.DB, t model.Transaction) error {
	res, err := db.Exec(
		"INSERT INTO transaction(uuid, name, year, month, day, amount) VALUES ($1, $2, $3, $4, $5, $6)", t.UUID, t.Name, t.Year, t.Month, t.Day, t.Amount)
	// res, err := db.Exec(fmt.Sprintf(
	// 	"INSERT INTO transactions(uuid,name,year,month,day,amount) VALUES (%s, %s, %d, %d, %d, %f)", t.UUID, t.Name, t.Year, t.Month, t.Day, t.Amount))
	if err != nil {
		fmt.Println("Error saveTransaction:", err)
		fmt.Println("Culprit:", t.Name)
		return err
	}
	id, err := res.RowsAffected()
	fmt.Println("Affected Rows:", id)
	return nil

}
