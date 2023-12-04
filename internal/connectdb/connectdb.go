package connectdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("Connecting DB...")
}

func GetDBValue() (string, error) {
	rows, err := db.Query(`SELECT kanji_month FROM month WHERE id=1;`)
	if err != nil {
		fmt.Println("db.Query: ", err)
		return "no kanji", err
	}

	fmt.Println("rows:", rows)

	var kanji string
	for rows.Next() {
		if err := rows.Scan(&kanji); err != nil {
			fmt.Println("rows.Next: ", err)
			return "no kanji", err
		}
	}
	return kanji, nil
}
