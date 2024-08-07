/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"testing"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

func Test_newExpectedArgs(t *testing.T) {
	p1 := `a| b| c | :d | e: | f |`
	w1 := &tExpectedOpts{previous: p1}

	tests := []struct {
		name    string
		pattern string
		want    *tExpectedOpts
	}{
		{"1", p1, w1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newExpectedOpts(tt.pattern)
			if //!reflect.DeepEqual(got, tt.want) {
			got.previous != tt.want.previous {
				t.Errorf("%q: newExpectedArgs() =\n%v\n want \n%v",
					tt.name, got.previous, tt.want.previous)
			}
		})
	}
} // Test_newExpectedArgs()

func Test_tExpectedOpts_isValid(t *testing.T) {
	/* used by `getopts.init()` in testing/debugging mode:
	args = []string{
		"testingApplication",
		`-a`, // Flag option
		`-i`, // Error: intended with argument => ignored
		`--infile`, `config.in`,
		`--help`, // Flag option
	}
	*/
	eo := newExpectedOpts(`|a|i:|-infile:|-help|`)

	o0, a0, w0 := tOpt("0"), TArg("0"), false
	o1, a1, w1 := tOpt("a"), TArg(""), true
	o2, a2, w2 := tOpt("i"), TArg(""), false
	o3, a3, w3 := tOpt("-infile"), TArg("dummy.txt"), true

	tests := []struct {
		name      string
		opt       tOpt
		arg       TArg
		wantValid bool
	}{
		{"0", o0, a0, w0},
		{"1", o1, a1, w1},
		{"2", o2, a2, w2},
		{"3", o3, a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid := eo.isValid(tt.opt, tt.arg)
			if gotValid != tt.wantValid {
				t.Errorf("%q: tExpectedOpts.isValid() gotValid = %v, want %v",
					tt.name, gotValid, tt.wantValid)
			}
		})
	}
} // Test_tExpectedOpts_isValid()

func Test_tExpectedArgs_parse(t *testing.T) {
	/* used by `getopts.init()` in testing/debugging mode:
	args = []string{
		"testingApplication",
		`-a`, // Flag option
		`-i`, // Error: intended with argument => ignored
		`--infile`, `config.in`,
		`--help`, // Flag option
	}
	*/

	p1 := `a|i:|-infile:|-help`
	e1 := newExpectedOpts(p1)
	p2 := `|a|i:|-k|-help|`
	e2 := newExpectedOpts(p2)
	p3 := `a|i::|-j:|-zzz`
	e3 := newExpectedOpts(p3)
	p4 := `a|i::|::j|-zzz`
	e4 := newExpectedOpts(p4)
	p5 := `a|i::|::|::m|z`
	e5 := newExpectedOpts(p5)
	p6 := `a|i::|j|::k::|z`
	e6 := newExpectedOpts(p6)
	p7 := `a | i:: | j | ::k:: | z`
	e7 := newExpectedOpts(p7)
	p8 := `a | i : | j | : k : | z`
	e8 := newExpectedOpts(p8)

	tests := []struct {
		name    string
		eo      *tExpectedOpts
		pattern string
	}{
		{"1", e1, p1},
		{"2", e2, p2},
		{"3", e3, p3},
		{"4", e4, p4},
		{"5", e5, p5},
		{"6", e6, p6},
		{"7", e7, p7},
		{"8", e8, p8},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.eo.parse(tt.pattern); tt.eo.previous != tt.pattern {
				t.Errorf("%q: tExpectedArgs.parse() = %q, want %q",
					tt.name, tt.eo.previous, tt.pattern)
			}
		})
	}
} // Test_tExpectedArgs_parse()

/* _EoF_ */
