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
	"testing"

	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

func TestContextCreations(t *testing.T) {
	ctx := TxContext(
		EnvContext(context.Background(), "myapp"),
	)

	env := EnvFromContext(ctx)
	if env.Application != "myapp" {
		t.Errorf("env info not properly initialized in context: %#v", env)
	}

	tx := TxFromContext(ctx)
	if tx.TransactionID == uuid.Nil {
		t.Errorf("tx info not properly initialized in context: %#v", tx)
	}
}
