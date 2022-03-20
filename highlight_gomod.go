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

// GoMod processes go.mod files.
// Read more: https://go.dev/ref/mod#go-mod-file-module
//
// Supported keywords (const StyleKeyword):
//   module
//   go
//   require
//   exclude
//   replace
//   retract
//
// Supported operators (const StyleOperator):
//   =>
//
// Single-line comments (//) are also supported.
func GoMod(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
	keywords := []string{
		"module", "go",
		"require", "exclude",
		"replace", "retract",
	}

	startChars := []string{
		"\n",
	}

	endChars := []string{
		" ", "(",
	}

	// Operators
	operators := []string{
		"=&gt",
	}

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, startChars, endChars, StyleKeyword)
	}

	// Operators
	for _, word := range operators {
		code = formatWord(code, word, []string{" "}, []string{" "}, StyleOperator)
	}

	return code
}
