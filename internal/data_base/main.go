package data_base

import (
	"database/sql"
	"fmt"
)

func New(driverName, login, password, host, dbName string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", driverName, login, password, host, dbName)
	//connStr := "postgres://postgres:100@localhost/bank?sslmode=disable"
	return sql.Open("postgres", conn)
}

type Human struct {
	FIO string
	Age string
}

func getMock() {
	connStr := "postgres://postgres:100@localhost/bank?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	result, err := db.Query(`insert into clientcashaccountlink values ('1',1)`)
	if err != nil {
		panic(err)
	}

	var humans []Human
	for result.Next() {
		var h Human
		result.Scan(&h.FIO, &h.Age)
		humans = append(humans, h)
	}

	fmt.Println(humans)
}
