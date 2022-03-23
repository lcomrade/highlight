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
	"strings"
)

// Replace line prefix if it exist.
func replacePrefix(line string, prefixOld string, prefixNew string) string {
	if strings.HasPrefix(line, prefixOld) {
		return prefixNew + line[len(prefixOld):]
	}

	return line
}

// Checks if rune is a number.
func isNumber(char rune) bool {
	switch char {
	case '0':
		return true
	case '1':
		return true
	case '2':
		return true
	case '3':
		return true
	case '4':
		return true
	case '5':
		return true
	case '6':
		return true
	case '7':
		return true
	case '8':
		return true
	case '9':
		return true
	}

	return false
}

// Checks if the string is in the list.
func isInStrList(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}

	return false
}

// Shields HTML tags.
func shieldHTML(code string) string {
	var result string = ""

	codeRune := []rune(code)

	for i := range codeRune {
		char := string(codeRune[i])

		switch char {
		case "<":
			result = result + "&lt"
		case ">":
			result = result + "&gt"
		case "&":
			result = result + "&amp"
		default:
			result = result + char
		}
	}

	return result
}
