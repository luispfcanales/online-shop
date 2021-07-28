package models

import "time"

type Order struct {
	IdOrder           int
	IdProduct         int
	IdPeople          int
	Amount            float64
	TotalAmount       float64
	ShopDate          time.Time
	DirectionDelivery string
}
