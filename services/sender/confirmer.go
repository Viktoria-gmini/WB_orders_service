package sender

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
	str "github.com/nats-io/go-nats-streaming/services/models"
)

func Confirm(data []byte, cache *cacher.Cache) error {
	var order str.Order
	var err error
	err = json.Unmarshal(data, &order)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON:", err)
	}
	connStr := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	err = InsertOrder(db, order)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = cache.Set(order.OrderUID, data, 0)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
