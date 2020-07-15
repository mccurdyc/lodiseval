package replica

func defaultReplicaOptions() replicaOptions {
	return replicaOptions{}
}

type replicaOptions struct{}

type ReplicaOptions interface {
	apply(*replicaOptions)
}

type funcReplicaOptions struct {
	f func(*replicaOptions)
}

func (fro *funcReplicaOptions) apply(ro *replicaOptions) {
	fro.f(ro)
}

func newFuncReplicaOption(f func(*replicaOptions)) *funcReplicaOptions {
	return &funcReplicaOptions{
		f: f,
	}
}

// New creates a new replica with the specified identifier, given a store and input stream.
func New(id string, store StateHandler, opts ...ReplicaOptions) *Replica {
	dopts := defaultReplicaOptions()

	for _, opt := range opts {
		opt.apply(&dopts)
	}

	return &Replica{
		id: id,
		sh: store,
	}
}

type Replica struct {
	id string
	sh StateHandler
}

// ID returns the specified identifier for the replica.
func (r *Replica) ID() string {
	return r.id
}

// Store returns the replica's store.
func (r *Replica) StateHandler() StateHandler {
	return r.sh
}

// Algorithm defines the minimal interface that an algorithm must implement for
// storing and retreiving state from a Store.
type Algorithm interface {
	// Set determines how to set a key to a specified value across multiple state handlers.
	Set(interface{}, interface{}) error
	// Get determines how to retreive the value for a specified key from multiple state handlers.
	Get(interface{}) (interface{}, error)
	// Replicas retreives the replicas where state is being replicated.
	Replicas() map[string]StateHandler
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

type StateHandler interface {
	Setter
	Getter
}
