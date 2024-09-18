package comptime_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/frobware/comptime"
)

func init() {
	// ListenAndServe used by the Makefile target
	// benchmark-profile only.
	go func() {
		if port := os.Getenv("BENCHMARK_PROFILE_PORT"); port != "" {
			if err := http.ListenAndServe("localhost:"+port, nil); err != nil {
				panic(err)
			}
		}
	}()
}

// Standard Input Parsing Comparison

// BenchmarkStandardLibraryParseDuration tests the performance of the
// standard library's time.ParseDuration function with a typical
// input.
func BenchmarkStandardLibraryParseDuration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := time.ParseDuration("20h31m23s647ms")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComptimeParseStandardInput tests the performance of
// comptime.ParseDuration with the same input as the standard library
// benchmark.
func BenchmarkComptimeParseStandardInput(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := comptime.ParseDuration("20h31m23s647ms", comptime.Millisecond, comptime.ParseModeMultiUnit, comptime.NoRangeChecking)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Short Duration Parsing Comparison

// BenchmarkStandardLibraryParseDurationShort tests the performance of
// the standard library's time.ParseDuration function with a short
// duration input.
func BenchmarkStandardLibraryParseDurationShort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := time.ParseDuration("1h")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkParseDurationShort tests the performance of comptime
// parsing a short duration string.
func BenchmarkParseDurationShort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := comptime.ParseDuration("1h", comptime.Millisecond, comptime.ParseModeMultiUnit, comptime.NoRangeChecking)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Comptime Specific Features

// BenchmarkParseDurationMultiUnitMode tests comptime's performance
// with a complex duration string including days, which is not
// supported by the standard library.
func BenchmarkParseDurationMultiUnitMode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := comptime.ParseDuration("24d20h31m23s647ms", comptime.Millisecond, comptime.ParseModeMultiUnit, comptime.NoRangeChecking)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkParseDurationSingleUnitMode tests comptime's performance
// in single-unit mode, demonstrating its flexibility in parsing
// modes.
func BenchmarkParseDurationSingleUnitMode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := comptime.ParseDuration("2147483647ms", comptime.Millisecond, comptime.ParseModeSingleUnit, comptime.NoRangeChecking)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkParseDurationWithRangeChecker tests comptime's performance
// with a custom range checker, showcasing this unique feature.
func BenchmarkParseDurationWithRangeChecker(b *testing.B) {
	checker := func(position int, value time.Duration, totalSoFar time.Duration) bool {
		return totalSoFar < 30*24*time.Hour // Allow up to 30 days
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := comptime.ParseDuration("24d20h31m23s647ms", comptime.Millisecond, comptime.ParseModeMultiUnit, checker)
		if err != nil {
			b.Fatal(err)
		}
	}
}
