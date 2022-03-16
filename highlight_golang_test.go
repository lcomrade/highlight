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

func TestGolang(t *testing.T) {
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
package main

import(
	"fmt" // this is import
)

func main() {
	// comment
	var string hi = "Hellow World!"
	fmt.Println(hi)
}
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>package</span> main

<span class='` + StyleKeyword + `'>import</span>(
	<span class='` + StyleBrackets + `'>"fmt"</span> <span class='` + StyleComment + `'>// this is import</span>
)

<span class='` + StyleKeyword + `'>func</span> main() {
	<span class='` + StyleComment + `'>// comment</span>
	<span class='` + StyleKeyword + `'>var</span> <span class='` + StyleVarType + `'>string</span> hi = <span class='` + StyleBrackets + `'>"Hellow World!"</span>
	fmt.Println(hi)
}
`,
		},
	}

	runTest(Golang, testData, t)
}
