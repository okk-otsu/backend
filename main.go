package main

import (
	"fmt"
	"nilchanpub/feature1"
	"nilchanpub/feature_postgres/simple_connection"
)

func main() {
	feature1.Feature1()
	feature1.Feature1()

	simple_connection.CheckConnection()
	fmt.Println("main end.")
}
