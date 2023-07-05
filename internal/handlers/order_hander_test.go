package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	order := GetOrderTest()
	assert.NotNil(t, order)
	assert.NotEqual(t, 20.0, order.Value)
	assert.NotEmpty(t, order.IdOrder)
	assert.NotEmpty(t, order.IdCustumer)
}
