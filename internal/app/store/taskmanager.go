package store

import (
	"sort"
	"time"

	"github.com/egorbolychev/internal/app/models"
)

type TaskManager struct {
	taskStack models.TaskSlice
}

func (t *TaskManager) AddTask(task *models.Task) {
	t.taskStack = append(t.taskStack, task)
}

func (t *TaskManager) GenerateTask(time time.Time, username string, id, tableNum int) {
	task := models.NewTask()
	task.Id = id
	task.Time = time
	task.Username = username
	if tableNum != 0 {
		task.TableNum = tableNum
	}
	t.AddTask(task)
}

func (t *TaskManager) GetTask() *models.Task {
	tasksLen := len(t.taskStack) - 1
	task := t.taskStack[tasksLen]
	t.taskStack = t.taskStack[:tasksLen]

	return task
}

func (t *TaskManager) GetLen() int {
	return t.taskStack.Len()
}

func (t *TaskManager) Reverse() {
	sort.Sort(sort.Reverse(t.taskStack))
}
