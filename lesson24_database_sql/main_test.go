package main

import "testing"

func TestCreateStudentRejectsEmptyName(t *testing.T) {
	_, err := createStudent(t.Context(), nil, "   ", 90)
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}
}
