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
//   <span style='code-keyword'>User-Agent:</span>
//   <span style='code-comment'># My comment</span>
const (
	StyleKeyword     = "code-keyword"
	StyleOperator    = "code-operator"
	StyleVarType     = "code-var-type"
	StyleBuildInVar  = "code-build-in-var"
	StyleBuildInFunc = "code-build-in-func"
	StyleComment     = "code-comment"
	StyleBrackets    = "code-brackets"
)

// ByName helps to highlight code based on the language name.
// This can be useful for Markdown and some other cases.
// The name of the language is not case sensitive.
//
//   | Function name | Language name   |
//   |---------------|-----------------|
//   | C             | c               |
//   | Dockerfile    | dockerfile      |
//   | Golang        | go, golang      |
//   | Python        | python, python3 |
//   | RobotsTxt     | robots.txt      |
func ByName(code string, language string) (string, error) {
	language = strings.ToLower(language)

	switch strings.ToLower(language) {
	case "c":
		return C(code), nil

	case "dockerfile":
		return Dockerfile(code), nil

	case "go", "golang":
		return Golang(code), nil

	case "python", "python3":
		return Python(code), nil

	case "robots.txt":
		return RobotsTxt(code), nil
	}

	return "", errors.New("highlight: unknown language name: " + language)
}
