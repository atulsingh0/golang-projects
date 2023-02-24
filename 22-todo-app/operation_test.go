package todo_test

import (
	"os"
	"testing"

	todo "github.com/atulsingh0/golang-projects/22-todo-app"
)

func TestAdd(t *testing.T) {

	ls := todo.List{}
	note := "One task"
	ls.Add(note)

	if ls[0].Task != note {
		t.Errorf("Expected %s, Got %s", note, ls[0].Task)
	}
}

func TestComplete(t *testing.T) {
	ls := todo.List{}
	note := "Complete the task"
	ls.Add(note)

	if ls[0].Task != note {
		t.Errorf("Task: %s not added to list", note)
	}

	if ls[0].Done {
		t.Errorf("Task completion status should not be %v", ls[0].Done)
	}

	err := ls.Complete(1) // Item #, It is not a index
	if err != nil {
		t.Errorf("err: %v", err.Error())
	}

	if ls[0].Task != note {
		t.Errorf("Task: %s not added to list", note)
	}

	if !ls[0].Done {
		t.Errorf("Expected %v, Got %v", true, ls[0].Done)
	}
}

func TestDelete(t *testing.T) {
	ls := todo.List{}
	notes := []string{"Deleted the task #1", "Deleted the task #2", "Deleted the task #3"}

	for _, note := range notes {
		ls.Add(note)
	}

	err := ls.Delete(2) // Item #, It is not a index
	if err != nil {
		t.Errorf("err: %v", err.Error())
	}

	if len(ls) > 2 {
		t.Errorf("Expected count was 2 but got 3")

	}

	if ls[1].Task == notes[1] {
		t.Errorf("Task: %s is not deleted to list", notes[1])
	}
}

func TestSave(t *testing.T) {
	ls := todo.List{}
	notes := []string{"the task #1", "the task #2", "the task #3"}

	for _, note := range notes {
		ls.Add(note)
	}

	fl := "data.txt"
	err := ls.Save(fl)
	defer os.Remove(fl)

	if err != nil {
		t.Errorf("Unable to save file with error: %s", err.Error())

	}

}

func TestLoad(t *testing.T) {
	ls := todo.List{}
	notes := []string{"the task #1", "the task #2", "the task #3"}

	for _, note := range notes {
		ls.Add(note)
	}

	fl := "data.txt"
	err := ls.Save(fl)
	if err != nil {
		t.Errorf("Unable to save file with error: %s", err)

	}

	// Reading the data
	data := todo.List{}
	err = data.Load(fl)
	if err != nil || len(data) != 3 {
		t.Errorf("Unable to Load file with error: %s", err)
	}
	defer os.Remove(fl)
}
