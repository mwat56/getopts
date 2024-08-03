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

func Test_newOptIterator(t *testing.T) {

	tests := []struct {
		name string
		amap tArgList
		want *tIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newIterator(tt.amap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q: newOptIterator() = %v, want %v",
					tt.name, got, tt.want)
			}
		})
	}
} // TestNewOptIterator()

func Test_tOptIterator_Next(t *testing.T) {

	tests := []struct {
		name     string
		oi       *tIterator
		wantROpt tOpt
		wantRArg TArg
		wantROK  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotROpt, gotRArg, gotROK := tt.oi.Next()
			if gotROpt != tt.wantROpt {
				t.Errorf("%q: tOptIterator.Next() gotROpt = %v, want %v",
					tt.name, gotROpt, tt.wantROpt)
			}
			if gotRArg != tt.wantRArg {
				t.Errorf("%q: tOptIterator.Next() gotRArg = %v, want %v",
					tt.name, gotRArg, tt.wantRArg)
			}
			if gotROK != tt.wantROK {
				t.Errorf("%q: tOptIterator.Next() gotROK = %v, want %v",
					tt.name, gotROK, tt.wantROK)
			}
		})
	}
} // Test_tOptIterator_Next()

func Test_tOptIterator_Reset(t *testing.T) {

	tests := []struct {
		name string
		oi   *tIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.oi.Reset()
			if 0 != tt.oi.index {
				t.Errorf("%q: tOptIteratorReset.() gotROK = %v, want %v",
					tt.name, tt.oi.index, 0)

			}
		})
	}
} // Test_tOptIterator_Reset()
