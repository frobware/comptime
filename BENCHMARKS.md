# comptime Benchmark Analysis

## Standard Input Parsing Comparison

| Benchmark | Time (ns/op) | Allocations (B/op) | Allocs (allocs/op) |
|-----------|--------------|--------------------|--------------------|
| Standard Library | 79.43 | 0 | 0 |
| comptime | 31.73 | 0 | 0 |

**Analysis:** For parsing standard duration inputs (e.g., "20h31m23s647ms"), comptime is approximately 2.5 times faster than the standard library, while maintaining zero allocations.

## Short Duration Parsing Comparison

| Benchmark | Time (ns/op) | Allocations (B/op) | Allocs (allocs/op) |
|-----------|--------------|--------------------|--------------------|
| Standard Library | 21.27 | 0 | 0 |
| comptime | 6.256 | 0 | 0 |

**Analysis:** For short duration inputs (e.g., "1h"), comptime demonstrates even more significant performance advantages, being about 3.4 times faster than the standard library. This showcases comptime's efficiency in handling simple, common cases.

## comptime Specific Features

| Benchmark | Time (ns/op) | Allocations (B/op) | Allocs (allocs/op) |
|-----------|--------------|--------------------|--------------------|
| Multi-Unit Mode | 38.68 | 0 | 0 |
| Single-Unit Mode | 13.83 | 0 | 0 |
| With Range Checker | 38.71 | 0 | 0 |

**Analysis:**
1. **Multi-Unit Mode:** Even when parsing complex inputs including days (e.g., "24d20h31m23s647ms"), comptime performs efficiently at 38.68 ns/op. This is still faster than the standard library parsing a simpler input.

2. **Single-Unit Mode:** For single-unit parsing (e.g., "2147483647ms"), comptime is exceptionally fast at 13.83 ns/op, demonstrating its efficiency in handling uniform time units.

3. **Range Checking:** The addition of a custom range checker has a negligible impact on performance (38.71 ns/op vs 38.68 ns/op without), showcasing comptime's ability to provide additional functionality with minimal performance cost.

## Key Observations

1. **Consistent Performance Advantage:** comptime outperforms the standard library in all comparable scenarios, with improvements ranging from 2.5x to 3.4x faster.

2. **Scalability:** comptime maintains its performance advantage even when parsing complex, multi-unit durations that include units not supported by the standard library (e.g., days).

3. **Flexibility without Compromise:** The different parsing modes (multi-unit and single-unit) and additional features like range checking do not significantly impact performance.

4. **Zero Allocations:** All operations, both in comptime and the standard library, show zero memory allocations, which is excellent for performance-critical applications.

## Conclusions

1. For standard duration parsing, comptime offers a significant speed improvement over the standard library.
2. comptime excels particularly in parsing short durations, which could be beneficial for applications that frequently parse simple time expressions.
3. The ability to handle more complex inputs (including days) without a substantial performance penalty demonstrates comptime's versatility.
4. The negligible performance impact of range checking allows users to benefit from this additional feature without concerns about speed degradation.
5. The consistent zero-allocation performance across all operations makes comptime suitable for high-performance, memory-sensitive applications.

These benchmark results clearly demonstrate that comptime offers substantial performance improvements over the standard library's duration parsing capabilities, while providing additional functionality such as support for days and custom range checking. Its consistent speed advantage across different input types and complexities, combined with its memory efficiency, makes it an excellent choice for applications where parsing time durations is a frequent operation or where performance is critical.
