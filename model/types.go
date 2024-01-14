package model

import ()

type Transaction struct {
	UUID   string  `field:"uuid"`
	Name   string  `field:"name"`
	Day    int     `field:"day"`
	Month  int     `field:"month"`
	Year   int     `field:"year"`
	Amount float32 `field:"amount"`
}

func AddTransaction(uuid, name string, day, month, year int, amount float32) Transaction {
	return Transaction{
		UUID:   uuid,
		Name:   name,
		Day:    day,
		Month:  month,
		Year:   year,
		Amount: amount,
	}
}

type Category struct {
	Name            string
	CategoryFilters []string
	Transactions    []Transaction
	Budget          float32
}

func AddCategory(name string, budget float32) Category {
	return Category{
		Name:            name,
		Budget:          budget,
		Transactions:    nil,
		CategoryFilters: nil,
	}
}

type MonthlyBudget struct {
	Month      int
	Year       int
	Categories []Category
}

func AddMonthlyBudget(month, year int) MonthlyBudget {
	return MonthlyBudget{
		Month: month,
		Year:  year,
	}
}
