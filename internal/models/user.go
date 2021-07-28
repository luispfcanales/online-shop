package models

import "time"

type User struct {
	Id                    int
	email                 string
	pass                  string
	rol                   int
	current_date_register time.Time
}
