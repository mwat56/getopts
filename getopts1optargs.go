/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"fmt"
	"strconv"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// `TArg` represents an argument of a commandline option.
	TArg string

	// `tOpt` represents a commandline option.
	// Type definition provided for better readability and clarity.
	tOpt string

	// `tOptarg` is a key/value pair of a single commandline option
	// and its (potentially empty) argument.
	tOptArg struct {
		opt tOpt
		arg TArg
	}

	// `tOptArgList` is a list of option/argument pairs.
	//
	// A slice holding the _actual_ commandline options with their argument.
	// Type definition provided for better readability and clarity.
	tOptArgList []tOptArg
)

// --------------------------------------------------------------------
// TArg methods

// `Bool()` returns the argument's value as a boolean value.
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
// - `bool`: The argument's value as a Boolean.
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

// `Equal()` checks if this argument is equal to another one.
//
// Parameters:
//   - `aArg`: The argument to compare with this one.
//
// Returns:
//   - `bool`: An indicator for whether this argument equals `aArg`.
func (a TArg) Equal(aArg TArg) bool {
	return a == aArg
} // Equal()

// `Float()` returns the argument's value as a 64bit floating point.
//
// If the string is well-formed and near a valid floating point number,
// [Float] returns the nearest floating point number rounded using
// IEEE754 unbiased rounding.
//
// In case the option can't be converted to a float the method's result
// will be the value `float64(0.0)`.
//
// Returns:
// - `float64`: The argument's value as a 64bit floating point.
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

// `Int()` returns the argument's value as an integer.
//
// In case the option can't be converted to an integer the method's
// result will be the value `int(0)` (zero).
//
// Returns:
// - `int`: The argument's value as an integer.
func (a TArg) Int() int {
	if "" == string(a) {
		return int(0)
	}

	if i64, err := strconv.ParseInt(string(a), 10, 0); nil == err {
		return int(i64)
	}

	return int(0)
} // Int()

// `String()` returns a stringified version of the argument.
//
// Note: This is mainly for debugging purposes and has no real life use.
//
// Returns:
//   - `string`: The stringified version of the current argument.
func (a TArg) String() string {
	return string(a)
} // String()

// --------------------------------------------------------------------
// tOpt methods

// `Equal()` checks if this option is equal to another one.
//
// Parameters:
//   - `aOpt`: The option to compare with this one.
//
// Returns:
//   - `bool`: An indicator for whether this option equals `aOpt`.
func (o tOpt) Equal(aOpt tOpt) bool {
	return o == aOpt
} // Equal()

// `String()` returns a stringified version of the option.
//
// Note: This is mainly for debugging purposes and has no real life use.
//
// Returns:
//   - `string`: The stringified version of the current option.
func (o tOpt) String() string {
	return string(o)
} // String()

// --------------------------------------------------------------------

// `Equal()` checks if this option/argument is equal to another one.
//
// Parameters:
//   - `aOpt`: The option/argument to compare with this one.
//
// Returns:
//   - `bool`: Indicator for whether this option/argument equals `aOptArg`.
func (oa tOptArg) Equal(aOptArg tOptArg) bool {
	if oa.opt != aOptArg.opt {
		return false
	}
	return oa.arg == aOptArg.arg
} // Equal()

// `String()` returns a stringified version of the option/argument.
//
// Note: This is mainly for debugging purposes and has no real life use.
//
// Returns:
//   - `string`: The stringified version of the current option/argument.
func (oa tOptArg) String() string {
	return fmt.Sprintf("[%q: %q]", oa.opt, oa.arg)
} // String()

// --------------------------------------------------------------------
// tOptArgList Constructor

// `newOptArgList()` creates a new option/argument list based on `aArgList`.
//
// The function checks whether at least one argument is passed and then
// initialises an options/arguments list to store the current commandline
// options and arguments passed as `aArgList`. If the provided list of
// commandline options is empty (apart from the name of the running
// application), the function sets up the standard options "-h" and
// "--help" with an empty argument (such making them a `flag option`).
//
// Note that only the options's name is stored, i.e. the leading hyphen used
// on the actual commandline is removed. This means that so-called `short
// options` will be only a single letter and the `long options` will have
// their first hyphen removed. For example, the `short option` of "-h" will
// be stored as "h", and the `long option` of "--help" will be stored as
// "-help".
//
// Parameters:
//   - `aArgList`: A list of commandline options and arguments.
//
// Returns:
//   - `*tArgList`: A pointer to the newly created argument list.
func newOptArgList(aArgList []string) *tOptArgList {
	// Previously we used a map to store the key/value pairs. However,
	// because a map can not keep it's assigned order option/argument
	// pairs could only accessed in a random order. Hence we switched
	// to using a slice instead to keep the order of the commandline
	// options intact.
	oal := make(tOptArgList, 0, len(aArgList))

	if 1 >= len(aArgList) {
		var empty TArg
		// Set up some standard default options:
		oal = append(oal, tOptArg{tOpt(`h`), empty})
		oal = append(oal, tOptArg{tOpt(`-help`), empty})
		// nothing more to do here:
		return &oal
	}

	// Get the commandline arguments without the app's path/name
	// but with an added (empty) argument to allow for a peek ahead
	// without any range problems when we happen to process the
	// very last real option or argument.
	optList := append(aArgList[1:], "")

	// Exclude the peek-ahead dummy added at the end:
	oLen := len(optList) - 1

	for i := 0; i < oLen; i++ {
		o := optList[i]
		if len(o) < 2 {
			// We expect at least `-o` i.e. two characters.
			continue
		}
		if '-' == o[0] {
			// It's actually an option (not
			// an unexpected argument)
			o = o[1:]

			//
			//TODO: should we check for `--` as stdOut indicator?
			//

			p := optList[i+1] // peek ahead
			// If there's another value that is not
			// an argument, ignore it here:
			if (0 < len(p)) && ('-' == p[0]) {
				// leave it for the next loop step
				p = ""
			} else {
				i++
			}

			// This might assign an empty argument to the option,
			// a situation which will be handled by the iterator's
			// `Next()` method by checking whether an argument is
			// actually required according to the options pattern.
			oal = append(oal, tOptArg{tOpt(o), TArg(p)})
		}
	}

	return &oal
} // newOptArgList()

// --------------------------------------------------------------------
// tOptArgList methods

// `Equal()` checks if this option argument list is equal to another
// one: the same length and all elements equal.
//
// If the lengths are different, [Equal] returns `false`. Otherwise, the
// elements are compared in increasing index order, and the comparison
// stops at the first unequal pair.
//
// Parameters:
//   - `aList`: The slice to compare with this one.
//
// Returns:
//   - `bool`: An indicator for whether this slice is equal to `aList`.
func (oal tOptArgList) Equal(aList tOptArgList) bool {
	if len(oal) != len(aList) {
		return false
	}

	for idx, oa := range oal {
		if !oa.Equal(aList[idx]) {
			return false
		}
	}

	return true
} // Equal()

// `String()` returns a stringified version of the option argument list
//
// Note: This is mainly for debugging purposes and has no real life use.
//
// Returns:
//   - `string`: The stringified version of the current options list.
func (oal tOptArgList) String() (rStr string) {
	for _, oa := range oal {
		rStr += fmt.Sprintf("%s\n", oa)
	}

	return
} // String()

/* _EoF_ */
