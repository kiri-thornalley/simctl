package experiment

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)

	if got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

//func TestCreateExperiment(t *testing.T) {
// create temp directory
// create experiment
// verify experiment.md exists
// verify experiment.json exists
// verify notes directory exists
//}
