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
	for n := 0; n < b.N; n++ {
		pc := make([]uintptr, 15)
		runtime.Callers(1, pc)
		runtime.CallersFrames(pc)
	}

	// BenchmarkRuntimeCallersAndFrame-10    	 3327930	       366.0 ns/op	     352 B/op	       2 allocs/op
}
