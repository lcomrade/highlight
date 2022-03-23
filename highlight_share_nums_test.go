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

func TestShareFormatNumber(t *testing.T) {
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
			Input:        "100",
			ExpectResult: "<span class='" + StyleValue + "'>100</span>",
		},
		{
			Input:        "1.25",
			ExpectResult: "<span class='" + StyleValue + "'>1</span>.<span class='" + StyleValue + "'>25</span>",
		},
		{
			Input:        ".255",
			ExpectResult: ".<span class='" + StyleValue + "'>255</span>",
		},
		{
			Input:        "\n100\n",
			ExpectResult: "\n<span class='" + StyleValue + "'>100</span>\n",
		},
		{
			Input:        "a=100+b",
			ExpectResult: "a=<span class='" + StyleValue + "'>100</span>+b",
		},
		{
			Input:        "<span class='" + StyleOperator + "'>+</span>100",
			ExpectResult: "<span class='" + StyleOperator + "'>+</span><span class='" + StyleValue + "'>100</span>",
		},
		{
			Input:        "Русский\nтекст\n",
			ExpectResult: "Русский\nтекст\n",
		},
	}

	testFunc := func(text string) string {
		return formatNumber(text, defaultNumberChars, defaultNumberChars)
	}

	runTest(testFunc, testData, t)
}
