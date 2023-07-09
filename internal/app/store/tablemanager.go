package store

import (
	"time"

	"github.com/egorbolychev/internal/app/models"
)

type TableManager struct {
	tables map[int]*models.Table
}

func (m *TableManager) AddTable(id int) {
	table := models.NewTable()
	table.Id = id
	m.tables[id] = table
}

func (m *TableManager) NewUser(num int, user string, startTime time.Time) {
	t := m.tables[num]
	t.User = user
	t.TimeStart = startTime
}

func (m *TableManager) IsBusy(num int) bool {
	return m.tables[num].User != ""
}

func (m *TableManager) HasEmptyTable() bool {
	for _, val := range m.tables {
		if val.User == "" {
			return true
		}
	}

	return false
}

func (m *TableManager) ClearTable(num, hourCost int, endTime time.Time) {
	table := m.tables[num]
	usageTime := endTime.Sub(table.TimeStart)
	table.AllTime = table.AllTime.Add(usageTime)
	cost := int(usageTime.Minutes()/60) * hourCost
	if int(usageTime.Minutes())%60 > 0 {
		cost += hourCost
	}
	table.Money += cost
	table.User = ""
	table.TimeStart = time.Time{}
}

func (m *TableManager) Len() int {
	return len(m.tables)
}

func (m *TableManager) GetSum(num int) (int, time.Time) {
	table := m.tables[num]
	return table.Money, table.AllTime
}
