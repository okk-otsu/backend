package main

import (
	"fmt"
	"nilchanpub/httpserver"
)

func main() {
	fmt.Println("Запуск http сервера!")
	fmt.Println("Новый вывод!")
	err := httpserver.StartHTTPServer()
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Сервер завершился успешно")
	}
	fmt.Println("main end.")
}
