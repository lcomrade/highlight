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

	lines := strings.Split(code, "\n")
	linesNum := len(lines)

	for i := range lines {
		line := lines[i]
		// Comments
		line = formatSharpComment(line)

		// Commands
		line = strings.Replace(line, "FROM ", "<span class='"+StyleKeyword+"'>FROM</span> ", 1)
		line = strings.Replace(line, "RUN ", "<span class='"+StyleKeyword+"'>RUN</span> ", 1)
		line = strings.Replace(line, "CMD ", "<span class='"+StyleKeyword+"'>CMD</span> ", 1)
		line = strings.Replace(line, "LABEL ", "<span class='"+StyleKeyword+"'>LABEL</span> ", 1)
		line = strings.Replace(line, "MAINTAINER ", "<span class='"+StyleKeyword+"'>MAINTAINER</span> ", 1)
		line = strings.Replace(line, "EXPOSE ", "<span class='"+StyleKeyword+"'>EXPOSE</span> ", 1)
		line = strings.Replace(line, "ADD ", "<span class='"+StyleKeyword+"'>ADD</span> ", 1)
		line = strings.Replace(line, "COPY ", "<span class='"+StyleKeyword+"'>COPY</span> ", 1)
		line = strings.Replace(line, "ENTRYPOINT ", "<span class='"+StyleKeyword+"'>ENTRYPOINT</span> ", 1)
		line = strings.Replace(line, "VOLUME ", "<span class='"+StyleKeyword+"'>VOLUME</span> ", 1)
		line = strings.Replace(line, "USER ", "<span class='"+StyleKeyword+"'>USER</span> ", 1)
		line = strings.Replace(line, "FROM ", "<span class='"+StyleKeyword+"'>FROM</span> ", 1)
		line = strings.Replace(line, "WORKDIR ", "<span class='"+StyleKeyword+"'>WORKDIR</span> ", 1)
		line = strings.Replace(line, "ARG ", "<span class='"+StyleKeyword+"'>ARG</span> ", 1)
		line = strings.Replace(line, "ONBUILD ", "<span class='"+StyleKeyword+"'>ONBUILD</span> ", 1)
		line = strings.Replace(line, "STOPSIGNAL ", "<span class='"+StyleKeyword+"'>STOPSIGNAL</span> ", 1)
		line = strings.Replace(line, "HEALTHCHECK ", "<span class='"+StyleKeyword+"'>HEALTHCHECK</span> ", 1)
		line = strings.Replace(line, "SHELL ", "<span class='"+StyleKeyword+"'>SHELL</span> ", 1)

		if linesNum != i+1 {
			result = result + line + "\n"
		} else {
			result = result + line
		}
	}

	return result
}
