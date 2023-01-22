package cash_account

func (b CashAccount) Create(mail, currency string) (int, error) {
	_, err := b.clientActions.Get(mail) // проверяем есть ли такой клиент что бы добавить для него счет
	if err != nil {
		return 0, err
	}

	lastInsertId := 0
	err = b.db.
		QueryRow("INSERT INTO cashaccount (currency) VALUES ($1) RETURNING id", currency).
		Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	_, err = b.db.Exec("INSERT INTO clientcashaccountlink(mail, id) VALUES ($1, $2)", mail, lastInsertId)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
