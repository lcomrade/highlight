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
// Supported operators (const StyleOperator):
//   +
//   -
//   *
//   **
//   /
//   //
//   %
//   @
//   <<
//   >>
//   &
//   |
//   ^
//   ~
//   :=
//   <
//   >
//   <=
//   >=
//   ==
//   !=
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
// Also supported:
//   Single-line comments (#)
//   Multi-line comments (''' ''')
//   Single-line brackets (", ')
//   Multi-line brackets (""" """)
//   Numbers (100, 1.2, 1.25)
func Python(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
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

	// Operators
	operators := []string{
		"+", "-", "*", "**",
		"%", "@", "&lt", "&gt&gt",
		"&amp", "|", "^", "~",
		"=", ":=", "&lt", "&gt",
		"&lt=", "&gt=", "==", "!=",
	}

	// Build-in vars
	buildInVars := []string{
		"__name__", "__loader__",
		"__package__", "__spec__",
		"__path__", "__file__",
		"__cached__",
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

	// Multi-line comments
	code = formatOpenClose(code, `'''`, `'''`, StyleComment)

	// Multi-line brackets
	code = formatOpenClose(code, `"""`, `"""`, StyleBrackets)

	// Sinle-line comments
	code = formatOpenClose(code, "#", "\n", StyleComment)

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleKeyword)
	}

	// Operators
	for _, word := range operators {
		code = formatWord(code, word, defaultOperatorChars, defaultOperatorChars, StyleOperator)
	}

	// Build-in vars
	for _, word := range buildInVars {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleBuildInVar)
	}

	// Build-in functions
	for _, word := range buildInFuncs {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleBuildInFunc)
	}

	// Numbers
	code = formatNumber(code, defaultNumberChars, defaultNumberChars)

	return code
}
