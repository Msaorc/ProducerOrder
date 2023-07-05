package main

import (
	"encoding/json"

	"github.com/Msaorc/ProducerOrder/internal/handlers"
	"github.com/Msaorc/ProducerOrder/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel("")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		order := handlers.GetOrderTest()
		if order == nil {
			println("Order invalid!")
		}
		b, err := json.Marshal(order)
		if err != nil {
			panic(err)
		}
		err = rabbitmq.Producer(ch, string(b), "nameExtend")
		if err != nil {
			panic(err)
		}
	}
}
