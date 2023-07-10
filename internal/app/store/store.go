package store

import (
	"errors"

	"github.com/egorbolychev/internal/app/models"
)

// A structure that stores all data during the application's operation
type Store struct {
	TaskManager   *TaskManager
	TableManager  *TableManager
	QueueManager  *QueueManager
	ClientManager *ClientManager
}

func New() *Store {
	return &Store{}
}

// Configure Task stack and tables map from input
func (s *Store) ConfigureStore(tasksStr []string, maxTables int) error {
	var newTask *models.Task
	var prev models.Task

	for i := 0; i < maxTables; i++ {
		s.Table().AddTable(i + 1)
	}
	for _, task := range tasksStr {
		newTask = models.NewTask()
		err := newTask.ParseTask(task, maxTables)
		if err != nil {
			return err
		}
		s.Task().AddTask(newTask)

		//проверяется, что все события последовательны
		if !prev.Time.IsZero() && prev.Time.After(newTask.Time) {
			return errors.New(task)
		}
		prev = *newTask
	}
	s.Task().Reverse()

	return nil
}

func (s *Store) Task() *TaskManager {
	if s.TaskManager != nil {
		return s.TaskManager
	}

	s.TaskManager = &TaskManager{
		taskStack: make([]*models.Task, 0),
	}

	return s.TaskManager
}

func (s *Store) Table() *TableManager {
	if s.TableManager != nil {
		return s.TableManager
	}

	s.TableManager = &TableManager{
		tables: make(map[int]*models.Table),
	}

	return s.TableManager
}

func (s *Store) Queu() *QueueManager {
	if s.QueueManager != nil {
		return s.QueueManager
	}

	s.QueueManager = &QueueManager{
		queu: make([]string, 0),
	}

	return s.QueueManager
}

func (s *Store) Client() *ClientManager {
	if s.ClientManager != nil {
		return s.ClientManager
	}

	s.ClientManager = &ClientManager{
		clients: make(map[string]*models.Client),
	}

	return s.ClientManager
}
