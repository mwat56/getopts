/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import "testing"

//lint:file-ignore ST1017 - I prefer Yoda conditions

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
