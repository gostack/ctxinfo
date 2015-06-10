package main

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
