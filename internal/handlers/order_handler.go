package handlers

import (
	"math/rand"

	"github.com/Msaorc/ProducerOrder/internal/entity"
)

func GetOrderTest() *entity.Order {
	return entity.NewOrder(rand.Float64())
}
