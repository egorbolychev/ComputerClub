package store

import (
	"sort"

	"github.com/egorbolychev/internal/app/models"
)

// Manage the clients.
// Data struct: hash table, key - username
type ClientManager struct {
	clients map[string]*models.Client
}

func (m *ClientManager) Add(user string) {
	c := models.NewClient()
	c.IsInside = true
	m.clients[user] = c
}

func (m *ClientManager) IsIn(user string) bool {
	return m.clients[user] != nil
}

func (m *ClientManager) GetUserTable(user string) int {
	return m.clients[user].TableNum
}

func (m *ClientManager) SatDown(user string, num int) {
	m.clients[user].TableNum = num
}

func (m *ClientManager) Remove(user string) {
	delete(m.clients, user)
}

func (m *ClientManager) GetListByAlthabet() []string {
	var result []string
	for key := range m.clients {
		result = append(result, key)
	}
	sort.Strings(result)

	return result
}
