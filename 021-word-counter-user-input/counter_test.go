package main

import (
	"bytes"
	"testing"
)

func TestCounterWords(t *testing.T) {
	lines := bytes.NewBufferString("this is \n me calling \n via iPhone")
	expected := 6
	result := count(lines, false, false)

	if expected != result {
		t.Errorf("Expected words: %d, got %d instead.\n", expected, result)
	}
}

func TestCounterLines(t *testing.T) {
	lines := bytes.NewBufferString("this is \n me calling \n via iPhone")
	expected := 3
	result := count(lines, true, false)

	if expected != result {
		t.Errorf("Expected words: %d, got %d instead.\n", expected, result)
	}
}

func TestCounterBytes(t *testing.T) {
	lines := bytes.NewBuffer([]byte("this is me calling via iPhone"))
	expected := 29
	result := count(lines, false, true)

	if expected != result {
		t.Errorf("Expected words: %d, got %d instead.\n", expected, result)
	}
}

func TestCounterBytesWithNewLines(t *testing.T) {
	lines := bytes.NewBuffer([]byte("this is me \n calling via \n iPhone"))
	expected := 33
	result := count(lines, false, true)

	if expected != result {
		t.Errorf("Expected words: %d, got %d instead.\n", expected, result)
	}
}
