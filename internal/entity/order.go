package entity

import "github.com/Msaorc/ProducerOrder/pkg/entity"

type Order struct {
	IdOrder entity.ID
	IdCustumer entity.ID
	Value float64
}

func NewOrder(value float64) *Order {
	return &Order{
		IdOrder: entity.NewID(),
		IdCustumer: entity.NewID(),
		Value: value,
	}
}