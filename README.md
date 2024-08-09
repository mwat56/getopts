# GetOpts

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/mwat56/getopts?status.svg)](https://godoc.org/github.com/mwat56/getopts)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/getopts)](https://goreportcard.com/report/github.com/mwat56/getopts)
[![Issues](https://img.shields.io/github/issues/mwat56/getopts.svg)](https://github.com/mwat56/getopts/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/getopts.svg)](https://github.com/mwat56/getopts/)
[![Tag](https://img.shields.io/github/tag/mwat56/getopts.svg)](https://github.com/mwat56/getopts/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/getopts/blob/main/_demo/demo.go)
[![License](https://img.shields.io/github/mwat56/getopts.svg)](https://github.com/mwat56/getopts/blob/main/LICENSE)

- [GetOpts](#getopts)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
		- [Options pattern](#options-pattern)
			- [Bash usage](#bash-usage)
			- [Go usage](#go-usage)
	- [Libraries](#libraries)
	- [Licence](#licence)

----

## Purpose

    //TODO

## Installation

You can use `Go` to install this package for you:

```bash
go get -u github.com/mwat56/getopts
```

## Usage

Since you're interested in this module you probably know the `getopts` functionality as used in, for example, shell scripts etc. This module tries to mimic that. However, there are some small differences in the [pattern](#options-pattern) used to describe the expected commandline arguments.

As an example here a simple use case:

```go
func MySetup(aPattern string) {
	var (
		b bool
		f float64
		i int
		o string
		s string
	)

	// Loop through all available options:
	for {
		opt, arg, more := Get(aPattern)
		switch opt {
		case "b":
			b = arg.Bool()
		case "f":
			f = arg.Float()
		case "i":
			i = arg.Int()
		case "s":
			s = arg.String()
		default:
			o = opt
		}
		if !more {
			// No more options available
			break
		}
	}

	fmt.Printf("Bool: %t, Float: %f, Int: %d, String: %q, other: %v",
		b, f, i, s, o)
} // MySetup()
```

Of course, you would need to give a valid [`pattern`](#go-usage). And you would assign the respective options to your own configuration variables (instead of the locally declared dummies in the example above).

### Options pattern

Here comes a brief comparison between the use of `getopts` in a shell script and in Go.

#### Bash usage

In a `bash` shell script, the use would look something like this:

```bash
while getopts ":a:c:hjl:p:qs:v:" opt; do
	case ${opt} in
	a)	# use the `a` argument's option
		echo " -a option: ${OPTARG}"
		;;
	c)	# use the `c` argument's option
		echo " -c option: ${OPTARG}"
		;;
	h)	# handle the `h` argument's existence
		echo " -h: ${opt}"
		;;
	j)	# handle the `j` argument's existence
		echo " -j: ${opt}"
		;;
	l)	# use the `p` argument's option
		echo " -p option: ${OPTARG}"
		;;
	p)	# use the `p` argument's option
		echo " -p option: ${OPTARG}"
		;;
	q)	# handle the `q` argument's existence
		echo " -q: ${opt}"
		;;
	s)	# use the `s` argument's option
		echo " -s option: ${OPTARG}"
		;;
	v)	# use the `v` argument's option
		echo " -v option: ${OPTARG}"
		;;
	\?)	# unknown option
		echo -e "\n\tignoring unknown option: -${opt}=${OPTARG}\n" >&2
		;;
	:)	# invalid argument
		echo "Option -${opt} requires an argument." >&2
		;;
	esac
done
```

As can be seen above, the pattern used is this: `:a:c:hjl:p:qs:v:`. The expected options are given by their respective name. A following `:` (colon) signals that the option expects an argument following it. Let's break it down:

1. The leading colon (`:`) at the beginning of the string enables silent error reporting mode. In this mode, `getopts` handles errors internally without printing error messages.
2. The remaining characters define the valid options:
   * `a`: - Option `-a` requires an argument
   * `c`: - Option `-c` requires an argument
   * `h`: - Option `-h` does not require an argument
   * `j`: - Option `-j` does not require an argument
   * `l`: - Option `-l` requires an argument
   * `p`: - Option `-p` requires an argument
   * `q`: - Option `-q` does not require an argument
   * `s`: - Option `-s` requires an argument
   * `v`: - Option `-v` requires an argument

So, this `getopts` argument specifies that the script accepts the following options:

   * `-a`, `-c`, `-l`, `-p`, `-s`, and `-v`: These options _require an argument_. When using these options, you must provide a value immediately after the option letter.
   * `-h`, `-j`, and `-q`: These are _flag options_ that do not require an argument. They can be used to toggle certain behaviours or settings in the script.

#### Go usage

Now, Go does not support a `while` loop directly but as can be seen [above](#usage) it can be simply build by an `for{ ... }` loop that runs as long as it isn't broken.

The _pattern_ to use look similar but not identical:

- Shell: `:a:c:hjl:p:qs:v:`
- Go: `a:|c:|h|j|l:|p:|q|s:|v:`

A leading colon in the _pattern_ is not needed here because any problems are handled internally anyway. One common problem, for example, is giving an option on the commandline that requires an argument (e.g. a filename or a certain value) without providing that argument. This Go implementation of `getopts()` simply ignores such option, and it's up to the developer to decide what to do if the option wasn't provided by the app user (which, BTW, a developer has to do anyway).

While the *nix _getopts_ allows only for single letter options, we want to be able to work with long options like `--help` as well. Hence we need a separator between the options which is here the pipe symbol `|`. So a pattern for this Go implementation could look like this:

   - `a|i:|-input:|h|-help|o:|-output:|q|v`

This would handle a commandline with options like

```bash
$> myprog -a -i inFile1 --input inFile1 -o outFile1 --output outFile2 -q -v --help
```

It's up to the developer to decide how to handle such settings: Here are both, the short and the long form, options used for the input and output file names. And there are both, the `-q` (_quiet_) and `-v` (_verbose_), flags used. And shouldn't `myprog` be terminated if help was requested?

## Libraries

The following external libraries were used building `getopts`:

* No external modules were used apart from Go's standard library.

## Licence

	Copyright © 2024 M.Watermann, 10247 Berlin, Germany
			All rights reserved
		    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
[![GFDL](https://www.gnu.org/graphics/gfdl-logo-tiny.png)](http://www.gnu.org/copyleft/fdl.html)
