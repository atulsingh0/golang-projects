package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CompletedAt time.Time
	CreatedAt   time.Time
}

type List []item

// Adding a task
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete Task
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exists", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// Delete Task
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exists", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save Task
func (l *List) Save(file string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, js, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Get Task
func (l *List) Load(file string) error {
	data, err := os.ReadFile(file)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exists", file)
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, l)
}

// Display
func (l *List) Display() (string, error) {
	js, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	return string(js), nil
}
