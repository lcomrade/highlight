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

func TestJSON(t *testing.T) {
	testData := []testDataType{
		{
			Input: `

test


`,
			ExpectResult: `

test


`,
		},
		{
			Input: `
{
	{
	  "firstName": "John",
	  "lastName": "Smith",
	  "isAlive": true,
	  "age": 27,
	  "address": {
	    "streetAddress": "21 2nd Street",
	    "city": "New York",
	    "state": "NY",
	    "postalCode": "10021-3100"
	  },
	  "phoneNumbers": [
	    {
	      "type": "home",
	      "number": "212 555-1234"
	    },
	    {
	      "type": "office",
	      "number": "646 555-4567"
	    }
	  ],
	  "children": [],
	  "spouse": null
	}
}
`,
			ExpectResult: `
{
	{
	  <span class='` + StyleBrackets + `'>"firstName"</span>: <span class='` + StyleBrackets + `'>"John"</span>,
	  <span class='` + StyleBrackets + `'>"lastName"</span>: <span class='` + StyleBrackets + `'>"Smith"</span>,
	  <span class='` + StyleBrackets + `'>"isAlive"</span>: <span class='` + StyleValue + `'>true</span>,
	  <span class='` + StyleBrackets + `'>"age"</span>: <span class='` + StyleValue + `'>27</span>,
	  <span class='` + StyleBrackets + `'>"address"</span>: {
	    <span class='` + StyleBrackets + `'>"streetAddress"</span>: <span class='` + StyleBrackets + `'>"21 2nd Street"</span>,
	    <span class='` + StyleBrackets + `'>"city"</span>: <span class='` + StyleBrackets + `'>"New York"</span>,
	    <span class='` + StyleBrackets + `'>"state"</span>: <span class='` + StyleBrackets + `'>"NY"</span>,
	    <span class='` + StyleBrackets + `'>"postalCode"</span>: <span class='` + StyleBrackets + `'>"10021-3100"</span>
	  },
	  <span class='` + StyleBrackets + `'>"phoneNumbers"</span>: [
	    {
	      <span class='` + StyleBrackets + `'>"type"</span>: <span class='` + StyleBrackets + `'>"home"</span>,
	      <span class='` + StyleBrackets + `'>"number"</span>: <span class='` + StyleBrackets + `'>"212 555-1234"</span>
	    },
	    {
	      <span class='` + StyleBrackets + `'>"type"</span>: <span class='` + StyleBrackets + `'>"office"</span>,
	      <span class='` + StyleBrackets + `'>"number"</span>: <span class='` + StyleBrackets + `'>"646 555-4567"</span>
	    }
	  ],
	  <span class='` + StyleBrackets + `'>"children"</span>: [],
	  <span class='` + StyleBrackets + `'>"spouse"</span>: <span class='` + StyleValue + `'>null</span>
	}
}
`,
		},
	}

	runTest(JSON, testData, t)
}
