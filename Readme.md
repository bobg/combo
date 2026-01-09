# Combo - combinatorial operations for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/combo.svg)](https://pkg.go.dev/github.com/bobg/combo)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/combo)](https://goreportcard.com/report/github.com/bobg/combo)
[![Tests](https://github.com/bobg/combo/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/combo/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/combo/badge.svg?branch=main)](https://coveralls.io/github/bobg/combo?branch=main)

This is combo,
a library of combinatorial operations for Go.
Given a slice,
it can compute its permutations,
its n-element combinations,
and its n-element combinations-with-replacement.

There is also a command-line tool for invoking these operations on a list of strings.

## Installation and usage

For library usage,
please see the documentation at [pkg.go.dev](https://pkg.go.dev/github.com/bobg/combo).

Installing the command-line tool:

```sh
go install github.com/bobg/combo/cmd/combo@latest
```

Usage:

```sh
combo perm [ARG ARG ...]
combo comb N [ARG ARG ...]
combo rcomb N [ARG ARG ...]
```

The first argument is the operation to run:
`perm` for permutations;
`comb` for combinations;
and `rcomb` for combinations-with-replacement.

The `comb` and `rcomb` subcommands require a numeric argument, `N`,
the number of strings to choose in each combination.

The input is specified as additional arguments on the command line.
If no such arguments are supplied,
then the input is read from standard input,
one line per string.

Examples:

```sh
$ combo perm 1 2 3
1 2 3
2 1 3
3 1 2
1 3 2
2 3 1
3 2 1
```

```sh
$ combo comb 2 a b c
a b
a c
b c
```

```sh
$ combo rcomb 2 a b c
a a
a b
a c
b b
b c
c c
```
