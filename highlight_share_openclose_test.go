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

func TestShareFormatOpenClose(t *testing.T) {
	testData := []testDataType{
		{
			Input:        "",
			ExpectResult: "",
		},
		{
			Input:        "\ntest\n",
			ExpectResult: "\ntest\n",
		},
		{
			Input:        "Русский\nтекст\n",
			ExpectResult: "Русский\nтекст\n",
		},
		{
			Input:        `"my text"`,
			ExpectResult: `<span class='` + StyleBrackets + `'>"my text"</span>`,
		},
		{
			Input:        `111"my text"111`,
			ExpectResult: `111<span class='` + StyleBrackets + `'>"my text"</span>111`,
		},
		{
			Input:        `a:="my text"`,
			ExpectResult: `a:=<span class='` + StyleBrackets + `'>"my text"</span>`,
		},
	}

	testFunc := func(text string) string {
		for _, openClose := range []string{`"`, `'`} {
			text = formatOpenClose(text, openClose, openClose, StyleBrackets)
		}
		return text
	}

	runTest(testFunc, testData, t)
}
