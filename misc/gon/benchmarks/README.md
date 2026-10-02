# Go / Gon benchmarks

Run from the repository root after building the public Gon tools:

```sh
GON_BASELINE_GO=/absolute/path/to/unmodified/go python3 misc/gon/benchmark.py
```

This compares three standalone programs:

1. Unmodified Go compiling the legacy implementation.
2. Gon compiling exactly the same legacy source.
3. Gon compiling the equivalent implementation with new syntax.

The runner copies `fixtures/common.go`, `common_test.go` and one of
`legacy.go` / `modern.go` into isolated modules. The fixtures directory itself
is not a buildable package: the implementations are alternatives, and the
runner supplies `nonce.go` with the build measurement constant.

The default run takes ten runtime samples of 300 ms per benchmark and seven
build samples per variant. Use `--help` for overrides. All nine correctness
tests must pass in every variant, and the three executables must print the
same checksum, failure count and effects. The runner also verifies that the
legacy source files are byte-identical across toolchains.

The twelve benchmarks cover successful/failed propagation, wrapping handlers,
conditional branches, captured callbacks, present/absent guarded fields,
guarded calls, coalescing assignment, and combined parse/transform pipelines.
All reported operations process **64 elements**, including bytes and
allocations per operation. Shared counters make lazy evaluation observable;
they also add instrumentation cost equally to the implementations. Timed
workloads contain no I/O. Datasets and string inputs are prepared before timing.

Each run retains `report.html`, `summary.json`, raw samples, command logs,
source copies and binaries in a new directory under `pkg/gon-benchmarks/`.
These outputs are ignored by Git. The JSON includes compiler hashes, toolchain
versions, repository HEAD, integrated upstream provenance, machine details,
and environment settings. No compiler or language implementation is changed.

Runtime benchmarks use precompiled test binaries, default optimizations,
`GOMAXPROCS=1`, `GOGC=100`, `GOMEMLIMIT=off`, and no cgo. Each variant is warmed
before measurement. Variants run serially in randomized order within each
round. Medians and observed ranges are retained; the descriptive bootstrap
intervals use paired round ratios. Small changes require more evidence than
one machine/session, especially without CPU affinity or thermal control.

Build measurements use warmed dependency caches, `-p=1`, `-trimpath`, and an
observable constant change to force recompilation and linking of the main
package. Separate no-change builds measure the cached path. The runtime,
compiler toolchain and dependency cache are not rebuilt from scratch in these
measurements. Wall time includes the public command/launcher. RSS is the
maximum reported by `/usr/bin/time`, not summed concurrent memory across the
process tree. Binary size includes runtime and debug metadata; no stripping.

Use an upstream Go from the same revision as Gon's integrated base if the
goal is to isolate the fork's overhead. When versions differ, Go versus Gon
also includes upstream compiler and standard-library differences. Comparing
the two Gon variants isolates the effect of source syntax on this workload
with the same compiler/runtime. These benchmarks do not establish performance
for other programs or architectures.
