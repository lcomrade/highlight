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

// JSON processes JSON config files (*.json files).
// Read more: https://en.wikipedia.org/wiki/JSON
//
// Supported values (const StyleValue):
//   true
//   false
//   null
//   All numbers (100, 1.3, 2.25)
//
// Also supported:
//   Single-line comments (//)
//   Multi-line comments (/* */)
//   Single-line brackets (", ')
func JSON(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Values
	values := []string{
		"true", "false", "null",
	}

	// Multi-line comments
	code = formatOpenClose(code, "/*", "*/", StyleComment)

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Numbers
	code = formatNumber(code, defaultNumberChars, defaultNumberChars)

	// true, false and null
	for _, word := range values {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleValue)
	}

	return code
}
