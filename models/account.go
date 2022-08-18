package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Account struct {
	Id     int64   `json:"id"`
	Number string  `json:"account_number"`
	Amount float64 `json:"amount"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./db/desafio1.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetAccounts() ([]Account, error) {
	rows, err := DB.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := make([]Account, 0)
	for rows.Next() {
		singleAccount := Account{}
		err = rows.Scan(&singleAccount.Id, &singleAccount.Number, &singleAccount.Amount)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, singleAccount)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func InsertAccount(account Account) (id int64) {
	stmt, err := DB.Prepare("INSERT INTO account(account_number, amount) VALUES($1, $2)")
	if err != nil {
		return 0
	}

	res, err := stmt.Exec(account.Number, account.Amount)
	if err != nil {
		return 0
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return 0
	}

	err = stmt.Close()
	if err != nil {
		return 0
	}

	return lid
}
