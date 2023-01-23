package cash_account

func (b CashAccount) Create(mail, currency string) (int, error) {
	_, err := b.clientActions.Get(mail) // проверяем есть ли такой клиент что бы добавить для него счет
	if err != nil {
		return 0, err
	}

	lastInsertId := 0
	err = b.db.QueryRow("INSERT INTO clientcashaccountlink(mail) VALUES ($1) RETURNING cash_account_id", mail).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	_, err = b.db.
		Query("INSERT INTO cashaccount (id,currency) VALUES ($1,$2)", lastInsertId, currency)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
