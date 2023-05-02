package model

import "time"

type Question struct {
	ID      int
	Email   string
	Header  string
	Message string
	Answer  string
	Date    time.Time
	Status  string
}
