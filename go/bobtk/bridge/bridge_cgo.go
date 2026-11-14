//go:build bobtk_cgo

package bridge

/*
#cgo CFLAGS: -O2
static const char *bobtk_bridge_ping(void) { return "pong"; }
*/
import "C"

// Available reports whether the cgo bridge was compiled in.
func Available() bool { return true }

// Ping proves the cgo toolchain path end-to-end (Go -> C -> Go). The next
// step links this against libbobgui from the native _build and forwards a
// real call (e.g. Media AudioProcessor processBlock); until then this is a
// toolchain spike only, and the dashboard keeps badging the bridge Planned.
//
// Toolchain note (2026-11-14, this laptop): Go 1.27.4 cgo cannot parse
// GCC 15.2 (w64devkit) objects — "cannot parse gcc output ... as ELF,
// Mach-O, PE, XCOFF" because GCC emits pe-bigobj-x86-64. Reproduced in
// isolation; -g0 does not help and -mno-big-obj is not a GCC option here.
// Unblock by using an older GCC or a Go release with bigobj support; the
// default (stub) build is unaffected and fully tested.
func Ping() string { return C.GoString(C.bobtk_bridge_ping()) }
