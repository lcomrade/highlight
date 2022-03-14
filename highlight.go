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
	"errors"
	"strings"
)

// These constants describe the name of the class
// that will be assigned to the HTML tag '<span>'.
//
// Examples:
//   <span style='code-k'>User-Agent:</span>
//   <span style='code-c'># My comment</span>
const (
	StyleKeyword = "code-k"
	StyleArg     = "code-a"
	StyleVar     = "code-v"
	StyleComment = "code-c"
)

// ByName helps to highlight code based on the language name.
// This can be useful for Markdown and some other cases.
// The name of the language is not case sensitive.
//
//   | Function name | Language name |
//   |---------------|---------------|
//   | Dockerfile    | dockerfile    |
//   | RobotsTxt     | robots.txt    |
func ByName(code string, language string) (string, error) {
	language = strings.ToLower(language)

	switch strings.ToLower(language) {
	case "dockerfile":
		return Dockerfile(code), nil

	case "robots.txt":
		return RobotsTxt(code), nil
	}

	return "", errors.New("highlight: unknown language name: " + language)
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

// Replace line prefix if it exist.
func replacePrefix(line string, prefixOld string, prefixNew string) string {
	if strings.HasPrefix(line, prefixOld) {
		return prefixNew + line[len(prefixOld):]
	}

	return line
}

// Processes command.
// Content inside <span>....</span> is ignored.
// It is assumed that all HTML tags except <span>....</span> were shielded.
func formatCommand(line string, command string, cmdStartChars []string, cmdEndChars []string) string {
	result := ""

	lineRune := []rune(line)
	lineLen := len(lineRune)
	commandLen := len(command)

	otherSpanTagOpen := 0
	skip := 0

	for i := range lineRune {
		char := string(lineRune[i])

		// Skip
		if skip != 0 {
			skip = skip - 1
			continue
		}

		// Get last char
		lastChar := ""

		if i != 0 {
			lastChar = string(lineRune[i-1])
		}

		// Get command end char
		cmdEndChar := ""

		if lineLen > i+commandLen {
			cmdEndChar = string(lineRune[i+commandLen])
		}

		// Get sub string
		subLine := ""

		if lineLen > i+commandLen {
			subLine = line[i : i+commandLen]
		}

		// Find <span
		if lineLen > i+5 {
			if line[i:i+5] == "<span" {
				otherSpanTagOpen = otherSpanTagOpen + 1
			}
		}

		// Find </span>
		if lineLen > i+7 {
			if line[i:i+7] == "</span>" {
				otherSpanTagOpen = otherSpanTagOpen - 1
			}
		}

		// Check sub string
		if subLine == command && otherSpanTagOpen == 0 {
			okStart := false
			okEnd := false

			// Check command start char
			for _, ch := range cmdStartChars {
				if ch == lastChar {
					okStart = true
					break
				}
			}

			// Check command end char
			for _, ch := range cmdEndChars {
				if ch == cmdEndChar {
					okEnd = true
					break
				}
			}

			// Save
			if okStart == true && okEnd == true {
				result = result + "<span class='" + StyleKeyword + "'>" + command + "</span>"
				skip = commandLen - 1

			} else {
				result = result + char
			}

		} else {
			result = result + char
		}
	}

	return result
}

// Processes '#' (sharp) comments.
func formatSharpComment(line string) string {
	lineRune := []rune(line)
	line = ""
	commentFound := false

	for _, charRune := range lineRune {
		char := string(charRune)

		if commentFound == false && char == "#" {
			line = line + "<span class='" + StyleComment + "'>#"
			commentFound = true
			continue
		}

		line = line + char
	}

	if commentFound == true {
		line = line + "</span>"
	}

	return line
}
