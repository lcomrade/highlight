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

type testDataType struct {
	Input        string
	ExpectResult string
}

func runTest(function func(string) string, testData []testDataType, t *testing.T) {
	for _, test := range testData {
		result := function(test.Input)
		if result != test.ExpectResult {
			t.Error("\n" + "Input:    '" + test.Input + "'\n" + "Expected: '" + test.ExpectResult + "'\n" + "But get:  '" + result + "'")
		}
	}
}
