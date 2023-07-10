package controller

import (
	"testing"
	"time"

	"github.com/egorbolychev/internal/app/store"
)

func TestController_ClientCame(t *testing.T) {
	st := store.New()
	conf := TestConfig()
	c := NewController(conf, st)
	tasktime, _ := time.Parse("15:04", "12:00")
	c.store.Task().GenerateTask(tasktime, "user", 1, 0)
	task := c.store.Task().GetTask()
	c.ClientCame(task)
	if c.store.Task().GetLen() > 0 {
		t.Fail()
	}
	if !c.store.Client().IsIn("user") {
		t.Fail()
	}

	c.store.Task().GenerateTask(tasktime, "user", 1, 0)
	task = c.store.Task().GetTask()
	c.ClientCame(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	errTask := c.store.Task().GetTask()
	if errTask.Username != "YouShallNotPass" {
		t.Fail()
	}

	tasktime, _ = time.Parse("15:04", "09:00")
	c.store.Task().GenerateTask(tasktime, "user1", 1, 0)
	task = c.store.Task().GetTask()
	c.ClientCame(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	errTask = c.store.Task().GetTask()
	if errTask.Username != "NotOpenYet" {
		t.Fail()
	}
}

func TestController_SatDown(t *testing.T) {
	st := store.New()
	conf := TestConfig()
	c := NewController(conf, st)
	tasktime, _ := time.Parse("15:04", "12:00")
	c.store.Task().GenerateTask(tasktime, "user", 2, 1)
	task := c.store.Task().GetTask()
	c.ClientSatDown(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	if errTask := c.store.Task().GetTask(); errTask.Username != "ClientUnknown" {
		t.Fail()
	}

	c.store.Table().AddTable(1)
	c.store.Task().GenerateTask(tasktime, "user", 1, 0)
	task = c.store.Task().GetTask()
	c.ClientCame(task)

	c.store.Task().GenerateTask(tasktime, "user", 2, 1)
	task = c.store.Task().GetTask()
	c.ClientSatDown(task)

	if c.store.Task().GetLen() > 0 {
		t.Fail()
	}

	c.store.Task().GenerateTask(tasktime, "user", 2, 1)
	task = c.store.Task().GetTask()
	c.ClientSatDown(task)

	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	if errTask := c.store.Task().GetTask(); errTask.Username != "PlaceIsBusy" {
		t.Fail()
	}

	c.store.Table().AddTable(2)
	c.store.Task().GenerateTask(tasktime, "user", 2, 2)
	task = c.store.Task().GetTask()
	c.ClientSatDown(task)

	if c.store.Table().IsBusy(1) {
		t.Fail()
	}
	if c.store.Task().GetLen() > 0 {
		t.Fail()
	}
}

func TestController_Wait(t *testing.T) {
	st := store.New()
	conf := TestConfig()
	conf.MaxTables = 1
	c := NewController(conf, st)
	c.store.Table().AddTable(1)

	tasktime, _ := time.Parse("15:04", "12:00")
	c.store.Task().GenerateTask(tasktime, "user", 3, 1)
	task := c.store.Task().GetTask()
	c.Wait(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	if errTask := c.store.Task().GetTask(); errTask.Username != "ICanWaitNoLonger" {
		t.Fail()
	}

	c.store.Table().NewUser(1, "user1", task.Time)

	c.store.Task().GenerateTask(tasktime, "user", 3, 1)
	task = c.store.Task().GetTask()
	c.Wait(task)
	if c.store.Task().GetLen() > 0 {
		t.Fail()
	}

	c.store.Task().GenerateTask(tasktime, "user1", 3, 1)
	task = c.store.Task().GetTask()
	c.Wait(task)
	c.store.Task().GenerateTask(tasktime, "user2", 3, 1)
	task = c.store.Task().GetTask()
	c.Wait(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}

}

func TestController_ClientGone(t *testing.T) {
	st := store.New()
	conf := TestConfig()
	c := NewController(conf, st)
	c.store.Table().AddTable(1)
	tasktime, _ := time.Parse("15:04", "12:00")
	c.store.Task().GenerateTask(tasktime, "user1", 4, 0)
	task := c.store.Task().GetTask()
	c.ClientGone(task)
	if c.store.Task().GetLen() != 1 {
		t.Fail()
	}
	if errTask := c.store.Task().GetTask(); errTask.Username != "ClientUnknown" {
		t.Fail()
	}
	c.store.Task().GenerateTask(tasktime, "user1", 1, 0)
	task = c.store.Task().GetTask()
	c.ClientCame(task)
	c.store.Task().GenerateTask(tasktime, "user1", 2, 1)
	task = c.store.Task().GetTask()
	c.ClientSatDown(task)
	c.store.Task().GenerateTask(tasktime, "user2", 1, 0)
	task = c.store.Task().GetTask()
	c.ClientCame(task)
	c.store.Task().GenerateTask(tasktime, "user2", 3, 0)
	task = c.store.Task().GetTask()
	c.Wait(task)
	if c.store.Queu().Len() != 1 {
		t.Fail()
	}
	endTime, _ := time.Parse("15:04", "13:00")
	c.store.Task().GenerateTask(endTime, "user1", 4, 0)
	task = c.store.Task().GetTask()
	c.ClientGone(task)
	if c.store.Client().IsIn("user1") {
		t.Fail()
	}
	if money, time := c.store.Table().GetSum(1); money != 10 || time.Format("15:04") != "01:00" {
		t.Fail()
	}
	if c.store.Queu().Len() > 0 {
		t.Fail()
	}
	task = c.store.Task().GetTask()
	c.ClientSatDown(task)
	if c.store.Client().GetUserTable("user2") != 1 {
		t.Fail()
	}
	if !c.store.Table().IsBusy(1) {
		t.Fail()
	}

}
