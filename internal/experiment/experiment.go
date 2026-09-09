package experiment

import "time"

// A struct is used to create a collection of members of different data types into a single variable.
// A struct can be useful for organizing related data together.

// Struct tags control how a struct is formatted when it is read by the outside world.
// They are metadata attached to fields that other packages can read.

type Experiment struct {
	ExperimentID   int    `json:"id"`
	ExperimentName string `json:"name"`
	Description    string `json:"description"`

	Status ExperimentStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Tags []string `json:"tags"`
}

type ExperimentStatus string

const (
	StatusPlanned   ExperimentStatus = "planned"
	StatusActive    ExperimentStatus = "active"
	StatusCompleted ExperimentStatus = "completed"
)
