package main

import "time"

type Task struct {
    /* A struct for storing tasks data.
       If DoneAt is zero, task is not completed.

       Структура для хранения данных о задачах.
       Если поле DoneAt нулевое, значит задача не выполнена */
    ID        int
    Name      string
    DoneAt    time.Time
    CreatedAt time.Time
}

type TaskPost struct {
    // A struct for getting POST request data from tasks API
    // Структура для получения данных POST запроса из API задач
    Name string `json:"name"`
}

type TasksMgr struct {
    // A struct for managing tasks slice and counting IDs.
    // Структура для управления слайсом задач и счетом ID.
    Tasks *[]Task
    Count int
}

func (t *Task) Done() {
    /* Since Task.DoneAt is the only status property,
       we only need to update it to mark a task as done.

       Так как Task.DoneAt это единственное поле
       статуса задачи, нужно обновить только его, чтобы
       задача считалась выполненной. */
    t.DoneAt = time.Now()
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

    if t.DoneAt.IsZero() {
        t.Done()
    } else {
        t.DoneAt = time.Time{}
    }
}

func appendTask(mgr *TasksMgr, name string) {
    /* Adds a new task to in-memory tasks slice.
       Sets the CreatedAt property at UTC timestamp.
       Args:
         mgr *TasksMgr - a pointer to the task manager
         name   string  - a name for the new task

       Добавляет новую задачу в слайс задач в памяти.
       Устанавливает поле CreatedAt к текущему времени UTC.
       Аргументы:
         mgr *TasksMgr - указатель к менеджеру задач
         name   string  - имя новой задачи */
    loc, _ := time.LoadLocation("UTC")
    timestamp := time.Now().In(loc)
    (*mgr).Count++
    task := Task{
        ID:        (*mgr).Count,
        Name:      name,
        CreatedAt: timestamp,
    }
    *(*mgr).Tasks = append(*(*mgr).Tasks, task)
}
