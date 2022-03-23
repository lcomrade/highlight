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

func TestShareFormatWordKeyword(t *testing.T) {
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
			Input:        "COMMAND",
			ExpectResult: "<span class='" + StyleKeyword + "'>COMMAND</span>",
		},
		{
			Input:        "КОМАНДА",
			ExpectResult: "<span class='" + StyleKeyword + "'>КОМАНДА</span>",
		},
		{
			Input:        "123COMMAND123",
			ExpectResult: "123COMMAND123",
		},
		{
			Input:        "       COMMAND",
			ExpectResult: "       <span class='" + StyleKeyword + "'>COMMAND</span>",
		},
		{
			Input:        "myVar:=COMMAND(123)",
			ExpectResult: "myVar:=<span class='" + StyleKeyword + "'>COMMAND</span>(123)",
		},
	}

	testFunc := func(text string) string {
		for _, command := range []string{"COMMAND", "КОМАНДА"} {
			text = formatWord(text, command, defaultKeywordChars, defaultKeywordChars, StyleKeyword)
		}
		return text
	}

	runTest(testFunc, testData, t)
}


func TestShareFormatWordOperator(t *testing.T) {
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
			Input:        "a=2+1",
			ExpectResult: "a<span class='"+StyleOperator+"'>=</span>2<span class='"+StyleOperator+"'>+</span>1",
		},
		{
			Input:        "c=a+b",
			ExpectResult: "c<span class='"+StyleOperator+"'>=</span>a<span class='"+StyleOperator+"'>+</span>b",
		},
	}

	testFunc := func(text string) string {
		for _, operator := range []string{"=", "+", "-", "*", "/"} {
			text = formatWord(text, operator, defaultOperatorChars, defaultOperatorChars, StyleOperator)
		}
		return text
	}

	runTest(testFunc, testData, t)
}
