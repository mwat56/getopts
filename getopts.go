/*
Copyright © 2024  M.Watermann, 10247 Berlin, Germany

			All rights reserved
		EMail : <support@mwat.de>
*/

package getopts

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

//lint:file-ignore ST1017 - I prefer Yoda conditions

type (
	IHelpShower interface {
		// `ShowHelp` is supposed to show some helpful information
		// to the user if a help request was triggered by the
		// commandline options `-h` or `--help`.
		// If the function returns a non `nil` error value the getopts
		// processing is aborted.
		ShowHelp() error
	}
)

// `HelpShower` implements the `IHelpShower` interface to provide some
// helpful information to the user if the help request was triggered
// by the commandline options `-h` or `--help`.
//
// If the `ShowHelp()` function returns a non `nil` error value the
// getopts processing is aborted.
//
// Note: This variable must be setup before the [Get] function is called.
var HelpShower IHelpShower

// --------------------------------------------------------------------
// Internal functions

var (
	// Internal flag signalling whether we're in testing/debugging mode:
	gSomeTestsAreRunning bool

	// Barrier for concurrent tests
	gMtx sync.Mutex
)

// `init()` automatically initialises the command-line argument parser
// and sets up the associated parser for the arguments.
func init() {
	args := os.Args

	//
	//TODO: split os.Args into option groups using a pattern like:
	// 	([\p{L}]+(\s+((\-+[\p{L}]+)\s*)+)+)
	//

	// This env var is set by `go -test`:
	if "true" == os.Getenv("testing") {
		// Apparently we're testing or debugging
		gSomeTestsAreRunning = true

		args = []string{
			"testingApplication",
			`-a`, // Flag option
			`-i`, // Error: intended with argument => ignored
			`--infile`, `config.in`,
			`--help`, // Flag option
		}
	}

	realInit(args)
} // init()

// `(realInit)` initialises the commandline argument parser.
//
// The function checks if at least one argument is passed and then
// initialises an arguments map to store the command-line arguments
// and options. It then gets the command-line arguments without the
// app's path/name, iterates through them, and adds the argument and
// their corresponding options to the `arguments` map. Finally, it
// initialises the internal iterator `gIterator` with the arguments map.
//
// Parameters:
//   - `aArgList`: A list of commandline options and arguments.
func realInit(aArgList []string) {
	if gSomeTestsAreRunning {
		gMtx.Lock()
		defer gMtx.Unlock()
	}
	oal := newOptArgList(aArgList)

	// Set up the global/internal iterator:
	newIterator(oal)
} // realInit()

// --------------------------------------------------------------------
// public functions

// `Get()` retrieves the next commandline option and its argument.
//
// The function uses an internal iterator to retrieve the next option and
// its argument. If the iterator has no more items, it returns `false` for
// `rOK`. The retrieved option and argument are returned as `rOpt` and
// `rArg` respectively.
//
// The `aPattern` parameter is used to set up the internal iterator
// to know which options to expect/accept.
//
// Parameters:
//   - `aPattern`: The pattern declaring which commandline options to expect.
//
// Returns:
//   - `rOpt`: The current option in the iteration.
//   - `rArg`: The option's argument in the iteration.
//   - `rMore`: Indicator for whether there are more options to come.
func Get(aPattern string) (rOpt string, rArg TArg, rMore bool) {
	o, rArg, rMore := gIterator.setPattern(aPattern).Next()
	if `` == o {
		// This might happen if the last commandline option is
		// invalid (i.e. not defined in `aPattern` or missing
		// its required argument.)
		rOpt = string(`?`)
	} else {
		switch o {
		case `h`, `-help`:
			if nil != HelpShower {
				if err := HelpShower.ShowHelp(); nil != err {
					// Perhaps somebody needs time?
					runtime.Gosched()
					// And here we go ...
					log.Fatalln(err.Error())
				}
			}
		}
		rOpt = string(o)
	}

	return
} // Get()

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
		o, a, more := Get(aPattern)
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
