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
//
// Supported preprocessor directives (const StyleKeyword):
//   #define
//   #elif
//   #else
//   #endif
//   #error
//   #if
//   #ifdef
//   #ifndef
//   #import
//   #include
//   #line
//   #pragma
//   #undef
//   #using
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

	preprocessor := []string{
		"#define", "#elif", "#else",
		"#endif", "#error", "#if",
		"#ifdef", "#ifndef", "#import",
		"#include", "#line", "#pragma",
		"#undef", "#using",
	}

	preprocessorChars := []string{
		" ", "\t", "\n",
	}

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

	// Multi-line comments
	code = formatOpenClose(code, "/*", "*/", StyleComment)

	// Single-line comments
	code = formatOpenClose(code, "//", "\n", StyleComment)

	// Single-line brackets
	code = formatOpenClose(code, `"`, `"`, StyleBrackets)
	code = formatOpenClose(code, `'`, `'`, StyleBrackets)

	// Preprocessor
	for _, word := range preprocessor {
		code = formatWord(code, word, preprocessorChars, preprocessorChars, StyleKeyword)
	}

	// Keywords
	for _, word := range keywords {
		code = formatWord(code, word, defaultKeywordChars, defaultKeywordChars, StyleKeyword)
	}

	// Operators
	for _, word := range operators {
		code = formatWord(code, word, defaultOperatorChars, defaultOperatorChars, StyleOperator)
	}

	return code
}
