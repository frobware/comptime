# comptime: Parse composite time durations, including support for days

The `comptime` package (composite time) extends the capabilities of the standard library's `time.ParseDuration` function. It introduces support for 'days', amongst other features.

## Key Features

- Supports 'days' as an additional time unit (denoted by 'd').
- Flexible parsing modes:
  - Multi-unit mode for parsing composite strings (e.g., "24d20h31m23s647ms").
  - Single-unit mode to ensure only one unit type is present.
- Custom range checking via a callback function, allowing for user-defined limits and early termination.
- Ability to specify a default unit for values without an explicit unit (e.g., interpreting "500" as 500ms).

## Performance

`comptime` is designed for efficiency. For detailed performance comparisons with the standard library, see [BENCHMARKS.md](./BENCHMARKS.md).

## Development

```sh
$ make
$ make benchmark
$ make benchmark-profile
```

For more detailed information and usage examples, please refer to the [package documentation](https://pkg.go.dev/github.com/frobware/comptime).
