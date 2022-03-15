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
// Supported keywords:
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
// Single-line comments (#) and brackets(", ') are also supported.
func Golang(code string) string {
	var result string = ""

	// Shild HTML
	code = shieldHTML(code)

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	// Keywords list
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

	cmdStartChars := []string{
		"", " ", "	", ":", ";", "(", ")",
	}

	cmdEndChars := []string{
		"", " ", "	", ":", ";", "(", ")",
	}

	// Run parser
	for i := range lines {
		line := lines[i]

		// Comment
		line = formatCComment(line)

		// Brackets
		line = formatBrackets(line)

		// Keywords
		for _, word := range keywords {
			line = formatWord(line, word, cmdStartChars, cmdEndChars, StyleKeyword)
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
