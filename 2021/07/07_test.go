package main

import "testing"

// newTestCrabSwarm creates a new crabSwarm for testing.
func newTestCrabSwarm() crabSwarm {
	positions := []position{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	cs := make(crabSwarm)
	for _, p := range positions {
		cs[p]++
	}
	return cs
}

func TestCrabSwarm_cost(t *testing.T) {
	cs := newTestCrabSwarm()
	want := 37
	got := cs.cost(2)
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCrabSwarm_mode(t *testing.T) {
	cs := newTestCrabSwarm()
	want := position(2)
	got := cs.mode()
	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}
