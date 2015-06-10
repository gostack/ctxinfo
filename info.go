package ctxinfo

import (
	"os"

	"github.com/satori/go.uuid"
)

// Info is a struct that stores information about the system and the request currently
// being processed.
type Info struct {
	Service       string
	Hostname      string
	TransactionID uuid.UUID
}

// NewInfo returns a new initialized Info instance.
func NewInfo(service string) Info {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return Info{
		Service:       service,
		Hostname:      hostname,
		TransactionID: uuid.NewV4(),
	}
}
