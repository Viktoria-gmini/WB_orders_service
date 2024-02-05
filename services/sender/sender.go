package sender

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"

	str "github.com/nats-io/go-nats-streaming/services/models"
)

func InsertOrder(db *sql.DB, order str.Order) error {
	query := `INSERT INTO orders (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES($1, $2, $3, $4::jsonb, $5::jsonb, $6::jsonb[], $7, $8, $9, $10, $11, $12, $13, $14)`
	var err error
	deliveryJSON, err := json.Marshal(order.Delivery)
	if err != nil {
		return err
	}
	paymentJSON, err := json.Marshal(order.Payment)
	if err != nil {
		return err
	}
	var itemsArray []string
	for _, item := range order.Items {
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return err
		}
		// Экранируем двойные кавычки внутри JSON строки.
		itemJSONStr := strconv.Quote(string(itemJSON))
		itemsArray = append(itemsArray, itemJSONStr)
	}

	// Соединяем все JSON строки в один литерал массива для PostgreSQL.
	itemsArrayLiteral := "{" + strings.Join(itemsArray, ",") + "}"
	_, err = db.Exec(query, order.OrderUID, order.TrackNumber, order.Entry, deliveryJSON, paymentJSON, itemsArrayLiteral, order.Locale, order.InternalSig, order.CustomerID, order.DeliveryService, order.Shardkey, order.SMID, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	return nil
}
