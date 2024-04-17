# Apache Beam Golang SDK Sample Project

This repository demonstrates the usage of Apache Beam Golang SDK as part of a small use case whereby words from a file are counted by omitting duplicates.

## Before you begin

Make sure you have a [Go](https://go.dev/) development environment ready.
If you don't, you can follow the instructions in the
[Download and install](https://go.dev/doc/install) page.

## Running the pipeline

```sh
go run main.go

# To run the tests.
go test ./transformations -v
```

## Using other runners

To keep this template small, it only includes the [Direct Runner](https://beam.apache.org/documentation/runners/direct/).

For a comparison of what each runner currently supports, look at the [Beam Capability Matrix](https://beam.apache.org/documentation/runners/capability-matrix/).

To add a new runner, visit the runner's page for instructions on how to include it.