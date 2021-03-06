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

func TestGoMod(t *testing.T) {
	testData := []testDataType{
		{
			Input: `
module example.com/mymodule

go 1.14

require (
    example.com/othermodule v1.2.3
    example.com/thismodule v1.2.3
    example.com/thatmodule v1.2.3
)

replace example.com/thatmodule => ../thatmodule
exclude example.com/thismodule v1.3.0
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>module</span> example.com/mymodule

<span class='` + StyleKeyword + `'>go</span> <span class='` + StyleValue + `'>1</span>.<span class='` + StyleValue + `'>14</span>

<span class='` + StyleKeyword + `'>require</span> (
    example.com/othermodule v<span class='` + StyleValue + `'>1</span>.<span class='` + StyleValue + `'>2</span>.<span class='` + StyleValue + `'>3</span>
    example.com/thismodule v<span class='` + StyleValue + `'>1</span>.<span class='` + StyleValue + `'>2</span>.<span class='` + StyleValue + `'>3</span>
    example.com/thatmodule v<span class='` + StyleValue + `'>1</span>.<span class='` + StyleValue + `'>2</span>.<span class='` + StyleValue + `'>3</span>
)

<span class='` + StyleKeyword + `'>replace</span> example.com/thatmodule <span class='` + StyleOperator + `'>=&gt</span> ../thatmodule
<span class='` + StyleKeyword + `'>exclude</span> example.com/thismodule v<span class='` + StyleValue + `'>1</span>.<span class='` + StyleValue + `'>3</span>.<span class='` + StyleValue + `'>0</span>
`,
		},
	}

	runTest(GoMod, testData, t)
}
