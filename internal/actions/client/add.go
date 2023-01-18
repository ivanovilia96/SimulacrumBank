package client

func (c Client) Add(mail, fio string, age int8) error {
	_, err := c.db.Exec("INSERT INTO client (fio, age, mail) VALUES ($1, $2, $3);", fio, age, mail)
	if err != nil {
		return err
	}

	return nil
}
