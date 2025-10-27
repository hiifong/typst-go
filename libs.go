package typst

// #cgo LDFLAGS: -ltypst
// #cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/libs/x86_64-apple-darwin -Wl,-rpath,${SRCDIR}/libs/x86_64-apple-darwin
// #cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/libs/aarch64-apple-darwin -Wl,-rpath,${SRCDIR}/libs/aarch64-apple-darwin
// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/libs/x86_64-unknown-linux-musl -Wl,-rpath=${SRCDIR}/libs/x86_64-unknown-linux-musl
// #cgo linux,arm64 LDFLAGS: -L${SRCDIR}/libs/aarch64-unknown-linux-musl -Wl,-rpath=${SRCDIR}/libs/aarch64-unknown-linux-musl
import "C"
