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

func TestPython(t *testing.T) {
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
# if __name__ == "__main__":
#	main()
`,
			ExpectResult: `
<span class='` + StyleComment + `'># if __name__ == "__main__":</span>
<span class='` + StyleComment + `'>#	main()</span>
`,
		},
		{
			Input: `
#!/usr/bin/python3
"""
This is test.
"""

def main() -> None:
	name = input("Your name:")

	if name == "Ivan" or name == "Nikolai":
		print("Vodka, matryoshka, balalaika!")
		
	elif name == "Franklin":
		print("Greetings, USA President!")

	else:
		print("I don't know what to say about that name:(")

'''
My multi-line comment:
   line 1
   line 2
   line 3

End
'''

if __name__ == "__main__":
	main()
`,
			ExpectResult: `
<span class='` + StyleComment + `'>#!/usr/bin/python3</span>
<span class='` + StyleBrackets + `'>"""
This is test.
"""</span>

<span class='` + StyleKeyword + `'>def</span> main() -&gt <span class='code-k'>None</span>:
	name <span class='` + StyleOperator + `'>=</span> <span class='` + StyleBuildInFunc + `'>input</span>(<span class='` + StyleBrackets + `'>"Your name:"</span>)

	<span class='` + StyleKeyword + `'>if</span> name <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"Ivan"</span> <span class='` + StyleKeyword + `'>or</span> name <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"Nikolai"</span>:
		<span class='` + StyleBuildInFunc + `'>print</span>(<span class='` + StyleBrackets + `'>"Vodka, matryoshka, balalaika!"</span>)
		
	<span class='` + StyleKeyword + `'>elif</span> name <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"Franklin"</span>:
		<span class='` + StyleBuildInFunc + `'>print</span>(<span class='` + StyleBrackets + `'>"Greetings, USA President!"</span>)

	<span class='` + StyleKeyword + `'>else</span>:
		<span class='` + StyleBuildInFunc + `'>print</span>(<span class='` + StyleBrackets + `'>"I don't know what to say about that name:("</span>)

<span class='` + StyleComment + `'>'''
My multi-line comment:
   line 1
   line 2
   line 3

End
'''</span>

<span class='` + StyleKeyword + `'>if</span> <span class='` + StyleBuildInVar + `'>__name__</span> <span class='` + StyleOperator + `'>==</span> <span class='` + StyleBrackets + `'>"__main__"</span>:
	main()
`,
		},
	}

	runTest(Python, testData, t)
}
