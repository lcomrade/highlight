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
	"testing"
)

func TestDockerfile(t *testing.T) {
	testData := []testDataType{
		{
			Input: `

test


`,
			ExpectResult: `

test


`,
		},
		{
			Input: `
FROMalpine
`,
			ExpectResult: `
FROMalpine
`,
		},
		{
			Input: `
# RUN rf -rf /
`,
			ExpectResult: `
<span class='code-c'># RUN rf -rf /</span>
`,
		},
		{
			Input: `
FROM python:3-alpine

# comment 1
# comment 2

WORKDIR /app # comment

COPY ./ ./

VOLUME /data

CMD [ "python", "./main.py" ]
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>FROM</span> python:3-alpine

<span class='` + StyleComment + `'># comment 1</span>
<span class='` + StyleComment + `'># comment 2</span>

<span class='` + StyleKeyword + `'>WORKDIR</span> /app <span class='` + StyleComment + `'># comment</span>

<span class='` + StyleKeyword + `'>COPY</span> ./ ./

<span class='` + StyleKeyword + `'>VOLUME</span> /data

<span class='` + StyleKeyword + `'>CMD</span> [ <span class='` + StyleBrackets + `'>"python"</span>, <span class='` + StyleBrackets + `'>"./main.py"</span> ]
`,
		},
	}

	runTest(Dockerfile, testData, t)
}
