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
// Single-line comments (//) and multi-line comments (/* */).
// Brackets(", ') are also supported.
func Golang(code string) string {
	var result string = ""

	// Shild HTML
	code = shieldHTML(code)

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

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

	cmdChars := []string{
		"", " ", "\t", "\n", ":", ";", "(", ")",
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

	opsChars := []string{
		"", " ", "\t", ":", "(", ")",
		`"`, `'`, "`", "_",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "0",
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x",
		"y", "z",
		"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X",
		"Y", "Z",
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

	// Run parser
	commentOpen := false

	for i := range lines {
		line := lines[i]

		// Multi-line comments
		lastCommentOpen := commentOpen
		line, commentOpen = formatCMultiComment(line, commentOpen)

		if lastCommentOpen != true && commentOpen != true {
			// Single-line comment
			line = formatCComment(line)

			// Brackets
			line = formatBrackets(line)

			// Keywords
			for _, word := range keywords {
				line = formatWord(line, word, cmdChars, cmdChars, StyleKeyword)
			}

			// Operators
			for _, word := range operators {
				line = formatWord(line, word, opsChars, opsChars, StyleOperator)
			}

			// Varibles types
			for _, word := range varTypes {
				line = formatWord(line, word, cmdChars, cmdChars, StyleVarType)
			}
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
