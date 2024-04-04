package comptime_test

import (
	"fmt"
	"os"

	"github.com/frobware/comptime"
)

// ExampleParseDuration gets copied to the documentation.
func ExampleParseDuration() {
	if duration, err := comptime.ParseDuration("1h", comptime.Millisecond, comptime.ParseModeSingleUnit, comptime.NoRangeChecking); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%vms\n", duration.Milliseconds())
	}

	if duration, err := comptime.ParseDuration("24d20h31m23s647ms", comptime.Millisecond, comptime.ParseModeMultiUnit, comptime.NoRangeChecking); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%vms\n", duration.Milliseconds())
	}

	if duration, err := comptime.ParseDuration("500", comptime.Millisecond, comptime.ParseModeMultiUnit, comptime.NoRangeChecking); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%vms\n", duration.Milliseconds())
	}

	// Output:
	// 3600000ms
	// 2147483647ms
	// 500ms

}
