package todo_test

import (
	"testing"

	todo "github.com/atulsingh0/golang-projects/22-todo-app/cmd/todo"
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

	err := ls.Complete(1) // Item #, It is not a index
	if err != nil {
		t.Errorf("err: %v", err.Error())
	}
	if !ls[0].Done {
		t.Errorf("Expected %v, Got %v", true, ls[0].Done)
	}
}
