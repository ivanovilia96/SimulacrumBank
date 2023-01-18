package client

import (
	"fmt"

	"simulacrumBank/internal/adapters/bisenes_logic"
	"simulacrumBank/internal/adapters/data_base"
)

func NewActions(db data_base.DataBaseActions) bisenes_logic.ClientActions {
	// сюда нужно передать бд чере адаптер\интерфейс
	return Client{
		db: db,
	}
}

type Client struct {
	db data_base.DataBaseActions
}

func (c Client) Add(mail, fio string, age int8) error {
	_, err := c.db.Exec("INSERT INTO client (fio, age, mail) VALUES ($1, $2, $3);", fio, age, mail)
	if err != nil {
		return err
	}

	return nil
}

// Delete - delete client with mail (PK)
func (c Client) Delete(mail string) error {
	result, err := c.db.Exec("DELETE FROM client WHERE mail=$1;", mail)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("deleted %d rows", rowsAffected)

	return nil
}

func (c Client) Update(mail string, fio string, age int8) error {
	return nil
}

func (c Client) Get(mail string) (*bisenes_logic.Client, error) {
	var client bisenes_logic.Client

	rows := c.db.QueryRow("SELECT mail, fio, age FROM client WHERE mail = $1", mail)
	err := rows.Scan(&client.Mail, &client.FIO, &client.Age)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
