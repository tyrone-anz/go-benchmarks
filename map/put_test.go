package _map

import "testing"

func BenchmarkMapPut10ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string)
		for i := 0; i < 10; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut10ToZeroSliceCapacity-10    	 3316593	       341.8 ns/op	     419 B/op	       1 allocs/op
}

func BenchmarkMapPut20ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string)
		for i := 0; i < 20; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut20ToZeroSliceCapacity-10    	 1658184	       728.1 ns/op	    1393 B/op	       2 allocs/op
}

func BenchmarkMapPut50ToZeroSliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string)
		for i := 0; i < 50; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut50ToZeroSliceCapacity-10    	  595094	      1966 ns/op	    3532 B/op	       5 allocs/op
}

func BenchmarkMapPut10To10SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string, 10)
		for i := 0; i < 10; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut10To10SliceCapacity-10    	 5846304	       199.7 ns/op	     419 B/op	       1 allocs/op
}

func BenchmarkMapPut20To20SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string, 20)
		for i := 0; i < 20; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut20To20SliceCapacity-10    	 2766314	       426.2 ns/op	     924 B/op	       1 allocs/op
}

func BenchmarkMapPut50To50SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string, 50)
		for i := 0; i < 50; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut50To50SliceCapacity-10    	 1000000	      1146 ns/op	    2028 B/op	       2 allocs/op
}

func BenchmarkMapPut50To20SliceCapacity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make(map[int]string, 20)
		for i := 0; i < 50; i++ {
			s[i] = "more important data"
		}
	}

	// BenchmarkMapPut50To20SliceCapacity-10    	  682447	      1637 ns/op	    3066 B/op	       3 allocs/op
}
