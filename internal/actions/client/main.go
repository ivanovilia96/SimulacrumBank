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
	fmt.Print("action Client.Add called", mail, fio, age)
	return nil
}

func (c Client) Delete(mail string) error {
	return nil
}

func (c Client) Update(mail string, fio string, age int8) error {
	return nil
}

func (c Client) Get(mail string) (*bisenes_logic.Client, error) {
	return nil, nil
}
