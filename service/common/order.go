package common

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model        `json:"-"`
	OrderUID          string    `json:"order_uid,required" validate:"required" fake:"{uuid}"`
	TrackNumber       string    `json:"track_number" validate:"required" fake:"{regex:[A-Z]{16}}"`
	Entry             string    `json:"entry" validate:"required" fake:"{regex:[A-Z]{9}}"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required" fake:"{language}"`
	InternalSignature string    `json:"internal_signature" validator:"required"`
	CustomerID        string    `json:"customer_id" validate:"required" fake:"{uuid}"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required" fake:"{uuid}"`
	SmID              int       `json:"sm_id" validate:"required" fake:"{number:1,100000}"`
	DateCreated       time.Time `json:"date_created" validate:"required" fake:"{date}"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	ID      int    `json:"-"`
	Name    string `json:"name" validate:"required" fake:"{name}"`
	Phone   string `json:"phone" validate:"required" fake:"{phone}"`
	Zip     string `json:"zip" validate:"required" fake:"{zip}"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
	OrderID int    `json:"-"`
}

type Payment struct {
	ID           int    `json:"-"`
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee" validate:"required"`
	OrderID      int    `json:"-"`
}

type Item struct {
	ID          int    `json:"-"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
	OrderID     int    `json:"-"`
}
