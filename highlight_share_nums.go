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

var defaultNumberChars = []string{
	" ", "\t", "\n",
	".", "=", ":", ";",
	"(", ")", "[", "]", "{", "}",
}

// Processes numbers.
// Content inside <span>....</span> is ignored.
// It is assumed that all HTML tags except <span>....</span> were shielded.
func formatNumber(text string, cmdStartChars []string, cmdEndChars []string) string {
	result := ""

	textRune := []rune(text)
	textLen := len(textRune)

	buffer := ""
	bufferOpen := false
	otherSpanTagOpen := 0

	for i := range textRune {
		charRune := textRune[i]
		char := string(charRune)

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

		// Check
		if otherSpanTagOpen == 0 {
			// Close buffer
			if bufferOpen == true && isNumber(charRune) == false {
				// Good close
				if isInStrList(cmdEndChars, char) && buffer != "" {
					result = result + "<span class='" + StyleNumber + "'>" + buffer + "</span>"

					// Bad close
				} else {
					result = result + buffer
				}

				// Save
				result = result + char
				buffer = ""
				bufferOpen = false
				continue
			}

			// Open buffer
			if bufferOpen == false && isInStrList(cmdStartChars, char) {
				result = result + char
				bufferOpen = true
				continue
			}

			// Continue read
			if bufferOpen == true {
				buffer = buffer + char

			} else {
				result = result + char
			}

		} else {
			if bufferOpen == true {
				if buffer != "" {
					result = result + "<span class='" + StyleNumber + "'>" + buffer + "</span>"
					buffer = ""
				}

				bufferOpen = false
			}

			result = result + char
		}
	}

	if bufferOpen == true && buffer != "" {
		result = result + "<span class='" + StyleNumber + "'>" + buffer + "</span>"
	}

	return result
}

func isInStrList(list []string, char string) bool {
	for _, ch := range list {
		if ch == char {
			return true
		}
	}

	return false
}
