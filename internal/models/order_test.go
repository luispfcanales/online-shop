package models

import (
	"fmt"
	"testing"
	"time"
)

const FORMAT_DATE = "2006-01-02"

func TestCreateOrder(t *testing.T) {
	currentDate := time.Now().Format(FORMAT_DATE)
	dateStamp, _ := time.Parse(FORMAT_DATE, currentDate)
	order := &Order{
		IdOrder:           1,
		IdProduct:         1,
		IdPeople:          1,
		Amount:            23.90,
		TotalAmount:       23.90,
		ShopDate:          dateStamp,
		DirectionDelivery: "jr.Acash 945",
	}
	dateRepository := fmt.Sprintf("%v", dateStamp.Format(FORMAT_DATE))

	if currentDate != dateRepository {
		t.Errorf("not equals date of Order: %v to %v", order.ShopDate, currentDate)
	}

}
