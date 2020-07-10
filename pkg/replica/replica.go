package replica

// New creates a new replica with the specified identifier, given a store and input stream.
func New(id string, store Store, in <-chan interface{}) *Replica {
	return &Replica{}
}

// TODO: delete?
// type (
// 	// SetFn specifies how to set the value for a key.
// 	SetFn func(k, v interface{})
// 	// GetFn specifies how to retreive the value for a key.
// 	GetFn func(k interface{}) interface{}
// 	// GetStateFn specifies how to retreive the entire state of the store.
// 	GetStateFn func() interface{}
// )

type Store interface {
	// Set sets the value for a key.
	Set(k, v interface{})
	// Get retreives the value for a key.
	Get(k interface{}) interface{}
	// GetState returns the entire state of the store.
	GetState() interface{}
}

// Replica is a single instance of a, potentially replicated, store.
//
// TODO: Currently, there is no concept of a leader / follower replica.
// Some options for this could be either a leader / follower field on the Replica
// struct or something like a {Leader,Follower}Replica implementation of a Replica interface.
type Replica struct {
	id    string
	in    <-chan interface{}
	store Store
}

// ID returns the specified identifier for the replica.
func (r *Replica) ID() string {
	return r.id
}

// Store returns the replica's store.
func (r *Replica) Store() interface{} {
	return r.store
}
