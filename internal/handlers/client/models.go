package client

type Client struct {
	Mail string `json:"mail"`
	Fio  string `json:"fio"`
	Age  int8   `json:"age"`
}
