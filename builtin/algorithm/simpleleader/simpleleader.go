package simpleleader

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mccurdyc/lodiseval/algorithm"
)

func Factory(ctx context.Context, conf *algorithm.Config) (algorithm.Algorithm, error) {
	return &Algorithm{
		leader: conf.Leader,
	}, nil
}

type Algorithm struct {
	leader   string
	replicas map[string]algorithm.Replica
}

// Set sets the specified key to the specified value in the leader. If setting state
// in the leader is successful, state is attempted to be replicated to the follower
// replicas.
func (a *Algorithm) Set(k interface{}, v interface{}) error {
	if err := a.replicas[a.leader].Store().Set(k, v); err != nil {
		return fmt.Errorf("failed to set in leader (%s): %w", a.leader, err)
	}

	// Serially set state across replicas.
	for id, r := range a.replicas {
		if err := r.Store().Set(k, v); err != nil {
			// We shouldn't really care if a replica fails in this example, so just log.
			log.Println(fmt.Errorf("failed to set in follower (%s): %w", id, err))
		}
	}

	return nil
}

// Get retreives the state for a specified key from the leader replica.
//
// EXAMPLE: This is a trivial example where state is always retreived from the leader.
func (a *Algorithm) Get(k interface{}) (interface{}, error) {
	v, err := a.replicas[a.leader].Store().Get(k)
	if err != nil {
		return nil, fmt.Errorf("failed to get from leader (%s): %w", a.leader, err)
	}

	return v, nil
}

// Replicas retreives the replicas where state is being replicated.
func (a *Algorithm) Replicas() map[string]algorithm.Replica {
	m := make(map[string]algorithm.Replica)

	for _, r := range a.replicas {
		m[r.ID()] = r
	}

	return m
}

func (a *Algorithm) Describe() string {
	return strings.TrimSpace(`
simpleleader defines a leader replica where write and read requests are sent. The
leader replica handles these requests and after successfully replicating state in
a majority of other replicas, commits them to a log.
`)
}
