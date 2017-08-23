package main

//#cgo LDFLAGS: -F /System/Library/PrivateFrameworks -framework login
//extern void SACLockScreenImmediate();
import "C"

func main() {
	C.SACLockScreenImmediate()
}
