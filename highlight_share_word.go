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

var defaultKeywordChars = []string{
	"", " ", "\t", "\n",
	".", ",", "<", ">", ":", ";",
	"=", "+", "-", "*", "/", "&", "^",
	"(", ")", "[", "]", "{", "}",
}

var defaultOperatorChars = []string{
	"", " ", "\t", ":", "(", ")",
	`"`, `'`, "`", "{", "}", "_",
	"1", "2", "3", "4", "5",
	"6", "7", "8", "9", "0",
	"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x",
	"y", "z",
	"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X",
	"Y", "Z",
}

// Processes keywords, commands and var types.
// Content inside <span>....</span> is ignored.
// It is assumed that all HTML tags except <span>....</span> were shielded.
func formatWord(text string, command string, cmdStartChars []string, cmdEndChars []string, styleClass string) string {
	result := ""

	textRune := []rune(text)
	textLen := len(textRune)
	commandLen := len([]rune(command))

	otherSpanTagOpen := 0
	skip := 0

	for i := range textRune {
		char := string(textRune[i])

		// Skip
		if skip != 0 {
			skip = skip - 1
			continue
		}

		// Get last char
		lastChar := ""

		if i != 0 {
			lastChar = string(textRune[i-1])
		}

		// Get command end char
		cmdEndChar := ""

		if textLen > i+commandLen {
			cmdEndChar = string(textRune[i+commandLen])
		}

		// Get sub string
		subLine := ""

		if textLen > i+commandLen-1 {
			subLine = string(textRune[i : i+commandLen])
		}

		// Find <span
		if textLen > i+5 {
			if string(textRune[i:i+5]) == "<span" {
				otherSpanTagOpen = otherSpanTagOpen + 1
			}
		}

		// Find </span>
		if textLen > i+7 {
			if string(textRune[i:i+7]) == "</span>" {
				otherSpanTagOpen = otherSpanTagOpen - 1
			}
		}

		// Check sub string
		if subLine == command && otherSpanTagOpen == 0 {
			// OK
			if isInStrList(cmdStartChars, lastChar) && isInStrList(cmdEndChars, cmdEndChar) {
				result = result + "<span class='" + styleClass + "'>" + command + "</span>"
				skip = commandLen - 1

				// Not OK
			} else {
				result = result + char
			}

		} else {
			result = result + char
		}
	}

	return result
}
