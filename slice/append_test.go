package slice

import (
	"testing"
)

func BenchmarkAppend10ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0)
		for i := 0; i < 10; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend10ToZeroSliceCapacity-10    	 7263372	       162.7 ns/op	     496 B/op	       5 allocs/op
}

func BenchmarkAppend20ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0)
		for i := 0; i < 20; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend20ToZeroSliceCapacity-10    	 4936579	       237.1 ns/op	    1008 B/op	       6 allocs/op
}

func BenchmarkAppend50ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0)
		for i := 0; i < 50; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend50ToZeroSliceCapacity-10    	 3219450	       370.9 ns/op	    2032 B/op	       7 allocs/op
}

func BenchmarkAppend10To10SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0, 10)
		for i := 0; i < 10; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend10To10SliceCapacity-10    	136048238	         8.809 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkAppend20To20SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0, 20)
		for i := 0; i < 20; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend20To20SliceCapacity-10    	62652823	        17.30 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkAppend50To50SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0, 50)
		for i := 0; i < 50; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend50To50SliceCapacity-10    	27602334	        42.82 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkAppend50To20SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]string, 0, 20)
		for i := 0; i < 50; i++ {
			s = append(s, "more important data")
		}
	}

	// BenchmarkAppend50To20SliceCapacity-10    	 4560397	       270.1 ns/op	    1920 B/op	       2 allocs/op
}
