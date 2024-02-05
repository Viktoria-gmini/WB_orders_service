package sender

import (
	"encoding/json"
	"errors"
	"log"

	str "github.com/nats-io/go-nats-streaming/services/models"
)

func ValidOrder(data []byte, order str.Order) (str.Order, error) {

	err := json.Unmarshal(data, &order)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON:", err)
	}
	// Проверка длины полей в структуре Order
	if len(order.OrderUID) > 100 {
		return order, errors.New("orderUID is too long")
	}
	if len(order.TrackNumber) > 100 {
		return order, errors.New("trackNumber is too long")
	}
	if len(order.Entry) > 100 {
		return order, errors.New("entry is too long")
	}
	return order, err
}
