package main

import (
	"sync"
	"time"
)

type Task struct {
	/* A struct for storing tasks data.
	   If DoneAt is zero, task is not completed.

	   Структура для хранения данных о задачах.
	   Если поле DoneAt нулевое, значит задача не выполнена */
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Done      bool       `json:"done"`
	DoneAt    *time.Time `json:"done_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

type TaskPost struct {
	// A struct for getting POST request data from tasks API
	// Структура для получения данных POST запроса из API задач
	Name string `json:"name"`
}

type TasksMgr struct {
	// A struct for managing tasks slice and counting IDs.
	// Структура для управления слайсом задач и счетом ID.
	Tasks []Task
	Count int
	sync.RWMutex
}

func (t *Task) MarkDone() {
	/* Update Task.DoneAt and Task.Done
	   Обновить Task.DoneAt и Task.Done */
	now := time.Now()
	t.DoneAt = &now
	t.Done = true
}

func (t *Task) Toggle() {
	/* A method to toggle task.
	   Since Task.DoneAt is a Time object, not a boolean
	   a Task is not done if its DoneAt is zero.
	   Un-checking a task is setting DoneAt to zero.

	   Метод для переключения состояния задачи
	   Так как Task.DoneAt это объект Time, не булевой,
	   задача не считается сделанной пока его DoneAt нулевое.
	   Выключение задачи устанавливает DoneAt к нулю. */

	if t.DoneAt == nil {
		t.MarkDone()
	} else {
		t.DoneAt = nil
		t.Done = false
	}
}

func (m *TasksMgr) AddTask(name string) Task {
	/* Adds a new task to in-memory tasks slice.
	   Sets the CreatedAt property at UTC timestamp.
	   Increments the task manager count.
	   Args:
	     name   string  - a name for the new task

	   Добавляет новую задачу в слайс задач в памяти.
	   Устанавливает поле CreatedAt к текущему времени UTC.
	   Увеличивает счетчик менеджера задач.
	   Аргументы:
	     name   string  - имя новой задачи */
	m.Lock()
	defer m.Unlock()

	m.Count++
	task := Task{
		ID:        m.Count,
		Name:      name,
		CreatedAt: time.Now(),
	}
	m.Tasks = append(m.Tasks, task)
	return task
}

func (m *TasksMgr) GetTasks() []Task {
	/* A method for getting tasks from TaskMgr's in-memory storage.
	   Метод для получения задач из хранилища менеджера задач. */
	m.RLock()
	defer m.RUnlock()
	return m.Tasks
}
