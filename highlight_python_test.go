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
<span class='code-c'># if __name__ == "__main__":</span>
<span class='code-c'>#	main()</span>
`,
		},
		{
			Input: `
#!/usr/bin/python3

def main() -> None:
	name = input("Your name:")

	if name == "Ivan" or name == "Nikolai":
		print("Vodka, matryoshka, balalaika!")
		
	elif name == "Franklin":
		print("Greetings, USA President!")

	else:
		print("I don't know what to say about that name:(")

if __name__ == "__main__":
	main()
`,
			ExpectResult: `
<span class='code-c'>#!/usr/bin/python3</span>

<span class='code-k'>def</span> main() -&gt <span class='code-k'>None</span>:
	name = <span class='code-bf'>input</span>("Your name:")

	<span class='code-k'>if</span> name == "Ivan" <span class='code-k'>or</span> name == "Nikolai":
		<span class='code-bf'>print</span>("Vodka, matryoshka, balalaika!")
		
	<span class='code-k'>elif</span> name == "Franklin":
		<span class='code-bf'>print</span>("Greetings, USA President!")

	<span class='code-k'>else</span>:
		<span class='code-bf'>print</span>("I don't know what to say about that name:(")

<span class='code-k'>if</span> <span class='code-bv'>__name__</span> == "__main__":
	main()
`,
		},
	}

	runTest(Python, testData, t)
}
