package main

import (
	"go-enigma-laundry/delivery"

	_ "github.com/lib/pq"
)

func main() {
	//run
	delivery.NewConsole().Run()

}
