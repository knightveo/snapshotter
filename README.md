# snapshotter

Incremental snapshot writer and restore tooling.

## Design

The implementation favors explicit state and small interfaces. Error paths, serialization boundaries, and concurrency assumptions are kept visible instead of being hidden behind framework code.

## Development

Build and test commands are documented next to the project files. Benchmarks, when present, are separate from correctness tests.
