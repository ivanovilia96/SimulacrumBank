package cash_account

import (
	"simulacrumBank/internal/adapters/bisenes_logic"
)

type Handlers struct {
	Actions bisenes_logic.CashAccountActions
}

type CashAccount struct {
	id       int
	currency string
	count    int
}

type CreateCashAccountData struct {
	Mail, Currency string
}
