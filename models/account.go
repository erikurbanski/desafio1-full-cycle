package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Account struct {
	Id     int64   `json:"id"`
	Number string  `json:"account_number"`
	Amount float64 `json:"amount"`
}

type Transfer struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
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

func TransferValues(transfer Transfer) (msg string) {
	stmtFrom, errFrom := DB.Prepare("SELECT * FROM account WHERE account_number = ?")
	checkErr(errFrom)

	rowsFrom, errQueryFrom := stmtFrom.Query(transfer.From)
	checkErr(errQueryFrom)

	var accountFrom Account
	for rowsFrom.Next() {
		errFrom = rowsFrom.Scan(&accountFrom.Id, &accountFrom.Number, &accountFrom.Amount)
		checkErr(errFrom)
	}

	stmtTo, errTo := DB.Prepare("SELECT * FROM account WHERE account_number = ?")
	checkErr(errTo)

	rowsTo, errQueryTo := stmtTo.Query(transfer.To)
	checkErr(errQueryTo)

	var accountTo Account
	for rowsTo.Next() {
		errTo = rowsTo.Scan(&accountTo.Id, &accountTo.Number, &accountTo.Amount)
		checkErr(errTo)
	}

	var newAmountFrom = accountFrom.Amount - transfer.Amount
	var newAmountTo = accountTo.Amount + transfer.Amount

	fmt.Println(newAmountFrom, newAmountTo)

	stmtUpdateFrom, errUpdateFrom := DB.Prepare("UPDATE account SET amount = $1 WHERE account_number = $2")
	checkErr(errUpdateFrom)

	resultUpdateFrom, errExecUpdateFrom := stmtUpdateFrom.Exec(newAmountFrom, transfer.From)
	checkErr(errExecUpdateFrom)

	rowAffectedFrom, errLastUpdateFrom := resultUpdateFrom.RowsAffected()
	checkErr(errLastUpdateFrom)

	stmtUpdateTo, errUpdateTo := DB.Prepare("UPDATE account SET amount = $1 WHERE account_number = $2")
	checkErr(errUpdateTo)

	resultUpdateTo, errExecUpdateTo := stmtUpdateTo.Exec(newAmountTo, transfer.To)
	checkErr(errExecUpdateTo)

	rowAffectedTo, errLastUpdateTo := resultUpdateTo.RowsAffected()
	checkErr(errLastUpdateTo)

	if rowAffectedFrom > 0 && rowAffectedTo > 0 {
		return "Transfer carried out!"
	}

	return ""
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
