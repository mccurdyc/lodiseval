package replica

import (
	"log"
)

type server struct {
	UnimplementedReplicaSvcServer

	logger *log.Logger

	id string
}
