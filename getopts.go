/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"os"
	"slices"
	"strconv"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// `TOpt` represents an option of a command line argument
	TOpt string
)

// --------------------------------------------------------------------
// TOpt methods

// `Bool()` returns the option's value as a boolean value.
//
// `0`, `f`, `F`, `n`, and `N` are considered `false` while
// `1`, `t`, `T`, `y`, `Y`, `j`, `J`, `o`, `O` are considered `true`.
//
// Note that the mere existence of a commandline argument automatically
// makes its option `true` by default.
//
// This method actually checks only the first character of the option's
// value so one can write e.g. "false" or "NO" (for a `false` result),
// or "True" or "yes" (for a `true` result).
//
// Returns:
// - `bool`:The options's value as a Boolean.
func (o TOpt) Bool() bool {
	if 0 < len(o) {
		switch o[:1] {
		case `0`, `f`, `F`, `n`, `N`:
			return false

		// case `1`, `t`, `T`, `y`, `Y`, `j`, `J`, `o`, `O`:
		// 	// True, Yes (English), Ja (German), Oui (French)`
		default:
			return true
		}
	}

	return true
} // Bool()

// `Float()` returns the option's value as a 64bit floating point.
//
// If the string is well-formed and near a valid floating point number,
// [Float] returns the nearest floating point number rounded using
// IEEE754 unbiased rounding.
//
// In case the option can't be converted to a float the method's result
// will be the value `float64(0.0)`.
//
// Returns:
// - `float64`: The options's value as a 64bit floating point.
func (o TOpt) Float() float64 {
	if "" == string(o) {
		return float64(0.0)
	}

	if f64, err := strconv.ParseFloat(string(o), 64); (nil == err) && (f64 == f64) {
		// for NaN the inequality comparison with itself returns true
		return f64
	}

	return float64(0.0)
} // Float()

// `Int()` returns the option's value as an integer.
//
// In case the option can't be converted to an integer the method's result
// will be the value `int(0)` (zero).
//
// Returns:
// - `int`: The options's value as an integer.
func (o TOpt) Int() int {
	if "" == string(o) {
		return 0
	}

	if i64, err := strconv.ParseInt(string(o), 10, 0); nil == err {
		return int(i64)
	}

	return int(0)
} // Int()

// `String()` returns the option's value as a string.
//
// Returns:
// - `string`: The value of `aKey` as a string.
func (o TOpt) String() string {
	return string(o)
} // String()

// --------------------------------------------------------------------
// Internal types and their methods

type (
	// The internal map holding the command line arguments and options.
	tArgs map[string]TOpt

	// tArgIterator is a struct that holds the map and
	// the current iteration state.
	tArgIterator struct {
		data  tArgs
		keys  []string // Slice to hold keys of the map for iteration
		index int      // Current index for iteration
	}
)

// `newArgIterator()` initialises a `tArgIterator` with the provided map.
//
// Returns:
// - `*tArgIterator`: The new iterator for `aMap`.
func newArgIterator(aMap tArgs) *tArgIterator {
	keys := make([]string, 0, len(aMap))
	for k := range aMap {
		keys = append(keys, k)
	}

	// Sort the slice
	slices.Sort(keys)
	//

	return &tArgIterator{
		data:  aMap,
		keys:  keys,
		index: 0,
	}
} // newArgIterator()

/*
// xNext returns the next key-value pair in the iteration.
// It returns false if there are no more items.
*/
// `Next()` returns the next key-value pair in the iteration.
// It returns `false` if there are no more items.
//
// Returns:
// - rArg: The current key in the iteration.
// - rOpt: The corresponding value for the current key.
// - rOK: A boolean indicating whether there are more items to iterate over.
func (m *tArgIterator) Next() (rArg string, rOpt TOpt, rOK bool) {
	if m.index < len(m.keys) {
		rArg = m.keys[m.index]
		rOpt = m.data[rArg]
		m.index++
		rOK = true
	} else {
		rOK = false
	}

	return
} // Next()

// `Reset()` resets the iterator to the beginning.
//
// This method resets the iterator's current index to 0, effectively
// starting the iteration from the beginning.
func (m *tArgIterator) Reset() {
	m.index = 0
} // Reset()

var (
	gIterator *tArgIterator
	gPattern  []string
)

// --------------------------------------------------------------------
// Internal functions

// `init()` automatically initialises the command-line argument parser
// and sets up the associated parser for the arguments.
func init() {
	realInit()
} // init()

// `(realInit)` initialises the command-line argument parser.
//
// The function checks if at least one argument is passed and then initialises
// an `arguments` map to store the command-line arguments and options. It
// then gets the command-line arguments without the app's path/name,
// iterates through them, and adds the argument and their corresponding
// options to the `arguments` map. Finally, it initialises the iterator
// `gIterator` with the `arguments` map.
func realInit() {
	// Check if at least one argument is passed
	if 2 > len(os.Args) {
		return
	}

	arguments := make(tArgs)
	// Get the commandline arguments w/o the app's path/name:
	args := append(os.Args[1:], "")

	// Ignore the peek-ahead dummy added at the end:
	aLen := len(args) - 1

	for i := 0; i < aLen; i++ {
		a := args[i]
		if 1 > len(a) {
			continue
		}
		if '-' == a[0] {
			a = a[1:]
			p := args[i+1] // peek ahead
			// If there's another value that is not an argument,
			// ignore it:
			if (0 < len(p)) && ('-' == p[0]) {
				p = ""
			} else {
				i++
			}
			arguments[a] = TOpt(p)
		}
	}

	// Initialise the iterator
	gIterator = newArgIterator(arguments)
} // realInit()

func prepPattern(aPattern string) {

}

// --------------------------------------------------------------------
// public functions

func Getopts(aPattern string) (rArg string, rOpt TOpt, rOK bool) {
	if nil == gPattern {
		prepPattern(aPattern)
	}

	rArg, rOpt, rOK = gIterator.Next()

	return
} // Getopts()

func Xmain() {
	for {
		a, o, n := Getopts("")
		switch a {
		case "b":
			b := o.Bool()
		case "f":
			f := o.Float()
		case "i":
			i := o.Int()
		case "s":
			s := o.String()

		}
		if !n {
			break
		}
	}

} // main()

/* _EoF_ */
