package replica

import "github.com/mccurdyc/lodiseval/store"

func defaultReplicaOptions() replicaOptions {
	return replicaOptions{}
}

type replicaOptions struct{}

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

type ReplicaOptions interface {
	apply(*replicaOptions)
}

// New creates a new replica with the specified identifier, given a store and input stream.
func New(id string, sh store.Store, opts ...ReplicaOptions) *Replica {
	dopts := defaultReplicaOptions()

	for _, opt := range opts {
		opt.apply(&dopts)
	}

	return &Replica{
		id: id,
		sh: sh,
	}
}

type Replica struct {
	id string
	sh store.Store
}

// ID returns the specified identifier for the replica.
func (r *Replica) ID() string {
	return r.id
}

// StgaStore returns the replica's store.
func (r *Replica) Store() store.Store {
	return r.sh
}
