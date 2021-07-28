package models

import "time"

type Product struct {
	IdProduct    int
	NameProduct  string
	Description  string
	price        float64
	stock        int
	DateRegister time.Time
}
