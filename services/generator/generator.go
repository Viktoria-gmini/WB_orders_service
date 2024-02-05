package generator

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	OrderUID        string   `json:"order_uid"`
	TrackNumber     string   `json:"track_number"`
	Entry           string   `json:"entry"`
	Delivery        Delivery `json:"delivery"`
	Payment         Payment  `json:"payment"`
	Items           []Item   `json:"items"`
	Locale          string   `json:"locale"`
	InternalSig     string   `json:"internal_signature"`
	CustomerID      string   `json:"customer_id"`
	DeliveryService string   `json:"delivery_service"`
	Shardkey        string   `json:"shardkey"`
	SMID            int      `json:"sm_id"`
	DateCreated     string   `json:"date_created"`
	OofShard        string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	ZIP     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int64  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int64  `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func GenerateJSON() ([]byte, error) {
	order := Order{
		OrderUID:        generateMD5(time.Now().String()),
		TrackNumber:     generateMD5(time.Now().String()),
		Entry:           generateMD5(time.Now().String()),
		Delivery:        Delivery{Name: generateMD5(time.Now().String())},
		Payment:         Payment{Transaction: generateMD5(time.Now().String())},
		Items:           []Item{{ChrtID: time.Now().UnixNano()}},
		Locale:          generateMD5(time.Now().String())[:10],
		InternalSig:     generateMD5(time.Now().String()),
		CustomerID:      generateMD5(time.Now().String()),
		DeliveryService: generateMD5(time.Now().String()),
		Shardkey:        generateMD5(time.Now().String())[:10],
		SMID:            time.Now().Nanosecond(),
		DateCreated:     time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		OofShard:        generateMD5(time.Now().String())[:10],
	}

	jsonData, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func generateMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}
