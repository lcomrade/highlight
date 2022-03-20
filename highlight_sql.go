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

// SQL processes Structured Query Language.
// Read more: https://en.wikipedia.org/wiki/SQL
//
// Supported keywords (const StyleKeyword):
//   ALTER
//   AND
//   AS
//   BEGIN
//   COMMIT
//   CONNECT BY
//   CREATE
//   DELETE
//   DENY
//   DROP
//   FROM
//   GRANT
//   GROUP BY
//   HAVING
//   IN
//   INSERT
//   INTO
//   JOIN
//   OR
//   ORDER BY
//   RELEASE
//   REVOKE
//   ROLLBACK
//   ROLLBACK TO
//   SAVEPOINT
//   SELECT
//   SET
//   TRANSACTION
//   UPDATE
//   VALUES
//   WHERE
//   WORK
//
// Brackets (", ') are also supported.
func SQL(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
	keywords := []string{
		"ALTER", "AND", "AS",
		"BEGIN", "COMMIT", "CONNECT BY",
		"CREATE", "DELETE", "DENY",
		"DROP", "FROM", "GRANT",
		"GROUP BY", "HAVING",
		"IN", "INSERT", "INTO",
		"JOIN", "OR", "ORDER BY",
		"RELEASE", "REVOKE", "ROLLBACK",
		"ROLLBACK TO", "SAVEPOINT", "SELECT",
		"SET", "TRANSACTION", "UPDATE",
		"VALUES", "WHERE", "WORK",
	}

	goodChars := []string{
		" ", "(", ")", ";",
		"\n", "\t",
	}

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, goodChars, goodChars, StyleKeyword)
	}

	return code
}
