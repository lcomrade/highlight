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
func RobotsTxt(code string) string {
	var result string = ""

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	for i := range lines {
		line := lines[i]

		// Comment
		lineRune := []rune(line)
		line = ""
		commentFound := false

		for _, charRune := range lineRune {
			char := string(charRune)

			if commentFound == false && char == "#" {
				line = line + "<span class='" + StyleComment + "'>#"
				commentFound = true
				continue
			}

			line = line + char
		}

		if commentFound == true {
			line = line + "</span>"
		}

		// Standard
		line = strings.Replace(line, "User-agent:", "<span class='"+StyleKeyword+"'>User-agent:</span>", 1)
		line = strings.Replace(line, "Disallow:", "<span class='"+StyleKeyword+"'>Disallow:</span>", 1)

		// Nonstandard extensions
		line = strings.Replace(line, "Allow:", "<span class='"+StyleKeyword+"'>Allow:</span>", 1)
		line = strings.Replace(line, "Crawl-delay:", "<span class='"+StyleKeyword+"'>Crawl-delay:</span>", 1)
		line = strings.Replace(line, "Sitemap:", "<span class='"+StyleKeyword+"'>Sitemap:</span>", 1)
		line = strings.Replace(line, "Host:", "<span class='"+StyleKeyword+"'>Host:</span>", 1)

		if linesNum != i+1 {
			result = result + line + "\n"
		} else {
			result = result + line
		}
	}

	return result
}
