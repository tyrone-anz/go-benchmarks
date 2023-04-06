# Go benchmarks of common usage

- [fmt.Sprintf](fmt/sprintf_test.go)
- [regexp.Regexp.MatchString](regexp/match_test.go)
- [runtime.Caller](runtime/caller_test.go)
- [append](slice/append_test.go)

## Learnings

- Avoid use of fmt.Sprintf, use static strings
- Add capacity to slices and maps
- Having | on regex pattern seems not better compared to multiple pattern matching
