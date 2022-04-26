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

// Java processes Java source code (*.java files).
// Read more: https://en.wikipedia.org/wiki/Java_(programming_language)
//
// Supported keywords (const StyleKeyword):
//   abstract
//   assert
//   break
//   case
//   catch
//   class
//   const
//   continue
//   default
//   do
//   else
//   enum
//   extends
//   final
//   finally
//   for
//   goto
//   if
//   implements
//   import
//   instanceof
//   interface
//   native
//   new
//   package
//   private
//   protected
//   public
//   return
//   static
//   strictfp
//   super
//   switch
//   synchronized
//   this
//   throw
//   throws
//   transient
//   try
//   void
//   volatile
//   while
//
// Supported values (const StyleValue):
//   true
//   false
//   null
//   All numbers (100, 1.3, 2.25)
//
// Supported operators (const StyleOperator):
//   !
//   !=
//   %
//   %=
//   &&
//   &=
//   */
//   *=
//   +
//   ++
//   +=
//   -
//   --
//   -=
//   /=
//   <
//   <<=
//   <=
//   =
//   ==
//   >
//   >=
//   >>=
//   ^=
//   |=
//   ||
//
// Supported variable types (const StyleVarType):
//   boolean
//   byte
//   char
//   double
//   float
//   int
//   long
//   short
//   String
//
// Also supported:
//   Single-line comments (//)
//   Multi-line comments (/* */)
//   Single-line brackets (", ')
//   Multi-line brackets (""" """)
func Java(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
	keywords := []string{
		"abstract", "assert", "break",
		"case", "catch", "class",
		"const", "continue", "default",
		"do", "else", "enum",
		"extends", "final", "finally",
		"for", "goto", "if",
		"implements", "import", "instanceof",
		"interface", "native", "new",
		"package", "private", "protected",
		"public", "return", "static",
		"strictfp", "super", "switch",
		"synchronized", "this", "throw",
		"throws", "transient", "try",
		"void", "volatile", "while",
	}

	// Values
	values := []string{
		"true", "false", "null",
	}

	// Operators
	operators := []string{
		"!", "!=", "%", "%=",
		"&amp&amp", "&amp=", "*/",
		"*=", "+", "++", "+=",
		"-", "--", "-=", "/=",
		"&lt", "&lt&lt=", "&lt=",
		"=", "==", "&gt", "&gt=",
		"&gt&gt=", "^=", "|=", "||",
	}

	// Varibles types
	varTypes := []string{
		"boolean", "byte", "char",
		"double", "float", "int",
		"long", "short", "String",
	}

	// Multi-line comments
	code = formatOpenClose(code, "/*", "*/", StyleComment)

	// Multi-line brackets
	code = formatOpenClose(code, `"""`, `"""`, StyleBrackets)

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleKeyword)
	}

	// Operators
	for _, word := range operators {
		code = formatWord(code, word, defaultOperatorChars, defaultOperatorChars, StyleOperator)
	}

	// Varibles types
	for _, word := range varTypes {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleVarType)
	}

	// Numbers
	code = formatNumber(code, defaultNumberChars, defaultNumberChars)

	// true, false and null
	for _, word := range values {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleValue)
	}

	return code
}
