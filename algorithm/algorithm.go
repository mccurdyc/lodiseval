package algorithm

import (
	"context"

	"github.com/mccurdyc/lodiseval/replica"
)

// Algorithm defines the minimal interface that an algorithm must implement for
// storing and retreiving state.
type Algorithm interface {
	// Set determines how to set a key to a specified value across multiple state handlers.
	Set(interface{}, interface{}) error
	// Get determines how to retreive the value for a specified key from multiple state handlers.
	Get(interface{}) (interface{}, error)
	// Describe describes the algorithm in detail.
	Describe() string
}

// Config contains configuration parameters used in the algorithm factory function.
type Config struct {
	Leader   string
	Replicas map[string]*replica.Replica

	// Opaque is an opaque configuration.
	Opaque map[string]string
}

// Factory is a factory function for creating an algorithm.
type Factory func(context.Context, *Config) (Algorithm, error)
