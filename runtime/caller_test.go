package runtime

import (
	"runtime"
	"testing"
)

func BenchmarkRuntimeCaller(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _, _ = runtime.Caller(2)
	}

	// BenchmarkRuntimeCaller-10    	 2544073	       468.0 ns/op	     232 B/op	       2 allocs/op
}

func BenchmarkRuntimeCallers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pc := make([]uintptr, 15)
		runtime.Callers(1, pc)
	}

	// BenchmarkRuntimeCallers-10    	 4273722	       276.5 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkRuntimeCallersAndFrame(b *testing.B) {
	const depthMax = 15
	for n := 0; n < b.N; n++ {
		var pcs [depthMax]uintptr
		n := runtime.Callers(1, pcs[:])
		pc := pcs[:n]

		runtime.CallersFrames(pc)
	}

	// BenchmarkRuntimeCallersAndFrame-10    	 3292882	       364.3 ns/op	     352 B/op	       2 allocs/op
}

func BenchmarkRuntimeCallersAndFrameAndNextAll(b *testing.B) {
	const depthMax = 15
	for n := 0; n < b.N; n++ {
		var pcs [depthMax]uintptr
		n := runtime.Callers(1, pcs[:])
		pc := pcs[:n]

		frames := runtime.CallersFrames(pc)
		for _, has := frames.Next(); has; {
			_, has = frames.Next()
		}
	}

	// BenchmarkRuntimeCallersAndFrameAndNextAll-10    	 1478937	       786.8 ns/op	     352 B/op	       2 allocs/op
}
