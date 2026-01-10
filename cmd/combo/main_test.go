package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/bobg/seqs"
)

func TestRun2(t *testing.T) {
	cases := []struct {
		cmd  string
		inp  []string
		want []string
	}{{
		cmd:  "perm",
		inp:  []string{"a", "b", "c"},
		want: []string{"a b c", "b a c", "c a b", "a c b", "b c a", "c b a"},
	}, {
		cmd:  "comb 2",
		inp:  []string{"a", "b", "c"},
		want: []string{"a b", "a c", "b c"},
	}, {
		cmd:  "rcomb 2",
		inp:  []string{"a", "b", "c"},
		want: []string{"a a", "a b", "a c", "b b", "b c", "c c"},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			args := strings.Fields(tc.cmd)
			t.Run("args", func(t *testing.T) {
				out := new(bytes.Buffer)
				err := run2(strings.NewReader(""), out, append(args, tc.inp...))
				if err != nil {
					t.Fatal(err)
				}
				lines, errptr := seqs.Lines(out)
				got := slices.Collect(lines)
				if err := *errptr; err != nil {
					t.Fatal(err)
				}
				if !slices.Equal(got, tc.want) {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			})
			t.Run("stdin", func(t *testing.T) {
				out := new(bytes.Buffer)
				inp := strings.NewReader(strings.Join(tc.inp, "\n") + "\n")
				err := run2(inp, out, args)
				if err != nil {
					t.Fatal(err)
				}
				lines, errptr := seqs.Lines(out)
				got := slices.Collect(lines)
				if err := *errptr; err != nil {
					t.Fatal(err)
				}
				if !slices.Equal(got, tc.want) {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			})
		})
	}
}
