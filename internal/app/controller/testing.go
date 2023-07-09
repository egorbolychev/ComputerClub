package controller

import "time"

func TestConfig() *Config {
	timeStart, _ := time.Parse("15:04", "11:00")
	timeEnd, _ := time.Parse("15:04", "18:00")
	return &Config{
		MaxTables: 3,
		TimeStart: timeStart,
		TimeEnd:   timeEnd,
		Cost:      10,
	}
}
