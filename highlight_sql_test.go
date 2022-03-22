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

func TestSQL(t *testing.T) {
	testData := []testDataType{
		{
			Input: `
SELECT
  DeptID,
  SUM(SaleAmount)
FROM
  Sales
WHERE
  SaleDate = '01-Jan-2000'
GROUP BY
  DeptID
HAVING
  SUM(SaleAmount) > 1000
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>SELECT</span>
  DeptID,
  SUM(SaleAmount)
<span class='` + StyleKeyword + `'>FROM</span>
  Sales
<span class='` + StyleKeyword + `'>WHERE</span>
  SaleDate = <span class='` + StyleBrackets + `'>'01-Jan-2000'</span>
<span class='` + StyleKeyword + `'>GROUP BY</span>
  DeptID
<span class='` + StyleKeyword + `'>HAVING</span>
  SUM(SaleAmount) &gt <span class='` + StyleNumber + `'>1000</span>
`,
		},
	}

	runTest(SQL, testData, t)
}
