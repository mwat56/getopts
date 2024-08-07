/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// `tIterator` is a struct holding the actual commandline
	// options as well as the current iteration state.
	tIterator struct {
		// List of option arguments
		optArgs *tOptArgList

		// List of known/expected options and argument requirement
		expected *tExpectedOpts

		// Current index for iteration:
		index int
	}
)

// --------------------------------------------------------------------
// tIterator constructor

// `newIterator()` initialises a `tIterator` with the provided map
// of option arguments.
//
// This function treats the internal global variable `gIterator` as a global
// singleton, i.e. only the very first call to this function actually creates
// a new instance while all following calls simply update the global variable
// according to the provided `tOptArgMap` instance.
//
// Parameters:
//   - `aMap`: The list of commandline options and their respective argument.
//
// Returns:
//   - `*tIterator`: The iterator for `aList`.
func newIterator(aList *tOptArgList) *tIterator {
	if nil == gIterator {
		gIterator = &tIterator{
			optArgs: aList,
			// leave `expected` for lazy initialisation
			index: 0,
		}
	} else {
		gIterator.optArgs = aList
	}

	return gIterator
} // newIterator()

// --------------------------------------------------------------------
// tIterator methods

// `Next()` returns the next key-value pair in the iteration.
// It returns `false` if there are no more items.
//
// Returns:
//   - `rOpt`: The current option in the iteration.
//   - `rArg`: The current option's argument.
//   - `rMore`: Indicator for whether there are more options to process.
func (oi *tIterator) Next() (rOpt tOpt, rArg TArg, rMore bool) {
	if oi.index < len(*oi.optArgs) {
		oa := (*oi.optArgs)[oi.index]
		rOpt = oa.opt
		rArg = oa.arg
		oi.index++
		rMore = (oi.index < len(*oi.optArgs))
		//} else {	rMore = false
		// leave the other return values at their zero values
	}

	// if nil == oi.expected {
	// 	// This should never happen because this singleton
	// 	// instance actually IS the global iterator ...
	// 	oi.expected = gIterator.expected
	// }

	if !oi.expected.isValid(rOpt, rArg) {
		if rMore {
			// An unknown option or one without the required
			// argument is simply ignored and we continue by
			// going through the next iteration:
			rOpt, rArg, rMore = oi.Next()
		} else {
			rOpt = tOpt("")
			rArg = TArg("")
			// `rMore` remains false
		}
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

// `setPattern()` sets up the options pattern for the options iterator.
//
// The method sets the pattern for the iterator, which is used to
// determine the expected options and their required arguments. If this
// is the very first call, it initialises the private list of expected
// options (lazy initialisation). If there is already an existing
// pattern, the method ensures that the pattern is up-to-date.
//
// Note, that for the iterator to work properly, this method must be called
// at least once (usually by the `Getopts()` function) prior to the [Next]
// method. This requirement is satisfied by the `init()` function which
// passes the options pattern `h|-help` that is commonly used IRL. The
// actual options pattern used by `Getopts()` is passed by [Get] to
// method. So both, by default and by setting it explicitly, this method
// will be called internally, so there's no need to expose it publicly.
//
// Every call to this method will reset the iterator's index if the
// given options pattern is different from the one previously used.
//
// Parameters:
//   - `aPattern`: The new pattern to be used by the iterator.
//
// Returns:
//   - `*tIterator`: The iterator instance with the updated pattern.
func (oi *tIterator) setPattern(aPattern string) *tIterator {
	if nil == oi.expected {
		// This seems to be the very first call, hence we
		// need to initialise the list of expected options:
		oi.expected = newExpectedOpts(aPattern)
	} else if oi.expected.previous != aPattern {
		// Make sure the pattern is up-to-date:
		oi.expected = oi.expected.parse(aPattern)
		oi.index = 0
	}

	return oi
} // setPattern()

// --------------------------------------------------------------------

var (
	// Internal instance of the iterator to make it accessible for the
	// public `Get()` function. It is initialised automatically by
	// the getopts' `realInit()` function.
	gIterator *tIterator
)

/* _EoF_ */
