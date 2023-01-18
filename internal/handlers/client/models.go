package client

import (
	"simulacrumBank/internal/adapters/bisenes_logic"
)

type Handlers struct {
	Actions bisenes_logic.ClientActions
}

type Client struct {
	Mail string `json:"mail"`
	Fio  string `json:"fio"`
	Age  int8   `json:"age"`
}
