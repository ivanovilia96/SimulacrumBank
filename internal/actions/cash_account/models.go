package cash_account

import (
	"simulacrumBank/internal/adapters/bisenes_logic"
	"simulacrumBank/internal/adapters/data_base"
)

type CashAccount struct {
	db            data_base.DataBaseActions
	clientActions bisenes_logic.ClientActions
}
