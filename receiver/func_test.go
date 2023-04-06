package receiver

import (
	"testing"
)

type teststruct struct {
	val string
}

func (s teststruct) ValThroughCopy() string {
	return s.val
}

func (s *teststruct) ValThroughPointer() string {
	return s.val
}

func BenchmarkReceiverFuncThroughCopy(b *testing.B) {
	in := teststruct{val: "very important data"}

	for n := 0; n < b.N; n++ {
		in.ValThroughCopy()
	}

	// BenchmarkReceiverFuncThroughCopy-10    	1000000000	         0.3292 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkReceiverFuncThroughPointer(b *testing.B) {
	in := teststruct{val: "very important data"}

	for n := 0; n < b.N; n++ {
		in.ValThroughPointer()
	}

	// BenchmarkReceiverFuncThroughPointer-10    	1000000000	         0.3219 ns/op	       0 B/op	       0 allocs/op
}
