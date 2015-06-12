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

	"golang.org/x/net/context"

	"github.com/satori/go.uuid"
)

// keys uses to store information under context's values map
const (
	ctxKeyEnv = "github.com/gostack/ctxinfo:EnvInfo"
	ctxKeyTx  = "github.com/gostack/ctxinfo:TxInfo"
)

// EnvInfo stores information that are pertinent to the environment, which won't change
// over the execution of a program and can be safely cached.
type EnvInfo struct {
	Hostname    string
	Application string
}

// NewEnvInfo returns a new initialized EnvInfo instance.
func NewEnvInfo(application string) EnvInfo {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return EnvInfo{
		Hostname:    hostname,
		Application: application,
	}
}

// EnvContext creates a new context containing an EnvInfo object in it's values storage.
func EnvContext(ctx context.Context, application string) context.Context {
	return context.WithValue(ctx, ctxKeyEnv, NewEnvInfo(application))
}

// EnvFromContext retrieves an EnvInfo stored within a context.
func EnvFromContext(ctx context.Context) EnvInfo {
	if info, ok := ctx.Value(ctxKeyEnv).(EnvInfo); ok {
		return info
	}

	return EnvInfo{}
}

// Info is a struct that stores information about a specific transaction. It embeds the
// environment information but unlike EnvInfo, it is expected to be created multiple time
// over the execution of a program.
type TxInfo struct {
	TransactionID uuid.UUID
}

// NewTxInfo returns a new initialized TxInfo instance.
func NewTxInfo() TxInfo {
	return TxInfo{
		TransactionID: uuid.NewV4(),
	}
}

// TxContext creates a new context containing an TxInfo object in it's values storage.
func TxContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyTx, NewTxInfo())
}

// TxFromContext retrieves an TxInfo stored within a context.
func TxFromContext(ctx context.Context) TxInfo {
	if info, ok := ctx.Value(ctxKeyTx).(TxInfo); ok {
		return info
	}

	return TxInfo{}
}

// Info is simply a object that combines EnvInfo and TxInfo for easy of consumption.
type Info struct {
	EnvInfo
	TxInfo
}

// FromContext simply gets the EnvInfo and TxInfo object out of the context, returning
// a combined representation of it.
func FromContext(ctx context.Context) Info {
	return Info{EnvFromContext(ctx), TxFromContext(ctx)}
}
