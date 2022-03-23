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

// Golang processes go source code (*.go files).
// Read more: https://go.dev/ref/spec
//
// Supported keywords (const StyleKeyword):
//   break
//   case
//   chan
//   const
//   continue
//   default
//   defer
//   else
//   fallthrough
//   for
//   func
//   go
//   goto
//   if
//   import
//   interface
//   map
//   package
//   range
//   return
//   select
//   struct
//   switch
//   type
//   var
//
// Supported values (const StyleValue):
//   true
//   false
//   nil
//   All numbers (100, 1.3, 2.25)
//
// Supported operators (const StyleOperator):
//   !
//   !=
//   %
//   &
//   &&
//   &^
//   *
//   +
//   -
//   /
//   <
//   <-
//   <<
//   <=
//   ==
//   >
//   >=
//   >>
//   ^
//   |
//   ||
//   =
//   :=
//
// Supported variable types (const StyleVarType):
//   bool
//   uint
//   uint8
//   uint16
//   uint32
//   uint64
//   uintptr
//   int
//   int8
//   int16
//   int32
//   int64
//   float32
//   float64
//   complex64
//   complex128
//   byte
//   rune
//   string
//
// Also supported:
//   Single-line comments (//)
//   Multi-line comments (/* */)
//   Single-line brackets (", ')
//   Multi-line brackets (` `)
func Golang(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
	keywords := []string{
		"break", "case", "chan",
		"const", "continue", "default",
		"defer", "else", "fallthrough",
		"for", "func", "go",
		"goto", "if", "import",
		"interface", "map", "package",
		"range", "return", "select",
		"struct", "switch", "type",
		"var",
	}

	// Values
	values := []string{
		"true", "false", "nil",
	}

	// Operators
	operators := []string{
		"!", "!=", "%", "&amp",
		"&amp&amp", "&amp^", "*", "+",
		"-", "/", "&lt", "&lt-",
		"&lt&lt", "&lt=", "==", "&gt",
		"&gt=", "&gt&gt", "^", "|",
		"||", "=", ":=",
	}

	// Varibles types
	varTypes := []string{
		"bool", "uint", "uint8",
		"uint16", "uint32", "uint64",
		"uintptr", "int", "int8",
		"int16", "int32", "int64",
		"float32", "float64", "complex64",
		"complex128", "byte", "rune",
		"string",
	}

	// Multi-line comments
	code = formatOpenClose(code, "/*", "*/", StyleComment)

	// Multi-line brackets
	code = formatOpenClose(code, "`", "`", StyleBrackets)

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

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

	// Varibles types
	for _, word := range varTypes {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleVarType)
	}

	// Numbers
	code = formatNumber(code, defaultNumberChars, defaultNumberChars)

	// true, false and nil
	for _, word := range values {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleValue)
	}

	return code
}
