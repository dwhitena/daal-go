package cholesky

// #cgo CXXFLAGS: -I$DAALINCLUDE
// #cgo LDFLAGS: -L$DAALLIB -ldaal_core -ldaal_sequential -lpthread -lm
import "C"
