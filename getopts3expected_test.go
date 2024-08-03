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

func Test_newExpectedArgs(t *testing.T) {
	tests := []struct {
		name string
		want *tExpectedArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newExpectedArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q: newExpectedArgs() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_newExpectedArgs()

func Test_tExpectedArgs_needArgument(t *testing.T) {

	tests := []struct {
		name string
		ea   tExpectedArgs
		opt  tOpt
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ea.needArgument(tt.opt); got != tt.want {
				t.Errorf("%q: tExpectedArgs.argNeeded() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tExpectedArgs_needArgument()

func Test_tExpectedArgs_parse(t *testing.T) {

	tests := []struct {
		name    string
		ea      *tExpectedArgs
		pattern string
		want    *tExpectedArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ea.parse(tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q: tExpectedArgs.parse() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // Test_tExpectedArgs_parse()
