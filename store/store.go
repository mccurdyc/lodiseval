package store

import "context"

// Store handles setting and retreiving state from a state store.
type Store interface {
	Setter
	Getter
}

type Setter interface {
	// Set sets the value for a key.
	Set(interface{}, interface{}) error
}

type Getter interface {
	// Get retreives the value for a key.
	Get(interface{}) (interface{}, error)
	// GetState returns the entire state of the store.
	GetState() (interface{}, error)
}

// Config contains configuration parameters used in the algorithm factory function.
type Config struct {
	// Opaque is an opaque configuration.
	Opaque map[string]string
}

// Factory is a factory function for creating a state store.
type Factory func(context.Context, *Config) (Store, error)
