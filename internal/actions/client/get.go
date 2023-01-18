package client

import "simulacrumBank/internal/adapters/bisenes_logic"

func (c Client) Get(mail string) (*bisenes_logic.Client, error) {
	var client bisenes_logic.Client

	rows := c.db.QueryRow("SELECT mail, fio, age FROM client WHERE mail = $1", mail)
	err := rows.Scan(&client.Mail, &client.FIO, &client.Age)
	if err != nil {
		return nil, err
	}

	return &client, nil
}
