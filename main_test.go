package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	if got != 4 {
		t.Fatalf("expected 4, got %d", got)
	}
}
