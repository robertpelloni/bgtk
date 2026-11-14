//go:build !bobtk_cgo

// Package bridge is the Go↔C++ bridge spike. Without the bobtk_cgo build
// tag the package is a pure-Go stub, so the default build stays offline and
// dependency-free. See bridge_cgo.go for the real cgo path.
package bridge

// Available reports whether the cgo bridge was compiled in.
func Available() bool { return false }

// Ping returns the C bridge greeting, or "" when unavailable.
func Ping() string { return "" }
