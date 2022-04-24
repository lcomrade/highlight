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

// INI processes .INI configuration files.
// Read more: https://en.wikipedia.org/wiki/INI_file
//
// Supported:
//   Sections ([....])
//   Single-line comments (# and ;)
func INI(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Single-line comments
	code = formatOpenClose(code, "#", "\n", StyleComment)
	code = formatOpenClose(code, ";", "\n", StyleComment)

	// Sections
	code = formatOpenClose(code, "[", "]", StyleBrackets)

	return code
}
