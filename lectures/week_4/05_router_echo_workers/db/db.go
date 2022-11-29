package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type (
	dataMap[T any] map[int64]*T
	JsonDB[T any]  struct {
		lastID   int64
		filename string
		data     dataMap[T]
	}
)

func (d *JsonDB[T]) InsertInto(m *T, id int64, update *int64) error {
	if id <= d.lastID {
		id = d.getLastID()
	} else if id > d.lastID {
		d.lastID = id
	}
	*update = id
	d.data[id] = m
	data, err := json.Marshal(d.data)
	if err != nil {
		return err
	}
	return os.WriteFile(d.filename, data, 0o644)
}

func (d *JsonDB[T]) Records() int {
	return len(d.data)
}

func (d *JsonDB[T]) GetAll() []T {
	out := make([]T, 0, len(d.data))
	for _, v := range d.data {
		out = append(out, *v)
	}
	return out
}

func (d *JsonDB[T]) Select(number int64, m *T) error {
	data, ok := d.data[number]
	if !ok {
		return errors.New("not found")
	}
	if m == nil {
		return errors.New("input model is nil")
	}
	*m = *data
	return nil
}

func (d *JsonDB[T]) getLastID() int64 {
	if len(d.data) == 0 {
		d.lastID = 1
		return d.lastID
	}
	d.findLastId()
	d.lastID++
	return d.lastID
}

func (d *JsonDB[T]) findLastId() {
	if d.lastID == 0 {
		for idx := range d.data {
			if idx > d.lastID {
				d.lastID = idx
			}
		}
	}
}

func NewJsonDB[T any](filename string) *JsonDB[T] {
	members := make(dataMap[T])
	if fileData, err := os.ReadFile(filename); err == nil {
		err := json.Unmarshal(fileData, &members)
		if err != nil {
			log.Printf("can't unmarshal data from %s: %v", filename, err)
		}
	}
	tmp := JsonDB[T]{filename: filename, data: members}
	tmp.findLastId()
	return &tmp
}
