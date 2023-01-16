package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db_adapter "simulacrumBank/internal/adapters/data_base"
	db "simulacrumBank/internal/data_base"
	"simulacrumBank/internal/rest"
)

func main() {
	// что я хочу,
	//1 - блок где мы берем статические данные ( бд\файлы) типо " мы хотим получить пользоавателей. откуда мы их получим не важно, файл там или бд или рест"
	// 2 - блок бизнес логики типо "проверить есть ли у пользователя Х денег\ перевод от 1 пользователя к другому"
	// 3 выход из МС - где мы будем получать что надо делать в метаданных
	e := gin.Default()

	dataBase, err := db.New("postgres", "postgres", "100", "localhost", "bank")
	if err != nil {
		panic(err)
	}

	dbActions := db_adapter.New(dataBase)
	rest.InitRouters(e, dbActions)

	err = e.Run()
	if err != nil {
		panic(err)
	}
}
