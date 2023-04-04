package regexp

import (
	"regexp"
	"testing"
)

func BenchmarkRegexpMatchString(b *testing.B) {
	regex := regexp.MustCompile(`.*/package/file.*\.go$`)

	for n := 0; n < b.N; n++ {
		regex.MatchString("path/to/a/different/file.go")
	}

	// BenchmarkRegexpMatchString-10    	 2075474	       573.1 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkRegexpMatchStringWithTwoPatterns(b *testing.B) {
	regex := regexp.MustCompile(`.*/package/file.*\.go$|.*/directory/file.*\.go$`)

	for n := 0; n < b.N; n++ {
		regex.MatchString("path/to/a/different/file.go")
	}

	// BenchmarkRegexpMatchStringWithTwoPatterns-10    	  850474	      1205 ns/op	       0 B/op	       0 allocs/op
}

func BenchmarkRegexpMatchStringWithFivePatterns(b *testing.B) {
	regex := regexp.MustCompile(`.*/package/file.*\.go$|.*/directory/file.*\.go$|.*/folder/file.*\.go$|.*/bucket/file.*\.go$|.*/storage/file.*\.go$`)

	for n := 0; n < b.N; n++ {
		regex.MatchString("path/to/a/different/file.go")
	}

	// BenchmarkRegexpMatchStringWithFivePatterns-10    	  354511	      3088 ns/op	       0 B/op	       0 allocs/op
}
