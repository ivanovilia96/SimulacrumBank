package data_base

//type Bank interface {
//	Client
//	BankAccount
//}

// сущьность описывающая действия которые мы можем сделать с клиентом банка
type ClientActions interface {
	Add(mail, fio string, age int8) error    // добавляем ( регистрируем ) клиента в банке
	Delete(mail string) error                // удаляем клиента и его банковские счета, вернуть ошибку если хоть на 1 банковском счете есть деньги
	Update(mail, fio string, age int8) error // позволяем изменить у клиента информацию если такой (маил) еще не существует в банке
	Get(mail string) (Client, error)
}

// описывает действия которые мы можем сделать с банковским аккаунтом клиента
type BankAccountActions interface {
	Create(mail, currency string) (int, error) // создаем аккаунт для клиента в нужной валюте, возвращает айди
	Delete(id int) error                       // удаляем по id
	AddMoney(id int, count float64) error      // добавляем деньги на счет с определенным номером
	Get(id int) (BankAccount, error)           // получаем информацию по номеру счета
	GetAll(mail string) ([]BankAccount, error) // получаем информацию о всех счетах для 1 пользователя банка
}

type BankAccount struct {
	Id       int
	Currency string
	Count    float64
}

type Client struct {
	Mail string
	FIO  string
	Age  int
}
