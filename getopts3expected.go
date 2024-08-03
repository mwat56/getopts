/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"slices"
	"strings"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (

	// A map of the _expected_ options indicating whether
	// it needs an argument to work.
	tArgBool map[tOpt]bool

	tExpectedArgs struct {
		// List of all _expected_ commandline options
		argBool tArgBool

		// A previously used options pattern
		previous string
	}
)

func newExpectedArgs() *tExpectedArgs {
	return &tExpectedArgs{
		argBool: make(tArgBool),
	}
} // newExpectedArgs()

// `needArgument()` checks if a given commandline option requires an argument.
//
// This function takes a command-line option name as input and returns
// a boolean value indicating whether the option requires an argument.
//
// Parameters:
//   - `aOpt`: The commandline option name to check.
//
// Returns:
//   - `bool`: Indicator for whether the option requires an argument.
func (ea tExpectedArgs) needArgument(aOpt tOpt) bool {
	result, ok := ea.argBool[aOpt]
	if !ok {
		return false
	}

	return result
} // needArgument()

// `parse()` parses the provided pattern and updates the expected arguments.
//
// Parameters:
//   - `aPattern`: The pattern declaring which commandline options to expect.
//
// Returns:
//   - `*tExpectedArgs`: The updated expected arguments.
func (ea *tExpectedArgs) parse(aPattern string) *tExpectedArgs {
	if aPattern == ea.previous {
		return ea
	}

	// Split the pattern string by `|` into a slice of option names
	opts := strings.Split(aPattern, `|`)
	slices.Sort(opts)

	oLen := len(opts)
	for i := 0; i < oLen; i++ {
		opt := opts[i]
		opt = string(opt[len(opt)-1])
		if ":" == opt {
			o := tOpt(opt[:len(opt)-1])
			ea.argBool[o] = true
		} else {
			ea.argBool[tOpt(opt)] = false
		}
	}
	ea.previous = aPattern

	return ea
} // parse()

var (
	// global/private instance of `tExpectedArgs`
	gExpectedOpts *tExpectedArgs
)

/* _EoF_ */
