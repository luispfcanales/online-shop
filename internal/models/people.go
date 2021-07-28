package models

import "time"

type People struct {
	IdPeople    int
	IdUser      int
	Dni         string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
