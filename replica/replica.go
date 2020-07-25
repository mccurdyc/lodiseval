package replica

import "github.com/mccurdyc/lodiseval/store"

// TODO
type Replica struct {
	Store store.Store
}

func (r *Replica) ID() string {
	return "foo"
}
