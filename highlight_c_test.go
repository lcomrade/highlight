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

func TestC(t *testing.T) {
	testData := []testDataType{
		{
			Input: `
#include <stdio.h>

int main() {
	printf("Hello, world!");

	return 0;
}
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>#include</span> &ltstdio.h&gt

<span class='` + StyleKeyword + `'>int</span> main() {
	printf(<span class='` + StyleBrackets + `'>"Hello, world!"</span>);

	<span class='` + StyleKeyword + `'>return</span> <span class='` + StyleValue + `'>0</span>;
}
`,
		},
	}

	runTest(C, testData, t)
}
