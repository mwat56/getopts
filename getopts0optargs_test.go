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

func Test_TArg_Bool(t *testing.T) {
	a0, w0 := TArg(""), false
	a1, w1 := TArg("+"), false
	a2, w2 := TArg("1"), true
	a3, w3 := TArg("n"), false

	tests := []struct {
		name string
		arg  TArg
		want bool
	}{
		{"0", a0, w0},
		{"1", a1, w1},
		{"2", a2, w2},
		{"3", a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Bool(); got != tt.want {
				t.Errorf("%q: TArg.Bool() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_TArg_Bool()

func Test_TArg_Equal(t *testing.T) {
	a0, o0 := TArg(""), TArg("")
	a1, o1 := a0, TArg("1")
	a2, o2 := TArg("2"), TArg("2")

	tests := []struct {
		name string
		a    TArg
		arg  TArg
		want bool
	}{
		{"0", a0, o0, true},
		{"1", a1, o1, false},
		{"2", a2, o2, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.arg); got != tt.want {
				t.Errorf("%q: TArg.Equal() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_TArg_Equal()

func Test_TArg_Float(t *testing.T) {
	a0, w0 := TArg(""), 0.0
	a1, w1 := TArg("1.23"), 1.23
	a2, w2 := TArg("-2.34"), -2.34
	a3, w3 := TArg("n.a."), 0.0

	tests := []struct {
		name string
		arg  TArg
		want float64
	}{
		{"0", a0, w0},
		{"1", a1, w1},
		{"2", a2, w2},
		{"3", a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Float(); got != tt.want {
				t.Errorf("%q: TArg.Float() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_TArg_Float()

func Test_TArg_Int(t *testing.T) {
	a0, w0 := TArg(""), 0
	a1, w1 := TArg("123"), 123
	a2, w2 := TArg("-234"), -234
	a3, w3 := TArg("n.a."), 0

	tests := []struct {
		name string
		arg  TArg
		want int
	}{
		{"0", a0, w0},
		{"1", a1, w1},
		{"2", a2, w2},
		{"3", a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.Int(); got != tt.want {
				t.Errorf("%q: TArg.Int() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_TArg_Int()

func Test_TArg_String(t *testing.T) {
	a0, w0 := TArg(""), ""
	a1, w1 := TArg("12.3"), "12.3"
	a2, w2 := TArg("-234"), "-234"
	a3, w3 := TArg("n.a."), "n.a."

	tests := []struct {
		name string
		arg  TArg
		want string
	}{
		{"0", a0, w0},
		{"1", a1, w1},
		{"2", a2, w2},
		{"3", a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arg.String(); got != tt.want {
				t.Errorf("%q: TArg.String() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_TArg_String()

func Test_tOpt_Equal(t *testing.T) {
	a0, o0 := tOpt(""), tOpt("")
	a1, o1 := a0, tOpt("1")
	a2, o2 := tOpt("2"), tOpt("2")

	tests := []struct {
		name  string
		o     tOpt
		other tOpt
		want  bool
	}{
		{"0", a0, o0, true},
		{"1", a1, o1, false},
		{"2", a2, o2, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Equal(tt.other); got != tt.want {
				t.Errorf("%q: tOpt.Equal() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tOpt_Equal()

func Test_tOpt_String(t *testing.T) {
	a0, w0 := tOpt(""), ""
	a1, w1 := tOpt("12.3"), "12.3"
	a2, w2 := tOpt("-234"), "-234"
	a3, w3 := tOpt("n.a."), "n.a."

	tests := []struct {
		name string
		o    tOpt
		want string
	}{
		{"0", a0, w0},
		{"1", a1, w1},
		{"2", a2, w2},
		{"3", a3, w3},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("%q: tOpt.String() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tOpt_String()

func prepOptArgList() *tOptArgList {
	// This is the list set up by `init()` for testing purposes
	args := []string{
		"testingApplication",
		`-a`, // Flag option
		`-i`, // Error: intended with argument => ignored
		`--infile`, `config.in`,
		`--help`, // Flag option
	}
	return newOptArgList(args)
} // prepOptArgList()

func Test_tOptArg_Equal(t *testing.T) {
	oa0 := tOptArg{}
	oo0 := tOptArg{tOpt(""), TArg("")}
	oa1 := tOptArg{tOpt("one"), TArg("1")}
	oo1 := tOptArg{tOpt("one"), TArg("2")}
	oa2 := tOptArg{tOpt("oa2"), TArg("")}
	oo2 := tOptArg{tOpt("oo2"), TArg("")}

	tests := []struct {
		name  string
		oa    tOptArg
		other tOptArg
		want  bool
	}{
		{"0", oa0, oo0, true},
		{"1", oa1, oo1, false},
		{"2", oa2, oo2, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oa.Equal(tt.other); got != tt.want {
				t.Errorf("%q: tOptArg.Equal() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tOptArg_Equal()

func Test_tOptArg_String(t *testing.T) {
	oa0 := tOptArg{}
	ws0 := ``
	oa1 := tOptArg{tOpt("one"), TArg("")}
	ws1 := `["one": ""]`
	oa2 := tOptArg{tOpt("oa2"), TArg("2")}
	ws2 := `["oa2": "2"]`

	tests := []struct {
		name string
		oa   tOptArg
		want string
	}{
		{"0", oa0, ws0},
		{"1", oa1, ws1},
		{"2", oa2, ws2},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oa.String(); got != tt.want {
				t.Errorf("%q: tOptArg.String() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tOptArg_String()

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
	oal1 := prepOptArgList()
	w1 := `["a": ""]
["i": ""]
["-infile": "config.in"]
["-help": ""]
`
	tests := []struct {
		name    string
		oal     *tOptArgList
		wantStr string
	}{
		{"1", oal1, w1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := tt.oal.String(); gotStr != tt.wantStr {
				t.Errorf("%q: tOptArgList.String() =\n%s\n want \n%s",
					tt.name, gotStr, tt.wantStr)
			}
		})
	}
} // Test_tOptArgList_String()

/* _EoF_ */
