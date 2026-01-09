package main

import (
	"context"
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/bobg/errors"
	"github.com/bobg/seqs"
	"github.com/bobg/subcmd/v2"

	"github.com/bobg/combo"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	return subcmd.Run(context.Background(), maincmd{}, os.Args[1:])
}

type maincmd struct{}

func (m maincmd) Subcmds() subcmd.Map {
	return subcmd.Commands(
		"perm", doPerm, "permutations", nil,
		"comb", doComb, "combinations", subcmd.Params(
			"n", subcmd.Int, 0, "number of elements to choose",
		),
		"rcomb", doRComb, "combinations with replacement", subcmd.Params(
			"n", subcmd.Int, 0, "number of elements to choose",
		),
	)
}

func doPerm(ctx context.Context, args []string) error {
	slice, err := readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	perms := combo.Permutations(slice)
	return writeOutput(perms)
}

func doComb(ctx context.Context, n int, args []string) error {
	if n <= 0 {
		return fmt.Errorf("n must be positive")
	}
	slice, err := readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	combs := combo.Combinations(slice, n)
	return writeOutput(combs)
}

func doRComb(ctx context.Context, n int, args []string) error {
	if n <= 0 {
		return fmt.Errorf("n must be positive")
	}
	slice, err := readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	combs := combo.CombinationsWithReplacement(slice, n)
	return writeOutput(combs)
}

func readInput(args []string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}
	lines, errptr := seqs.Lines(os.Stdin)
	result := slices.Collect(lines)
	return result, errors.Wrap(*errptr, "reading lines from stdin")
}

func writeOutput(seq iter.Seq[[]string]) error {
	for item := range seq {
		if _, err := fmt.Printf("%s\n", strings.Join(item, " ")); err != nil {
			return errors.Wrap(err, "writing output")
		}
	}
	return nil
}
