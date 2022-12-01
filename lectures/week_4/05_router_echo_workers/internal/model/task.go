package model

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Task struct {
	Name    string      `json:"name"`
	InParam int         `json:"inParam"`
	Log     echo.Logger `json:"-"`
}

func NewTask(log echo.Logger) *Task {
	return &Task{Log: log}
}

func (t *Task) Execute(timeMS int) {
	t.Log.Infof("task %s started for %d ms", t.Name, t.InParam)
	time.Sleep(time.Duration(timeMS) * time.Millisecond)
	t.Log.Infof("task %s stopped", t.Name)
}
