package store

import (
	"fmt"
	"orchestrator-from-scratch/task"
)

type Store interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, error)
	List() (interface{}, error)
	Count() (int, error)
}

type InMemoryTaskStore struct {
	Db map[string]*task.Task
}

func NewInMemoryTaskStore() *InMemoryTaskStore {
	return &InMemoryTaskStore{
		Db: make(map[string]*task.Task),
	}
}

func (i *InMemoryTaskStore) Put(key string, value interface{}) error {
	t, ok := value.(*task.Task)
	if !ok {
		return fmt.Errorf("value %v is not a task.Task type", value)
	}
	i.Db[key] = t
	return nil
}

func (i *InMemoryTaskStore) Get(key string) (interface{}, error) {
	t, ok := i.Db[key]
	if !ok {
		return nil, fmt.Errorf("task with key %s does not exist", key)
	}
	return t, nil
}

func (i *InMemoryTaskStore) List() (interface{}, error) {
	var tasks []*task.Task
	for _, t := range i.Db {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (i *InMemoryTaskStore) Count() (int, error) {
	return len(i.Db), nil
}

type InMemoryTaskEventStore struct {
	Db map[string]*task.TaskEvent
}

func NewInMemoryTaskEventStore() *InMemoryTaskEventStore {
	return &InMemoryTaskEventStore{
		Db: make(map[string]*task.TaskEvent),
	}
}

func (i *InMemoryTaskEventStore) Put(key string, value interface{}) error {
	te, ok := value.(*task.TaskEvent)
	if !ok {
		return fmt.Errorf("value %v is not a task.TaskEvent type", value)
	}
	i.Db[key] = te
	return nil
}

func (i *InMemoryTaskEventStore) Get(key string) (interface{}, error) {
	te, ok := i.Db[key]
	if !ok {
		return nil, fmt.Errorf("task event with key %s does not exist", key)
	}
	return te, nil
}

func (i *InMemoryTaskEventStore) List() (interface{}, error) {
	var taskEvents []*task.TaskEvent
	for _, te := range i.Db {
		taskEvents = append(taskEvents, te)
	}
	return taskEvents, nil
}

func (i *InMemoryTaskEventStore) Count() (int, error) {
	return len(i.Db), nil
}
