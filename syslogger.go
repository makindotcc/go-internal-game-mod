package main

// #include "syslogger.h"
import "C"

import "unsafe"

type Syslogger struct {
}

func (s Syslogger) Write(p []byte) (n int, err error) {
	// very unsafe, rewrite in rust required!
	// just like the android team did!
	// https://android.googlesource.com/platform/system/bt/+/refs/tags/android-12.0.0_r2/gd/rust/
	// maybe now finally phones won't get hacked while listening to music in a car d-_-b
	nullTerminated := make([]byte, len(p)+1)
	copy(nullTerminated, p)
	nullTerminated[len(p)] = '\000'

	C.syslogWrapper(unsafe.Pointer(&nullTerminated[0]))
	return len(p), nil
}
