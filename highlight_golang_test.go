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

/*
Multi-line comment:
line 1
line 2
end
*/

func main() {
	// comment
	var answer string
	fmt.Println("Linux or FreeBSD? ")
	fmt.Scanln(&answer)

	if answer == "Linux" {
		fmt.Println("The GPL is a good thing")
		
	} else if answer == "FreeBSD" {
		fmt.Println("The ZFS file system is great for servers")
		
	} else {
		fmt.Println("Unknown answer:(")
	}
}
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>package</span> main

<span class='` + StyleKeyword + `'>import</span>(
	<span class='` + StyleBrackets + `'>"fmt"</span> <span class='` + StyleComment + `'>// this is import</span>
)

<span class='` + StyleComment + `'>/*
Multi-line comment:
line 1
line 2
end
*/</span>

<span class='` + StyleKeyword + `'>func</span> main() {
	<span class='` + StyleComment + `'>// comment</span>
	<span class='` + StyleKeyword + `'>var</span> answer <span class='` + StyleVarType + `'>string</span>
	fmt.Println(<span class='` + StyleBrackets + `'>"Linux or FreeBSD? "</span>)
	fmt.Scanln(<span class='` + StyleOperator + `'>&amp</span>answer)

	<span class='` + StyleKeyword + `'>if</span> answer <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"Linux"</span> {
		fmt.Println(<span class='` + StyleBrackets + `'>"The GPL is a good thing"</span>)
		
	} <span class='` + StyleKeyword + `'>else</span> <span class='` + StyleKeyword + `'>if</span> answer <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"FreeBSD"</span> {
		fmt.Println(<span class='` + StyleBrackets + `'>"The ZFS file system is great for servers"</span>)
		
	} <span class='` + StyleKeyword + `'>else</span> {
		fmt.Println(<span class='` + StyleBrackets + `'>"Unknown answer:("</span>)
	}
}
`,
		},
	}

	runTest(Golang, testData, t)
}
