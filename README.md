# comptime: Parse composite time durations, including support for days

The `comptime` package (composite time) extends the capabilities of the standard library's `time.ParseDuration` function. It introduces support for 'days' and enables parsing of composite durations from a single string, such as '1d5m200ms'.

## Key Features

- Supports time units: "d" (days), "h" (hours), "m" (minutes), "s" (seconds), "ms" (milliseconds), and "us" (microseconds).
- Parses composite durations like "24d20h31m23s647ms".
- Ensures parsed durations are non-negative.
- Custom Range Checking: Define range constraints on parsed durations through a BoundsChecker callback, enabling early termination based on user-defined limits.

## Performance

`comptime` is designed for efficiency. For detailed performance comparisons with the standard library, see [BENCHMARKS.md](./BENCHMARKS.md).

## Development

```sh
$ make
$ make benchmark
$ make benchmark-profile
```

For more detailed information and usage examples, please refer to the [package documentation](https://pkg.go.dev/github.com/frobware/comptime).
