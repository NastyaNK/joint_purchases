package main

import (
	"fmt"
	"log"
	"net/http"

	"mvp/repository"
	"mvp/router"
)

/*
добавить мигратор
добавить миграции
добавить докер файл и компос файл
сделать редми файл с описанием запуска и зависимостями(указать что за мигратор испотзуется)
*/
const (
	userDB = "postgres"
	// pwdDB  = "sbermarket_paas"
	pwdDB  = "postgrespw"
	nameDB = "mvp"
)

func main() {
	db, err := repository.New(userDB, pwdDB, nameDB)
	if err != nil {
		log.Fatal("no create connect to db", err)
	}

	r := router.New(db)

	fmt.Println("http://localhost:3333/mvp")
	if err := http.ListenAndServe(":3333", r); err != nil {
		log.Fatal("start server err: ", err)
	}
}
