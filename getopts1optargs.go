/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"strconv"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// `TArg` represents an argument of a commandline option.
	TArg string

	// `tOpt` represents a commandline option.
	tOpt string

	// A map holding the _actual_ commandline options and their arguments.
	tArgList map[tOpt]TArg
)

// --------------------------------------------------------------------
// TArg methods

// `Bool()` returns an option's argument as a boolean value.
//
// `0`, `f`, `F`, `n`, and `N` are considered `false` while
// `1`, `t`, `T`, `y`, `Y`, `j`, `J`, `o`, `O` are considered `true`.
//
// This method actually checks only the first character of a option's
// argument so one can write e.g. "false" or "NO" (for a `false` result),
// or "True" or "yes" (for a `true` result).
//
// If the argument is empty or holds any other value than the ones
// mentioned above, then the method's result will be `false`.
//
// Returns:
// - `bool`: The options's argument as a Boolean.
func (a TArg) Bool() bool {
	if 0 < len(a) {
		switch a[:1] {
		case `0`, `f`, `F`, `n`, `N`:
			return false

		case `1`, `t`, `T`, `y`, `Y`, `j`, `J`, `o`, `O`:
			// True, Yes (English), Ja (German), Oui (French)`
		}
	}

	return false
} // Bool()

// `Float()` returns an option's argument as a 64bit floating point.
//
// If the string is well-formed and near a valid floating point number,
// [Float] returns the nearest floating point number rounded using
// IEEE754 unbiased rounding.
//
// In case the option can't be converted to a float the method's result
// will be the value `float64(0.0)`.
//
// Returns:
// - `float64`: A options's argument as a 64bit floating point.
func (a TArg) Float() float64 {
	if "" == string(a) {
		return float64(0.0)
	}

	if f64, err := strconv.ParseFloat(string(a), 64); (nil == err) && (f64 == f64) {
		// for NaN the inequality comparison with itself returns true
		return f64
	}

	return float64(0.0)
} // Float()

// `Int()` returns an option's argument as an integer.
//
// In case the option can't be converted to an integer the method's result
// will be the value `int(0)` (zero).
//
// Returns:
// - `int`: An option's argument as an integer.
func (a TArg) Int() int {
	if "" == string(a) {
		return 0
	}

	if i64, err := strconv.ParseInt(string(a), 10, 0); nil == err {
		return int(i64)
	}

	return int(0)
} // Int()

// `String()` returns an option's argument as a string.
//
// Returns:
// - `string`: An option's argument as a string.
func (a TArg) String() string {
	return string(a)
} // String()

// --------------------------------------------------------------------
// tOpt methods

func (o tOpt) String() string {
	return string(o)
} // String()

/* _EoF_ */
