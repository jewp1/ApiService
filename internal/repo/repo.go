package repo

import (
	"errors"
	"sync"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type inMemoryCache struct {
	currentId int
	items     map[int]*Task
	mu        sync.RWMutex
}

type Repository interface {
	CreateTask(task Task) (int, error)
	GetTasks() (map[int]*Task, error)
	GetTaskById(id int) (Task, error)
	UpdateTask(id int, task Task) (int, error)
	DeleteTask(id int) (int, error)
}

func NewRepo() (Repository, error) {
	var id int
	return &inMemoryCache{items: make(map[int]*Task), currentId: id}, nil
}

func (cache *inMemoryCache) CreateTask(task Task) (int, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.currentId++
	_, ok := cache.items[cache.currentId]
	if ok {
		return cache.currentId, errors.New("task already exists")
	}
	cache.items[cache.currentId] = &task
	return cache.currentId, nil
}
func (cache *inMemoryCache) GetTasks() (map[int]*Task, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	if len(cache.items) == 0 {
		return nil, errors.New("tasks not found")
	}
	return cache.items, nil
}

func (cache *inMemoryCache) GetTaskById(id int) (Task, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	if task, ok := cache.items[id]; ok {
		return *task, nil
	}
	return Task{}, errors.New("task not found")
}
func (cache *inMemoryCache) UpdateTask(id int, newTask Task) (int, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.items[id]; !ok {
		return 0, errors.New("task not found")
	}
	cache.items[id] = &newTask
	return id, nil
}
func (cache *inMemoryCache) DeleteTask(id int) (int, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.items[id]; !ok {
		return 0, errors.New("task not found")
	}
	delete(cache.items, id)
	return 0, nil
}
