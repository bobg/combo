package main

import (
	"context"
	"fmt"
	"io"
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
	return run2(os.Stdin, os.Stdout, os.Args[1:])
}

func run2(r io.Reader, w io.Writer, args []string) error {
	return subcmd.Run(context.Background(), maincmd{r: r, w: w}, args)
}

type maincmd struct {
	r io.Reader
	w io.Writer
}

func (m maincmd) Subcmds() subcmd.Map {
	return subcmd.Commands(
		"perm", m.doPerm, "permutations", nil,
		"comb", m.doComb, "combinations", subcmd.Params(
			"n", subcmd.Int, 0, "number of elements to choose",
		),
		"rcomb", m.doRComb, "combinations with replacement", subcmd.Params(
			"n", subcmd.Int, 0, "number of elements to choose",
		),
	)
}

func (m maincmd) doPerm(ctx context.Context, args []string) error {
	slice, err := m.readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	perms := combo.Permutations(slice)
	return m.writeOutput(perms)
}

func (m maincmd) doComb(ctx context.Context, n int, args []string) error {
	if n <= 0 {
		return fmt.Errorf("n must be positive")
	}
	slice, err := m.readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	combs := combo.Combinations(slice, n)
	return m.writeOutput(combs)
}

func (m maincmd) doRComb(ctx context.Context, n int, args []string) error {
	if n <= 0 {
		return fmt.Errorf("n must be positive")
	}
	slice, err := m.readInput(args)
	if err != nil {
		return errors.Wrap(err, "reading input")
	}
	combs := combo.CombinationsWithReplacement(slice, n)
	return m.writeOutput(combs)
}

func (m maincmd) readInput(args []string) ([]string, error) {
	if len(args) > 0 {
		return args, nil
	}
	lines, errptr := seqs.Lines(m.r)
	result := slices.Collect(lines)
	return result, errors.Wrap(*errptr, "reading lines from stdin")
}

func (m maincmd) writeOutput(seq iter.Seq[[]string]) error {
	for item := range seq {
		if _, err := fmt.Fprintf(m.w, "%s\n", strings.Join(item, " ")); err != nil {
			return errors.Wrap(err, "writing output")
		}
	}
	return nil
}
