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

func TestRobotsTxt(t *testing.T) {
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
  Allow: /
`,
			ExpectResult: `
  Allow: /
`,
		},
		{
			Input: `
# Allow:
`,
			ExpectResult: `
<span class='code-c'># Allow:</span>
`,
		},
		{
			Input: `
User-agent: * # comment
Disallow: /faq

# comment 1
# comment 2

Allow: /
Crawl-delay: 10
Sitemap: https://example.org/sitemap.xml
Host: https://mirror.example.org
`,
			ExpectResult: `
<span class='` + StyleKeyword + `'>User-agent:</span> * <span class='` + StyleComment + `'># comment</span>
<span class='` + StyleKeyword + `'>Disallow:</span> /faq

<span class='` + StyleComment + `'># comment 1</span>
<span class='` + StyleComment + `'># comment 2</span>

<span class='` + StyleKeyword + `'>Allow:</span> /
<span class='` + StyleKeyword + `'>Crawl-delay:</span> 10
<span class='` + StyleKeyword + `'>Sitemap:</span> https://example.org/sitemap.xml
<span class='` + StyleKeyword + `'>Host:</span> https://mirror.example.org
`,
		},
	}

	runTest(RobotsTxt, testData, t)
}
