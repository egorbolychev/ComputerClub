package models

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Time     time.Time
	Id       int
	Username string
	TableNum int
}

func NewTask() *Task {
	return &Task{}
}

// Parse and validate task from string.
// If string isn't valid return that one
func (t *Task) ParseTask(taskStr string, maxTables int) error {
	s := strings.Split(taskStr, " ")
	lenS := len(s)
	if lenS > 4 || lenS < 3 {
		return errors.New(taskStr)
	}

	time, err := time.Parse("15:04", s[0])
	if err != nil {
		return errors.New(taskStr)
	}
	t.Time = time

	switch s[1] {
	case "1", "2", "3", "4":
	default:
		return errors.New(taskStr)
	}
	id, err := strconv.Atoi(s[1])
	if err != nil {
		return errors.New(taskStr)
	}
	t.Id = id

	if match, err := regexp.Match("(^[a-z0-9_-]{1,}$)", []byte(s[2])); err != nil || !match {
		return errors.New(taskStr)
	}
	t.Username = s[2]

	if lenS == 4 {
		tableNum, err := strconv.Atoi(s[3])
		if err != nil {
			return errors.New(taskStr)
		}
		t.TableNum = tableNum
	}

	if t.Id == 2 && (t.TableNum == 0 || t.TableNum > maxTables) {
		return errors.New(taskStr)
	}

	return nil
}

// Log task in stdout
func (t *Task) Log() {
	time := t.Time.Format("15:04")
	s := fmt.Sprintf("%s %d %s", time, t.Id, t.Username)
	if t.TableNum != 0 {
		s += fmt.Sprintf(" %d", t.TableNum)
	}
	log.Println(s)
}
