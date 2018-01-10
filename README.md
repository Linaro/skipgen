# Overview

Skiplist Generator (skipgen) is a program that will generate a skiplist given a
yaml file and optionally a board name, branch name, and environment name.

[![Build Status](https://travis-ci.org/Linaro/skipgen.svg?branch=master)](https://travis-ci.org/Linaro/skipgen)

## Download and Install

Download release for your OS and architecture at
https://github.com/Linaro/skipgen/releases. Extract and run the 'skipfile'
binary.

## Usage

    skipgen [--board <boardname>] [--branch <branchname>] [--environment <environmentname] [--version] <skipfile.yaml>

## Example Usage

Show all skips available:

    $ skipgen examples/skipfile.yaml
    breakpoint_test_arm64
    ftracetest
    fw_filesystem.sh
    pstore_tests
    run.sh
    run_fuse_test.sh
    run_vmtests
    seccomp_bpf
    ...

Show skips that apply to the x15 board in the production environment and branch 4.4:

    $ skipgen --board=x15 --environment=staging --branch=4.4 examples/skipfile.yaml
    run_vmtests
    seccomp_bpf

## Skipfile Format

See [examples/skipfile.yaml](examples/skipfile.yaml).

## Building

1. Install golang. i.e. on debian-based systems, run `apt-get install golang`.
2. Set GOPATH. See https://github.com/golang/go/wiki/SettingGOPATH.
3. Install go dependencies. `go get -t ./...`
4. install golint. `go get -u github.com/golang/lint/golint`
   Don't forget to setup the path PATH="$GOPATH/bin:$PATH"
5. make skipgen
6. `./skipgen`

## Testing

skipgen includes unit tests that can be run using `go test`. The `make test`
target will also run 'go vet' and 'golint'. golint may need to be installed
(`go get -u github.com/golang/lint/golint`)

## Releasing

If travis has a github api key provided (see goreleaser docs), it will upload
the binaries to the release after making a release in github (or pushing a
tag). Otherwise, after tagging, a release can be made by running the following,
where 'v0.1.2' is the recent tag:

    export GITHUB_TOKEN=xxxxxxxxxxyyyyyyyyzzzzzzzzz
    make clean
    TRAVIS_TAG=v0.1.2 goreleaser
