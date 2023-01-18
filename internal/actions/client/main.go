package client

import (
	"simulacrumBank/internal/adapters/bisenes_logic"
	"simulacrumBank/internal/adapters/data_base"
)

func NewActions(db data_base.DataBaseActions) bisenes_logic.ClientActions {
	// сюда нужно передать бд чере адаптер\интерфейс
	return Client{
		db: db,
	}
}
