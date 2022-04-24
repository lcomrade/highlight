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

func TestINI(t *testing.T) {
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
# My ini file

; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization=Acme Widgets Inc.
`,
			ExpectResult: `
<span class='` + StyleComment + `'># My ini file</span>

<span class='` + StyleComment + `'>; last modified 1 April 2001 by John Doe</span>
<span class='` + StyleBrackets + `'>[owner]</span>
name = John Doe
organization=Acme Widgets Inc.
`,
		},
	}

	runTest(INI, testData, t)
}
