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

import "golang.org/x/net/context"

// ctxKey is a string identifier to be used when storing Info into a context,
// keeping all package data under the same key avoids collision from other packages.
const ctxKey = "github.com/gostack/ctxinfo"

// NewContext creates a new context containing an Info object in it's values storage.
func NewContext(ctx context.Context, info Info) context.Context {
	return context.WithValue(ctx, ctxKey, info)
}

// InfoFromContext retrieves an Info stored within a context.
func FromContext(ctx context.Context) Info {
	if info, ok := ctx.Value(ctxKey).(Info); ok {
		return info
	}

	return Info{}
}
