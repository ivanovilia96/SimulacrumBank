package cash_account

import (
	"simulacrumBank/internal/adapters/bisenes_logic"
	"simulacrumBank/internal/adapters/data_base"
)

func NewActions(db data_base.DataBaseActions, clientActions bisenes_logic.ClientActions) bisenes_logic.CashAccountActions {
	// сюда нужно передать бд чере адаптер\интерфейс
	return CashAccount{
		db:            db,
		clientActions: clientActions,
	}
}
