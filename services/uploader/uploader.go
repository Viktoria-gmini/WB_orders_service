package uploader

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	str "github.com/nats-io/go-nats-streaming/services/models"
)

func Upload() []byte {
	connStr := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data, err := parseRowsToJSON(rows, db)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
func parseRowsToJSON(rows *sql.Rows, db *sql.DB) ([]byte, error) {
	var orders []str.Order

	// Итерируемся по строкам
	for rows.Next() {
		var order str.Order
		var delivery []byte
		var payment []byte
		var items []byte
		err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&delivery,
			&payment,
			&items,
			&order.Locale,
			&order.InternalSig,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SMID,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			log.Fatal("Failed scanning:", err)
		}
		err = json.Unmarshal(delivery, &order.Delivery)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON:", err)
		}
		err = json.Unmarshal(payment, &order.Payment)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON:", err)
		}
		order.Items, err = itemsFromLiteral(items)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON:", err)
		}

		orders = append(orders, order)
	}

	// Преобразуем список заказов в json
	jsonData, err := json.Marshal(orders)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func itemsFromLiteral(itemsArrayLiteral []byte) ([]str.Item, error) {
	var itemsArray []str.Item
	trimmed := strings.Trim(string(itemsArrayLiteral), "{}")
	unquoted, err := strconv.Unquote(trimmed)
	if err != nil {
		return nil, err
	}
	arrayed := "[" + unquoted + "]"
	err = json.Unmarshal([]byte(arrayed), &itemsArray)
	if err != nil {
		return nil, err
	}
	return itemsArray, nil
}
