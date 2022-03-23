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

// Processes structures that have an opening and closing tag.
// Content inside <span>....</span> is ignored.
// It is assumed that all HTML tags except <span>....</span> were shielded.
func formatOpenClose(text string, openStr string, closeStr string, styleClass string) string {
	result := ""

	textRune := []rune(text)
	textLen := len(textRune)
	openLen := len([]rune(openStr))
	closeLen := len([]rune(closeStr))

	spanTagOpen := false
	otherSpanTagOpen := 0
	skip := 0

	for i := range textRune {
		char := string(textRune[i])

		// Skip
		if skip != 0 {
			skip = skip - 1
			continue
		}

		// Get open sub string
		openSub := ""

		if textLen > i+openLen-1 {
			openSub = string(textRune[i : i+openLen])
		}

		// Get close sub string
		closeSub := ""

		if textLen > i+closeLen-1 {
			closeSub = string(textRune[i : i+closeLen])
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

		// Check open str
		if openSub == openStr && spanTagOpen == false && otherSpanTagOpen == 0 {
			result = result + "<span class='" + styleClass + "'>" + openSub
			spanTagOpen = true
			skip = openLen - 1
			continue
		}

		// Check close str
		if closeSub == closeStr && spanTagOpen == true && otherSpanTagOpen == 0 {
			if closeStr == "\n" {
				result = result + "</span>\n"

			} else {
				result = result + closeSub + "</span>"
			}

			spanTagOpen = false
			skip = closeLen - 1
			continue
		}

		// Else
		result = result + char
	}

	if spanTagOpen == true {
		result = result + "</span>"
		spanTagOpen = false
	}

	return result
}
