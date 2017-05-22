package bench

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
)

func Step1(index1, index2 int) {
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://test@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Create the "accounts" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS accounts (id INT PRIMARY KEY, balance INT)"); err != nil {
		log.Fatal(err)
	}

	// Insert two rows into the "accounts" table.
	query := fmt.Sprintf("INSERT INTO accounts (id, balance) VALUES (%d, 100000000000), (%d, 100000000000)", index1, index2)
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	// Print out the balances.
	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances:")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %d\n", id, balance)
	}
}

func transferFunds(tx *sql.Tx, from int, to int, amount int) error {
	// Read the balance.
	var fromBalance int
	if err := tx.QueryRow(
		"SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance); err != nil {
		return err
	}
	/*if _, err := tx.Exec("UPDATE accounts SET balance = $1 where id = $2", 10000000000, from); err != nil {
		println(err.Error())
	}*/

	if fromBalance < amount {

		return fmt.Errorf("insufficient funds")
	}

	// Perform the transfer.
	if _, err := tx.Exec(
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from); err != nil {
		return err
	}
	if _, err := tx.Exec(
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to); err != nil {
		return err
	}
	return nil
}

func Step2() {
	db, err := sql.Open("postgres", "postgresql://test@localhost:26257/bank?sslmode=disable")
	if err != nil {
		println("error connecting to the database: ", err)
	}

	// Run a transfer in a transaction.
	err = crdb.ExecuteTx(db, func(tx *sql.Tx) error {
		return transferFunds(tx, 3 /* from acct# */, 4 /* to acct# */, 100 /* amount */)
	})
	if err == nil {
		println("Success")
	} else {
		println("error: ", err.Error())
	}
}
