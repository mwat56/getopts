/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"reflect"
	"testing"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

func prep4Test() *tOptArgList {
	// if gSomeTestsAreRunning {
	// 	cmdLineArgs := []string{
	// 		`appname`,
	// 		`-a`,
	// 		`-i`,
	// 		`--infile`,
	// 		`config.in`,
	// 		`--help`,
	// 	}
	// 	realInit(cmdLineArgs)
	// 	return newOptArgMap(cmdLineArgs)
	// }
	return gIterator.optArgs
} // prep4Test()

func Test_newIterator(t *testing.T) {
	l1 := prep4Test()
	w1 := gIterator

	tests := []struct {
		name string
		list *tOptArgList
		want *tIterator
	}{
		{"1", l1, w1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newIterator(tt.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q: newOptIterator() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestNewOptIterator()

func Test_tIterator_Next(t *testing.T) {
	// Set getopts `init()` :: []string{
	// 	"testingApplication",
	// 	`-a`, // Flag option
	// 	`-i`, // Error: expected with argument => skipped
	// 	`--infile`, `config.in`,
	// 	`--help`, // Flag option
	// }
	a1 := prep4Test()
	p1 := "|a|i:|-infile:|-help"
	i1 := newIterator(a1).setPattern(p1)
	o1, r1, m1 := tOpt("a"), TArg(""), true

	// "-i" is skipped because of its missing argument
	i2 := i1
	o2, r2, m2 := tOpt("-infile"), TArg("config.in"), true

	i3 := i2
	o3, r3, m3 := tOpt("-help"), TArg(""), false

	tests := []struct {
		name     string
		oi       *tIterator
		wantOpt  tOpt
		wantArg  TArg
		wantMore bool
	}{
		{"1", i1, o1, r1, m1},
		{"2", i2, o2, r2, m2},
		{"3", i3, o3, r3, m3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpt, gotArg, gotMore := tt.oi.Next()
			if gotOpt != tt.wantOpt {
				t.Errorf("%q: tIterator.Next() gotOpt = %q, want %q;\npattern %q",
					tt.name, gotOpt, tt.wantOpt, tt.oi.expected.previous)
			}
			if gotArg != tt.wantArg {
				t.Errorf("%q: tIterator.Next() gotArg = %q, want %q;\npattern %q",
					tt.name, gotArg, tt.wantArg, tt.oi.expected.previous)
			}
			if gotMore != tt.wantMore {
				t.Errorf("%q: tIterator.Next() gotMore = %t, want %t;\npattern %q",
					tt.name, gotMore, tt.wantMore, tt.oi.expected.previous)
			}

		})
	}
} // Test_tIterator_Next()

func Test_tIterator_Reset(t *testing.T) {
	prep4Test()

	oal0 := &tOptArgList{}
	i0 := newIterator(oal0)

	tests := []struct {
		name string
		oi   *tIterator
	}{
		{"0", i0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.oi.Reset()
			if 0 != tt.oi.index {
				t.Errorf("%q: tIteratorReset.() gotROK = %v, want %v",
					tt.name, tt.oi.index, 0)

			}
		})
	}
} // Test_tIterator_Reset()

func Test_tIterator_setPattern(t *testing.T) {
	prep4Test()

	oal0 := &tOptArgList{}
	p0 := "h|-help"
	e0 := &tExpectedOpts{previous: p0}
	i0 := newIterator(oal0)
	w0 := &tIterator{
		optArgs:  oal0,
		expected: e0,
	}

	oal1 := &tOptArgList{}
	p1 := "a|b:|c:|d"
	e1 := &tExpectedOpts{previous: p1}
	i1 := newIterator(oal1)
	w1 := &tIterator{
		optArgs:  oal1,
		expected: e1,
	}

	c2 := []string{
		`appname`,
		`-a`,
		`-b`, // without the required argument
		`--infile`, `config.in`,
		`--help`,
	}
	oal2 := newOptArgList(c2)
	p2 := "||a|b:|c:||d||"
	i2 := newIterator(oal2)
	w2 := newIterator(oal2)

	oal3 := oal2
	p3 := "a|b:|:-infile|d"
	i3 := newIterator(oal3)
	w3 := newIterator(oal3)

	tests := []struct {
		name    string
		iter    *tIterator
		pattern string
		want    *tIterator
	}{
		{"0", i0, p0, w0},
		{"1", i1, p1, w1},
		{"2", i2, p2, w2},
		{"3", i3, p3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.iter.setPattern(tt.pattern)
			if (nil != got) && (nil != tt.want) {
				if nil != got.expected && nil != tt.want.expected {
					if got.expected.previous != tt.want.expected.previous {
						t.Errorf("%q 3: tIterator.setPattern() = %q, want %q",
							tt.name, got.expected.previous, tt.want.expected.previous)
					}
				} else {
					t.Errorf("%q 2: tIterator.setPattern() = %v, want %v",
						tt.name, got.expected, tt.want.expected)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%q 1: tIterator.setPattern() = %v, want %v",
						tt.name, got, tt.want)
				}
			}
		})
	}
} // Test_tIterator_setPattern()

/* _EoF_ */
