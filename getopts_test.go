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

func Test_realInit(t *testing.T) {
	a1 := []string{
		`/path/appname`,
	}
	a2 := []string{
		`appname`,
		`-a`,
		`-i`,
		`--infile`,
		`config.in`,
		`--help`,
	}
	tests := []struct {
		name string
		args []string
	}{
		{"1", a1},
		{"2", a2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			realInit(tt.args)
			if nil == gIterator {
				t.Errorf("%q: realInit() = %v, want %v",
					tt.name, nil, "not NIL")
			}
		})
	}
} // Test_realInit()

func TestGet(t *testing.T) {
	// Set getopts `init()` :: []string{
	// 	"testingApplication",
	// 	`-a`, // Flag option
	// 	`-i`, // Error: expected with argument => skipped
	// 	`--infile`, `config.in`,
	// 	`--help`, // Flag option
	// }
	p1 := ""
	o1 := "-help"
	a1 := TArg("")
	w1 := false

	p2 := "a|b:|-celler|d|h|help"
	o2 := "a"
	a2 := TArg("")
	w2 := true

	p3 := p2
	o3 := "?"
	a3 := TArg("")
	w3 := false

	tests := []struct {
		name     string
		pattern  string
		wantOpt  string
		wantArg  TArg
		wantMore bool
	}{
		{"1", p1, o1, a1, w1},
		{"2", p2, o2, a2, w2},
		{"3", p3, o3, a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpt, gotArg, gotMore := Get(tt.pattern)
			if gotOpt != tt.wantOpt {
				t.Errorf("%q: Getopts() gotOpt = %q, want %q",
					tt.name, gotOpt, tt.wantOpt)
			}
			if gotArg != tt.wantArg {
				t.Errorf("%q: Getopts() gotArg = %q, want %q",
					tt.name, gotArg, tt.wantArg)
			}
			if gotMore != tt.wantMore {
				t.Errorf("%q: Getopts() gotMore = %t, want %t",
					tt.name, gotMore, tt.wantMore)
			}
		})
	}
} // TestGet()

/* _EoF_ */
