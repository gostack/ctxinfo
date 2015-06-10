/*
Copyright 2015 Rodrigo Rafael Monti Kochenburger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
