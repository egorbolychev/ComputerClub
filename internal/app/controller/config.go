package controller

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Basic parameters of the club
type Config struct {
	MaxTables int
	TimeStart time.Time
	TimeEnd   time.Time
	Cost      int
}

func NewConfig() *Config {
	return &Config{}
}

// Parse and Validate data from string array to config struct.
// If data isn't valid - return wrong string as a error
func (conf *Config) ConfigureConfig(confStr []string) error {
	maxTables, err := strconv.Atoi(confStr[0])
	if err != nil {
		return errors.New(confStr[0])
	}
	conf.MaxTables = maxTables

	timeBounds := strings.Split(confStr[1], " ")
	if len(timeBounds) != 2 {
		return errors.New(confStr[1])
	}
	timeStart, err := time.Parse("15:04", timeBounds[0])
	if err != nil {
		return errors.New(confStr[1])
	}

	timeEnd, err := time.Parse("15:04", timeBounds[1])
	if err != nil {
		return errors.New(confStr[1])
	}
	if timeStart.After(timeEnd) {
		return errors.New(confStr[1])
	}

	conf.TimeStart = timeStart
	conf.TimeEnd = timeEnd

	cost, err := strconv.Atoi(confStr[2])
	if err != nil {
		return errors.New(confStr[2])
	}
	conf.Cost = cost

	return nil
}
