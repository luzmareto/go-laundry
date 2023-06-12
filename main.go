package main

import (
	"go-laundry/delivery"

	_ "github.com/lib/pq"
)

func main() {
	//run
	delivery.NewConsole().Run()

}
