/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"slices"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// A list of options
	tOptList []tOpt

	// `tIterator` is a struct holding the actual commandline
	// options and the current iteration state.
	tIterator struct {
		args  tArgList // Slice of option arguments
		opts  tOptList // Slice of option names of the map for iteration
		index int      // Current index for iteration
	}
)

// --------------------------------------------------------------------
// Constructor function

// `newIterator()` initialises a `tArgIterator` with the provided
// map of option arguments.
//
// Parameters:
//   - `aMap`: The list of commandline options and their respective argument.
//
// Returns:
//   - `*tArgIterator`: The new iterator for `aMap`.
func newIterator(aMap tArgList) *tIterator {
	options := make(tOptList, 0, len(aMap))
	for o := range aMap {
		options = append(options, o)
	}
	if 1 < len(options) {
		// Sort the slice:
		slices.Sort(options)
	}

	return &tIterator{
		args:  aMap,
		opts:  options,
		index: 0,
	}
} // newOptIterator()

// --------------------------------------------------------------------
// tOptIterator methods

// `Next()` returns the next key-value pair in the iteration.
// It returns `false` if there are no more items.
//
// Returns:
//   - `rOpt`: The current option in the iteration.
//   - `rArg`: The current option's argument in the iteration.
//   - `rOK`: Indicator for whether there are more options to iterate over.
func (oi *tIterator) Next() (rOpt tOpt, rArg TArg, rOK bool) {
	if oi.index < len(oi.opts) {
		rOpt = oi.opts[oi.index]
		rArg = oi.args[rOpt]
		oi.index++
		rOK = true
	} else {
		rOK = false
		// leave the other return values at their default
	}

	return
} // Next()

// `Reset()` resets the iterator to the beginning.
//
// This method resets the iterator's current index to `0` (zero),
// effectively starting the iteration from the beginning.
func (oi *tIterator) Reset() {
	oi.index = 0
} // Reset()

// --------------------------------------------------------------------

var (
	// Global/internal instance of the iterator to make it
	// accessible for the public `getopts` function. It is
	// initialised automatically by getopts' `init()` function.
	gIterator *tIterator
)

/* _EoF_ */
