package main

import (
	"context"
	"fmt"
	simple_connection "nilchanpub/featurepostgres/simpleconnection"
	simple_sql "nilchanpub/featurepostgres/simplesql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		// log.Println("check connection error:", err)
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("Таблица в бд была создана успешно!")

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("SelectRows отработал успешно!")

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Покормить кошку"
			task.Description = "купи корм Kitekat"
			task.Completed = true
			t := time.Now()
			task.CompletedAt = &t
			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("main end.")
}
