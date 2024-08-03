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

func TestGetopts(t *testing.T) {

	tests := []struct {
		name     string
		pattern  string
		wantROpt string
		wantRArg TArg
		wantROK  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotROpt, gotRArg, gotROK := Getopts(tt.pattern)
			if gotROpt != tt.wantROpt {
				t.Errorf("Getopts() gotROpt = %v, want %v", gotROpt, tt.wantROpt)
			}
			if gotRArg != tt.wantRArg {
				t.Errorf("Getopts() gotRArg = %v, want %v", gotRArg, tt.wantRArg)
			}
			if gotROK != tt.wantROK {
				t.Errorf("Getopts() gotROK = %v, want %v", gotROK, tt.wantROK)
			}
		})
	}
} // TestGetopts()

/* _EoF_ */
