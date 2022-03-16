// Copyright (C) 2022 Leonid Maslakov.

// This file is part of Highlight.

// Highlight is free software: you can redistribute it
// and/or modify it under the terms of the
// GNU Affero Public License as published by the
// Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.

// Highlight is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero Public License for more details.

// You should have received a copy of the GNU Affero Public License along with Highlight.
// If not, see <https://www.gnu.org/licenses/>.

package highlight

import (
	"strings"
)

// Python processes python source code (*.py files).
// Read more: https://docs.python.org/3/reference/index.html
//
// Supported keywords (const StyleKeyword):
//   False
//   None
//   True
//   and
//   as
//   assert
//   async
//   await
//   break
//   class
//   continue
//   def
//   del
//   elif
//   else
//   except
//   finally
//   for
//   from
//   global
//   if
//   import
//   in
//   is
//   lambda
//   nonlocal
//   not
//   or
//   pass
//   raise
//   return
//   try
//   while
//   with
//   yield
//
// Supported import-related module attributes (const StyleBuildInVar):
//   __name__
//   __loader__
//   __package__
//   __spec__
//   __path__
//   __file__
//   __cached__
//
// Supported built-in functions (const StyleBuildInFunc):
//   abs()
//   aiter()
//   all()
//   any()
//   anext()
//   ascii()
//   bin()
//   bool()
//   breakpoint()
//   bytearray()
//   bytes()
//   callable()
//   chr()
//   classmethod()
//   compile()
//   complex()
//   delattr()
//   dict()
//   dir()
//   divmod()
//   enumerate()
//   eval()
//   exec()
//   filter()
//   float()
//   format()
//   frozenset()
//   getattr()
//   globals()
//   hasattr()
//   hash()
//   help()
//   hex()
//   id()
//   input()
//   int()
//   isinstance()
//   issubclass()
//   iter()
//   len()
//   list()
//   locals()
//   map()
//   max()
//   memoryview()
//   min()
//   next()
//   object()
//   oct()
//   open()
//   ord()
//   pow()
//   print()
//   property()
//   range()
//   repr()
//   reversed()
//   round()
//   set()
//   setattr()
//   slice()
//   sorted()
//   staticmethod()
//   str()
//   sum()
//   super()
//   tuple()
//   type()
//   vars()
//   zip()
//   __import__()
//
// Single-line comments (#) and brackets(", ') are also supported.
func Python(code string) string {
	var result string = ""

	// Shild HTML
	code = shieldHTML(code)

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	// Keywords list
	keywords := []string{
		"False", "None", "True",
		"and", "as", "assert",
		"async", "await", "break",
		"class", "continue", "def",
		"del", "elif", "else",
		"except", "finally", "for",
		"from", "global", "if",
		"import", "in", "is",
		"lambda", "nonlocal", "not",
		"or", "pass", "raise",
		"return", "try", "while",
		"with", "yield",
	}

	cmdStartChars := []string{
		"", " ", "\t", ":", ";", "(", ")",
	}

	cmdEndChars := []string{
		"", " ", "\t", ":", ";", "(", ")",
	}

	// Build-in vars
	buildInVars := []string{
		"__name__", "__loader__",
		"__package__", "__spec__",
		"__path__", "__file__",
		"__cached__",
	}

	varStartChars := []string{
		"", " ", "	", ":", "(", ")",
	}

	varEndChars := []string{
		"", " ", "	", ":", "(", ")",
	}

	// Build-in functions
	buildInFuncs := []string{
		"abs", "aiter", "all",
		"any", "anext", "ascii",
		"bin", "bool", "breakpoint",
		"bytearray", "bytes", "callable",
		"chr", "classmethod", "compile",
		"complex", "delattr", "dict",
		"dir", "divmod", "enumerate",
		"eval", "exec", "filter",
		"float", "format", "frozenset",
		"getattr", "globals", "hasattr",
		"hash", "help", "hex",
		"id", "input", "int",
		"isinstance", "issubclass", "iter",
		"len", "list", "locals",
		"map", "max", "memoryview",
		"min", "next", "object",
		"oct", "open", "ord",
		"pow", "print", "property",
		"range", "repr", "reversed",
		"round", "set", "setattr",
		"slice", "sorted", "staticmethod",
		"str", "sum", "super",
		"tuple", "type", "vars",
		"zip", "__import__",
	}

	funcStartChars := []string{
		"", " ", "\t", ":", ";", "(", ")",
	}

	funcEndChars := []string{
		"", " ", "\t", ":", ";", "(", ")",
	}

	// Run parser
	for i := range lines {
		line := lines[i]

		// Comment
		line = formatSharpComment(line)

		// Brackets
		line = formatBrackets(line)

		// Keywords
		for _, word := range keywords {
			line = formatWord(line, word, cmdStartChars, cmdEndChars, StyleKeyword)
		}

		// Build-in vars
		for _, word := range buildInVars {
			line = formatWord(line, word, varStartChars, varEndChars, StyleBuildInVar)
		}

		// Build-in functions
		for _, word := range buildInFuncs {
			line = formatWord(line, word, funcStartChars, funcEndChars, StyleBuildInFunc)
		}

		// Save
		if linesNum != i+1 {
			result = result + line + "\n"
		} else {
			result = result + line
		}
	}

	return result
}
