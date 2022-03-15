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

// Dockerfile processes Docker builder files.
// Read more: https://docs.docker.com/engine/reference/builder/
//
// Supported commands:
//   FROM
//   RUN
//   CMD
//   LABEL
//   MAINTAINER
//   EXPOSE
//   ENV
//   ADD
//   COPY
//   ENTRYPOINT
//   VOLUME
//   USER
//   WORKDIR
//   ARG
//   ONBUILD
//   STOPSIGNAL
//   HEALTHCHECK
//   SHELL
//
// Comments (#) are also supported.
func Dockerfile(code string) string {
	var result string = ""

	// Shild HTML
	code = shieldHTML(code)

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	// Commands list
	commands := []string{
		"FROM", "RUN", "CMD",
		"LABEL", "MAINTAINER", "EXPOSE",
		"ENV", "ADD", "COPY",
		"ENTRYPOINT", "VOLUME", "USER",
		"WORKDIR", "ARG", "ONBUILD",
		"STOPSIGNAL", "HEALTHCHECK", "SHELL",
	}

	for i := range lines {
		line := lines[i]
		// Comment
		line = formatSharpComment(line)

		// Brackets
		line = formatBrackets(line)

		// Commands
		for _, cmd := range commands {
			line = formatWord(line, cmd, []string{"", " ", "	"}, []string{"", " ", "	"}, StyleKeyword)
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
