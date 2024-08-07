/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"strings"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	// A map of the _expected_ options indicating whether
	// it needs an argument to work.
	tArgBool map[tOpt]bool

	// A map of the known/expected options and argument requirement
	tExpectedOpts struct {
		// List of all _expected_ commandline options
		argBool tArgBool

		// A previously used options pattern
		previous string
	}
)

// `newExpectedOpts()` sets up a instance of `tExpectedOpts`.
//
// This function is a first step to manage the expected commandline
// options and their respective arguments.
//
// Parameters:
//   - `aPattern`: The pattern declaring which commandline options to expect.
//
// Returns:
//   - `*tExpectedOpts`: The requested `tExpectedArgs` instance.
func newExpectedOpts(aPattern string) *tExpectedOpts {
	var eo *tExpectedOpts

	if nil == gIterator {
		eo = &tExpectedOpts{
			argBool: make(tArgBool),
		}
	} else {
		if nil == gIterator.expected {
			eo = &tExpectedOpts{
				argBool: make(tArgBool),
			}
			gIterator.expected = eo.parse(aPattern)

			return eo
		}

		eo = gIterator.expected
		if eo.previous == aPattern {
			return eo
		}
	}

	return eo.parse(aPattern)
} // newExpectedOpts()

// `isValid()` checks if a given commandline option is recognised.
//
// This function takes a commandline option name as input and returns
// a boolean value indicating whether the option requires an argument.
//
// Parameters:
//   - `aOpt`: The commandline option name to check.
//   - `rArg`: The option's argument to check.
//
// Returns:
//   - `bool`: Indicator for whether the option is recognised.
func (eo tExpectedOpts) isValid(aOpt tOpt, rArg TArg) bool {
	needArg, valid := eo.argBool[aOpt]
	if needArg && ("" == string(rArg)) {
		return false
	}

	return valid
} // isValid()

// `parse()` parses the provided pattern and updates the expected arguments.
//
// NOTE: If the given `aPattern` is empty, then the pattern `h|-help` will
// be used which usually triggers a help request and the termination of
// the running application.
//
// Parameters:
//   - `aPattern`: The pattern declaring which commandline options to expect.
//
// Returns:
//   - `*tExpectedArgs`: The updated expected arguments.
func (eo *tExpectedOpts) parse(aPattern string) *tExpectedOpts {
	if "" == aPattern {
		aPattern = "h|-help"
	}

	// Shortcut to avoid unnecessary parsing:
	if aPattern == eo.previous {
		return eo
	}

	// Reset the map to remove all previous entries
	clear(eo.argBool)

	// Split the pattern string by `|` into a slice of options
	// and their arguments:
	optargs := strings.Split(aPattern, `|`)
	// The order of the options reflects the order in the given
	// pattern, but not order of options on the commandline.

	oLen := len(optargs)
	for i := 0; i < oLen; i++ {
		opt := optargs[i]
		if "" == opt {
			// leading/trailing or orphaned pipe separator
			continue
		}

		// Look for leading colons and spaces to determine whether
		// to remove leasing garbage:
		pos, optL := 0, len(opt)
		for (pos < optL) && ((`:` == string(opt[pos])) || (` ` == string(opt[pos]))) {
			pos++
		}
		if 0 < pos {
			if pos == optL {
				continue // ignore empty option
			}
			opt = opt[pos:]
			optL = len(opt)
		}

		// Now, look for trailing colons and spaces to determine whether
		// the option requires an argument:
		pos = optL
		needArg := false
	argLoop:
		for 0 < pos {
			// for (0 < pos) && ((`:` == string(opt[pos-1])) || (` ` == string(opt[pos-1]))) {
			switch string(opt[pos-1]) {
			case `:`:
				needArg = true
				pos--

			case ` `:
				pos--

			default:
				break argLoop
			}
		}

		if needArg || (pos < optL) {
			if 0 == pos {
				continue // ignore empty option
			}
			opt = opt[:pos]
		}
		eo.argBool[tOpt(opt)] = needArg
	}
	// Save the pattern for a possible future call:
	eo.previous = aPattern

	return eo
} // parse()

/* _EoF_ */
