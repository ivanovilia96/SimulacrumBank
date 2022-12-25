package data_base

import (
	"database/sql"
	"fmt"
)

type Human struct {
	FIO string
	Age string
}

func GetMock() {
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
