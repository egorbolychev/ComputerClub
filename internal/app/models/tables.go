package models

import (
	"time"
)

type Table struct {
	Id        int
	User      string
	TimeStart time.Time
	AllTime   time.Time
	Money     int
}

func NewTable() *Table {
	return &Table{}
}
