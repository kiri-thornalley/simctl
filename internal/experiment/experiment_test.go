package experiment

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)

	if got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}