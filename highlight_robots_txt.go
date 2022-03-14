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

// RobotsTxt processes robots.txt files (Robots exclusion standard).
// Read more: https://en.wikipedia.org/wiki/Robots_exclusion_standard
//
// Supports standard directives:
//   User-agent
//   Disallow
//
// And nonstandard extensions:
//   Allow
//   Crawl-delay
//   Sitemap
//   Host
//
// Comments (#) are also supported.
func RobotsTxt(code string) string {
	var result string = ""

	// Shild HTML
	code = shieldHTML(code)

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	for i := range lines {
		line := lines[i]

		// Comment
		line = formatSharpComment(line)

		// Standard
		line = replacePrefix(line, "User-agent:", "<span class='"+StyleKeyword+"'>User-agent:</span>")
		line = replacePrefix(line, "Disallow:", "<span class='"+StyleKeyword+"'>Disallow:</span>")

		// Nonstandard extensions
		line = replacePrefix(line, "Allow:", "<span class='"+StyleKeyword+"'>Allow:</span>")
		line = replacePrefix(line, "Crawl-delay:", "<span class='"+StyleKeyword+"'>Crawl-delay:</span>")
		line = replacePrefix(line, "Sitemap:", "<span class='"+StyleKeyword+"'>Sitemap:</span>")
		line = replacePrefix(line, "Host:", "<span class='"+StyleKeyword+"'>Host:</span>")

		if linesNum != i+1 {
			result = result + line + "\n"
		} else {
			result = result + line
		}
	}

	return result
}
