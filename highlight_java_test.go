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

func TestJava(t *testing.T) {
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
/*
---HELLO WORLD APP---
*/

public class HelloWorldApp {
    public static void main(String[] args) {
        System.out.println("Hello World!"); // Prints the string to the console.
    }
}
`,
			ExpectResult: `
<span class='` + StyleComment + `'>/*
---HELLO WORLD APP---
*/</span>

<span class='` + StyleKeyword + `'>public</span> <span class='` + StyleKeyword + `'>class</span> HelloWorldApp {
    <span class='` + StyleKeyword + `'>public</span> <span class='` + StyleKeyword + `'>static</span> <span class='` + StyleKeyword + `'>void</span> main(<span class='` + StyleVarType + `'>String</span>[] args) {
        System.out.println(<span class='` + StyleBrackets + `'>"Hello World!"</span>); <span class='` + StyleComment + `'>// Prints the string to the console.</span>
    }
}
`,
		},
	}

	runTest(Java, testData, t)
}
