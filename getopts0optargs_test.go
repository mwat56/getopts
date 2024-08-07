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

func Test_newOptArgList(t *testing.T) {
	/* This is the list set up by `init()` for testing purposes
	args = []string{
		"testingApplication",
		`-a`, // Flag option
		`-i`, // Error: intended with argument => ignored
		`--infile`, `config.in`,
		`--help`, // Flag option
	}
	*/
	var empty TArg
	a1 := []string{
		"testingApplication",
		`-a`, // Flag option
		`-i`, // Error: intended with argument => ignored
		`--infile`, `config.in`,
		`--help`, // Flag option
	}
	w1 := make(tOptArgList, 0, 4)
	w1 = append(w1, tOptArg{tOpt(`a`), empty})
	w1 = append(w1, tOptArg{tOpt(`i`), empty})
	w1 = append(w1, tOptArg{tOpt(`-infile`), TArg(`config.in`)})
	w1 = append(w1, tOptArg{tOpt(`-help`), empty})

	tests := []struct {
		name string
		args []string
		want *tOptArgList
	}{
		{"1", a1, &w1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOptArgList(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q: newOptArgList() =\n%v\n want \n%v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_newOptArgList()

func Test_tOptArgList_Equal(t *testing.T) {
	oal1 := tOptArgList{}
	woa1 := tOptArgList{}
	oal2 := tOptArgList{tOptArg{"o2", "a2"}}
	woa2 := tOptArgList{tOptArg{"wo2", "wa2"}}
	oal3 := oal1
	woa3 := tOptArgList{tOptArg{"wo3", "wa3"}}

	tests := []struct {
		name  string
		oal   tOptArgList
		other tOptArgList
		want  bool
	}{
		{"1", oal1, woa1, true},
		{"2", oal2, woa2, false},
		{"3", oal3, woa3, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oal.Equal(tt.other); got != tt.want {
				t.Errorf("%q: tOptArgList.Equal() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tOptArgList_Equal()

func Test_tOptArgList_String(t *testing.T) {

	tests := []struct {
		name     string
		oal      tOptArgList
		wantRStr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRStr := tt.oal.String(); gotRStr != tt.wantRStr {
				t.Errorf("%q: tOptArgList.String() = %q, want %q",
					tt.name, gotRStr, tt.wantRStr)
			}
		})
	}
} // Test_tOptArgList_String()

func TestTArg_Bool(t *testing.T) {

	tests := []struct {
		name string
		a    TArg
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Bool(); got != tt.want {
				t.Errorf("%q: TArg.Bool() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestTArg_Bool()

func TestTArg_Float(t *testing.T) {

	tests := []struct {
		name string
		a    TArg
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Float(); got != tt.want {
				t.Errorf("%q: TArg.Float() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestTArg_Float()

func TestTArg_Int(t *testing.T) {

	tests := []struct {
		name string
		a    TArg
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Int(); got != tt.want {
				t.Errorf("%q: TArg.Int() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestTArg_Int()

func TestTArg_String(t *testing.T) {

	tests := []struct {
		name string
		a    TArg
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("%q: TArg.String() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestTArg_String()

/* _EoF_ */
