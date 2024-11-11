package main

import "testing"

func TestAdd(t *testing.T) {
	s := Add(2, 3)
	if s != 5 {
		t.Errorf("Expected sum of 5, but got %v", s)
	}
}
