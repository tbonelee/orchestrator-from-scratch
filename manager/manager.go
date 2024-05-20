package manager

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"orchestrator-from-scratch/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Selecting a worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Updating tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("Sending work to workers")
}
