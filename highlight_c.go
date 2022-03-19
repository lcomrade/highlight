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

// C processes C source code (*.c files).
// Read more: https://go.dev/ref/spec
//
// Supported keywords (const StyleKeyword):
//   auto
//   break
//   case
//   char
//   const
//   continue
//   default
//   do
//   double
//   else
//   enum
//   extern
//   float
//   for
//   goto
//   if
//   inline
//   int
//   long
//   register
//   restrict
//   return
//   short
//   signed
//   sizeof
//   static
//   struct
//   switch
//   typedef
//   union
//   unsigned
//   void
//   volatile
//   while
//   _Alignas
//   _Alignof
//   _Atomic
//   _Bool
//   _Complex
//   _Decimal128
//   _Decimal32
//   _Decimal64
//   _Generic
//   _Imaginary
//   _Noreturn
//   _Static_assert
//   _Thread_local
//
// Supported operators (const StyleOperator):
//   ->
//   ==
//   !=
//   <=>
//   <<=
//   =>>
//   <<
//   >>
//   <=
//   <
//   >=
//   >
//   =
//   ^
//   &&
//   &
//   ||
//   ?:
//   %=
//   &=
//   ^=
//   |=
//   ++
//   --
//   +
//   -
//   !
//   ~
//   *
//   /
//   %
//
// Also supported:
//   Single-line comments (//)
//   Multi-line comments (/* */)
//   Single-line brackets (", ')
func C(code string) string {
	// Shild HTML
	code = shieldHTML(code)

	// Keywords
	keywords := []string{
		"auto", "break", "case",
		"char", "const", "continue",
		"default", "do", "double",
		"else", "enum", "extern",
		"float", "for", "goto",
		"if", "inline", "int",
		"long", "register", "restrict",
		"return", "short", "signed",
		"sizeof", "static", "struct",
		"switch", "typedef", "union",
		"unsigned", "void", "volatile",
		"while",
		"_Alignas", "_Alignof", "_Atomic",
		"_Bool", "_Complex", "_Decimal128",
		"_Decimal32", "_Decimal64", "_Generic",
		"_Imaginary", "_Noreturn",
		"_Static_assert", "_Thread_local",
	}

	cmdChars := []string{
		"", " ", "\t", "\n", ":", ";",
		"(", ")", "{", "}",
	}

	// Operators
	operators := []string{
		"-&ampgt", "==", "!=",
		"&amplt=&ampgt", "&amplt&amplt=",
		"=&ampgt&ampgt", "&amplt&amplt",
		"&ampgt&ampgt", "&amplt=",
		"&amplt", "&ampgt=", "&ampgt",
		"=", "^", "&amp&amp", "&amp",
		"||", "?:", "%=", "&amp=",
		"^=", "|=", "++", "--",
		"+", "-", "!", "~", "*",
		"/", "%",
	}

	opsChars := []string{
		"", " ", "\t", ":", "(", ")",
		`"`, `'`, "{", "}", "_",
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "0",
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x",
		"y", "z",
		"A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X",
		"Y", "Z",
	}

	// Multi-line comments
	code = formatOpenClose(code, "/*", "*/", StyleComment)

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, cmdChars, cmdChars, StyleKeyword)
	}

	// Operators
	for _, word := range operators {
		code = formatWord(code, word, opsChars, opsChars, StyleOperator)
	}

	return code
}
