package fmt

import (
	"fmt"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("%v.%v", "testname", "testoperation")
	}

	// BenchmarkSprintf-10    	21999354	        55.08 ns/op	      24 B/op	       1 allocs/op
}

func BenchmarkSprintfError(b *testing.B) {
	err := fmt.Errorf("unexpected error")

	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("something bad happened: %v", err)
	}

	// BenchmarkSprintfError-10    	13721917	        73.70 ns/op	      48 B/op	       1 allocs/op
}
