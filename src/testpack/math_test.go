package main

import "testing"

func TestAddPstv(t *testing.T) {
	sum, err := Add(5, 10)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 15 {
		t.Errorf("expected: 15. have: %d", sum)
	}
}

func TestAddNgtv(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}
