package client

import "simulacrumBank/internal/adapters/data_base"

type Client struct {
	db data_base.DataBaseActions
}
