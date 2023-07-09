package controller

import (
	"log"

	"github.com/egorbolychev/internal/app/models"
	"github.com/egorbolychev/internal/app/store"
)

type Controller struct {
	config *Config
	store  *store.Store
}

func NewController(config *Config, store *store.Store) *Controller {
	return &Controller{
		config: config,
		store:  store,
	}
}

func (c *Controller) ReadAndServe() {
	log.Println(c.config.TimeStart.Format("15:04"))

	for c.store.Task().GetLen() > 0 {
		task := c.store.Task().GetTask()
		if task.Time.After(c.config.TimeEnd) {
			break
		}
		task.Log()
		switch task.Id {
		case 1:
			c.ClientCame(task)
		case 2, 12:
			c.ClientSatDown(task)
		case 3:
			c.Wait(task)
		case 4:
			c.ClientGone(task)
		case 11:
			c.ClientGoneOver(task)
		}
	}
	c.OnExit()
}

func (c *Controller) ClientCame(task *models.Task) {
	if c.store.Client().IsIn(task.Username) {
		c.store.Task().GenerateTask(task.Time, "YouShallNotPass", 13, 0)
		return
	}
	if c.config.TimeStart.After(task.Time) || c.config.TimeEnd.Before(task.Time) {
		c.store.Task().GenerateTask(task.Time, "NotOpenYet", 13, 0)
		return
	}
	c.store.Client().Add(task.Username)
}

func (c *Controller) ClientSatDown(task *models.Task) {
	if !c.store.Client().IsIn(task.Username) {
		c.store.Task().GenerateTask(task.Time, "ClientUnknown", 13, 0)
		return
	}
	if c.store.Table().IsBusy(task.TableNum) {
		c.store.Task().GenerateTask(task.Time, "PlaceIsBusy", 13, 0)
		return
	}
	userTable := c.store.Client().GetUserTable(task.Username)
	if userTable != 0 {
		c.store.Table().ClearTable(userTable, c.config.Cost, task.Time)
	}
	c.store.Table().NewUser(task.TableNum, task.Username, task.Time)
	c.store.Client().SatDown(task.Username, task.TableNum)
}

func (c *Controller) Wait(task *models.Task) {
	if c.store.Table().HasEmptyTable() {
		c.store.Task().GenerateTask(task.Time, "ICanWaitNoLonger", 13, 0)
		return
	}
	if c.store.Queu().Len() > c.config.MaxTables {
		c.store.Task().GenerateTask(task.Time, task.Username, 11, 0)
		return
	}

	c.store.Queu().Add(task.Username)
}

func (c *Controller) ClientGone(task *models.Task) {
	if !c.store.Client().IsIn(task.Username) {
		c.store.Task().GenerateTask(task.Time, "ClientUnknown", 13, 0)
		return
	}
	tableNum := c.store.ClientManager.GetUserTable(task.Username)
	if tableNum != 0 {
		c.store.Table().ClearTable(tableNum, c.config.Cost, task.Time)
		if c.store.Queu().Len() > 0 {
			newSitter := c.store.Queu().Get()
			c.store.Task().GenerateTask(task.Time, newSitter, 12, tableNum)
		}
	}
	c.store.Client().Remove(task.Username)
}

func (c *Controller) ClientGoneOver(task *models.Task) {
	if !c.store.Client().IsIn(task.Username) {
		c.store.Task().GenerateTask(task.Time, "ClientUnknown", 13, 0)
		return
	}
	tableNum := c.store.ClientManager.GetUserTable(task.Username)
	if tableNum != 0 {
		c.store.Table().ClearTable(tableNum, c.config.Cost, task.Time)
	}
	c.store.Client().Remove(task.Username)
}

func (c *Controller) OnExit() {
	for _, user := range c.store.Client().GetListByAlthabet() {
		c.store.Task().GenerateTask(c.config.TimeEnd, user, 11, 0)
		task := c.store.Task().GetTask()
		task.Log()
		c.ClientGoneOver(task)
	}
	log.Println(c.config.TimeEnd.Format("15:04"))
	for i := 1; i <= c.config.MaxTables; i++ {
		money, time := c.store.Table().GetSum(i)
		log.Println(i, money, time.Format("15:04"))
	}
}
