package tasks

import (
	"sync"
	"time"
)

type Task struct {
	/* A struct for storing tasks data.
	   Структура для хранения данных о задачах. */
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskMgr struct {
	// A struct for managing tasks slice and counting IDs.
	// Структура для управления слайсом задач и счетом ID.
	Tasks []Task
	Count int
	sync.RWMutex
}

func (m *TaskMgr) AddTaskByName(name string) Task {
	return m.AddTask(Task{Name: name})
}

func (m *TaskMgr) AddTask(task Task) Task {
	/* Adds a new task to in-memory tasks slice.
	   Sets the CreatedAt property at the current timestamp.
	   Increments the task manager count.
	   Args:
	     task Task - a Task object with necessary params

	   Добавляет новую задачу в слайс задач в памяти.
	   Устанавливает поле CreatedAt к текущему времени.
	   Увеличивает счетчик менеджера задач.
	   Аргументы:
	     name   string  - имя новой задачи */
	m.Lock()
	defer m.Unlock()

	m.Count++
	aux := Task{
		ID:        m.Count,
		Name:      task.Name,
		CreatedAt: time.Now(),
		Done:      task.Done,
	}
	m.Tasks = append(m.Tasks, aux)
	return aux
}

func (m *TaskMgr) GetTasks() []Task {
	/* A method for getting tasks from TaskMgr's in-memory storage.
	   Метод для получения задач из хранилища менеджера задач. */
	m.RLock()
	defer m.RUnlock()
	return m.Tasks
}

func (m *TaskMgr) PutTask(data Task) Task {
	/* A method for PUT request, updating or creating a task.
	   Метод для PUT запроса, обновление или создание задачи. */
	m.Lock()

	for i, task := range m.Tasks {
		if task.ID == data.ID {
			m.Tasks[i].Name = data.Name
			m.Tasks[i].Done = data.Done
			m.Unlock()
			return m.Tasks[i]
		}
	}

	// Task doesn't exist, create new - задачи не существует, создание новой
	m.Unlock()
	return m.AddTask(data)
}
