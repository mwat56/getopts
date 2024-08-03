/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"fmt"
	"os"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

// --------------------------------------------------------------------
// Internal functions

// `init()` automatically initialises the command-line argument parser
// and sets up the associated parser for the arguments.
func init() {
	// This env var is set by `go -test`
	if "true" == os.Getenv("testing") {
		// This function is disabled during testing.
		return
	}
	realInit(os.Args)
} // init()

// `(realInit)` initialises the command-line argument parser.
//
// The function checks if at least one argument is passed and then
// initialises an arguments map to store the command-line arguments
// and options. It then gets the command-line arguments without the
// app's path/name, iterates through them, and adds the argument and
// their corresponding options to the `arguments` map. Finally, it
// initialises the internal iterator `gIterator` with the arguments map.
func realInit(aArgList []string) {
	arguments := make(tArgList)

	// Check if at least one argument (apart from the running
	// app's path/filename) is passed:
	if 2 > len(aArgList) {
		arguments[tOpt("0")] = TArg("0")

		// Set up the global/internal iterator to avoid
		// NIL pointer problems along the way
		gIterator = newIterator(arguments)
		return
	}

	// Get the commandline arguments w/o the app's path/name
	// but an added (empty) argument to allow for a peek ahead:
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
				// leave it as subject of the next loop step
				p = ""
			} else {
				i++
			}
			arguments[tOpt(o)] = TArg(p)
		}
	}

	// Set up the global/internal iterator:
	gIterator = newIterator(arguments)
} // realInit()

// --------------------------------------------------------------------
// public functions

// `Getopts()` retrieves the next command-line option and its argument.
//
// The function uses an private iterator to retrieve the next option and
// its argument. If the iterator has no more items, it returns `false` for
// `rOK`. The retrieved option and argument are returned as `rOpt` and
// `rArg` respectively.
//
// The `aPattern` parameter is used to initialise the iterator if it has
// not been initialised yet.
//
// Parameters:
//   - `aPattern`: The pattern declaring which commandline options to expect.
//
// Returns:
//   - `rOpt`: The current option in the iteration.
//   - `rArg`: The current option's argument in the iteration.
//   - `rOK`: Indicator for whether there are more options to come.
func Getopts(aPattern string) (rOpt string, rArg TArg, rOK bool) {
	var ea *tExpectedArgs

	if nil == gExpectedOpts {
		ea = newExpectedArgs()
		gExpectedOpts = ea
	} else {
		ea = gExpectedOpts
	}
	ea.parse(aPattern)

	o, a, ok := gIterator.Next()
	rOpt, rArg, rOK = string(o), a, ok
	if ea.needArgument(o) {
		if "" == string(rArg) {
			rOpt = fmt.Sprintf("!> option %q requires an argument <!", rOpt)
		}
	}

	return
} // Getopts()

func MySetup(aPattern string) {
	var (
		b   bool
		f   float64
		i   int
		s   string
		opt string
	)

	// Now loop through all available options:
	for {
		o, a, more := Getopts(aPattern)
		switch o {
		case "b":
			b = a.Bool()
		case "f":
			f = a.Float()
		case "i":
			i = a.Int()
		case "s":
			s = a.String()
		default:
			opt = o
		}
		if !more {
			// No more options available
			break
		}
	}
	fmt.Printf("Bool: %t, Float: %f, Int: %d, String: %q, other: %v",
		b, f, i, s, opt)
} // MySetup()

/* _EoF_ */
