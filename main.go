package main

import (
	"fmt"
	"nilchanpub/feature1"
	"nilchanpub/feature2"
	"nilchanpub/feature3"
	"nilchanpub/feature_postgres/simple_connection"
)

func main() {
	feature1.Feature1()
	feature2.Feature2()
	feature3.Feature3()

	simple_connection.CheckConnection()
	fmt.Println("main end.")
}
