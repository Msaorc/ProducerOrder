package main

import "github.com/Msaorc/ProducerOrder/internal/handlers"

func main() {
	for i := 0; i < 5; i++ {
		order := handlers.GetOrderTest()
		if order == nil {
			println("Order invalid!")
		}
	}
}
