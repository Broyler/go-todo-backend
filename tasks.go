package main

import "time"

type Task struct {
	name      string
	doneAt    time.Time
	createdAt time.Time
}

func (t *Task) Done() {
	t.doneAt = time.Now()
}

func (t *Task) Toggle() {
	if t.doneAt.IsZero() {
		t.Done()
	} else {
		t.doneAt = time.Time{}
	}
}
