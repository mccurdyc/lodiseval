package main

import (
	"fmt"
	"log"

	"github.com/mccurdyc/lodiseval/pkg/replica"
)

func main() {
	var (
		store1 = newMapStateHandler(map[interface{}]interface{}{})
		store2 = newMapStateHandler(map[interface{}]interface{}{})
		store3 = newMapStateHandler(map[interface{}]interface{}{})
	)

	r1 := replica.New("1", store1)
	r2 := replica.New("2", store2)
	r3 := replica.New("3", store3)

	alg := leaderAlg{
		leader:   "1",
		replicas: map[string]*replica.Replica{"1": r1, "2": r2, "3": r3},
	}

	// EXAMPLE: set state in leader; check that followers received state.
	alg.Set("a", "hello")
	alg.Set("b", "world")
	alg.Set("c", "colton")

	// EXAMPLE: get states from all replicas
	a, err := alg.Get("a")
	if err != nil {
		fmt.Printf("unexpected err: %+v", err)
	}

	if a != "hello" {
		fmt.Printf("failed - got: (%s) want: (%s)\n", a, "hello")
	}

	a, errA := alg.Get("a")
	b, errB := alg.Get("b")
	c, errC := alg.Get("c")

	fmt.Printf("a: %s, err: %+v\n", a, errA)
	fmt.Printf("b: %s, err: %+v\n", b, errB)
	fmt.Printf("c: %s, err: %+v\n", c, errC)
}

// EXAMPLES: Below are trivial example implementations and could probably be deleted.
type leaderAlg struct {
	leader   string
	replicas map[string]*replica.Replica
}

// Set sets the specified key to the specified value in the leader. If setting state
// in the leader is successful, state is attempted to be replicated to the follower
// replicas.
func (l *leaderAlg) Set(k interface{}, v interface{}) error {
	if err := l.replicas[l.leader].StateHandler().Set(k, v); err != nil {
		return fmt.Errorf("failed to set in leader (%s): %w", l.leader, err)
	}

	// Serially set state across replicas.
	for id, r := range l.replicas {
		if err := r.StateHandler().Set(k, v); err != nil {
			// We shouldn't really care if a replica fails in this example, so just log.
			log.Println(fmt.Errorf("failed to set in follower (%s): %w", id, err))
		}
	}

	return nil
}

// Get retreives the state for a specified key from the leader replica.
//
// EXAMPLE: This is a trivial example where state is always retreived from the leader.
func (l *leaderAlg) Get(k interface{}) (interface{}, error) {
	v, err := l.Replicas()[l.leader].StateHandler().Get(k)
	if err != nil {
		return nil, fmt.Errorf("failed to get from leader (%s): %w", l.leader, err)
	}

	return v, nil
}

// Replicas retreives the replicas where state is being replicated.
func (l *leaderAlg) Replicas() map[string]*replica.Replica {
	m := make(map[string]*replica.Replica)

	for _, r := range l.replicas {
		m[r.ID()] = r
	}

	return m
}

// TODO: move to a mapstore.go file in package main or something. This is just an example.
type mapStateHandler struct {
	m map[interface{}]interface{}
}

func newMapStateHandler(m map[interface{}]interface{}) *mapStateHandler {
	return &mapStateHandler{
		m: m,
	}
}

func (s *mapStateHandler) Set(k, v interface{}) error {
	s.m[k] = v

	return nil
}

func (s *mapStateHandler) Get(k interface{}) (interface{}, error) {
	return s.m[k], nil
}

func (s *mapStateHandler) GetState() (interface{}, error) {
	return s.m, nil
}
