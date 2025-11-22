package stringutil

import "github.com/microsoft/typescript-go/internal/stringutil"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Comparison = stringutil.Comparison

// Functions (exported as variables)
var AddUTF8ByteOrderMark = stringutil.AddUTF8ByteOrderMark
var CompareStringsCaseInsensitive = stringutil.CompareStringsCaseInsensitive
var CompareStringsCaseInsensitiveEslintCompatible = stringutil.CompareStringsCaseInsensitiveEslintCompatible
var CompareStringsCaseInsensitiveThenSensitive = stringutil.CompareStringsCaseInsensitiveThenSensitive
var CompareStringsCaseSensitive = stringutil.CompareStringsCaseSensitive
var EncodeURI = stringutil.EncodeURI
var EquateStringCaseInsensitive = stringutil.EquateStringCaseInsensitive
var EquateStringCaseSensitive = stringutil.EquateStringCaseSensitive
var Format = stringutil.Format
var GetStringComparer = stringutil.GetStringComparer
var GetStringEqualityComparer = stringutil.GetStringEqualityComparer
var GuessIndentation = stringutil.GuessIndentation
var HasPrefix = stringutil.HasPrefix
var HasSuffix = stringutil.HasSuffix
var IsASCIILetter = stringutil.IsASCIILetter
var IsDigit = stringutil.IsDigit
var IsHexDigit = stringutil.IsHexDigit
var IsLineBreak = stringutil.IsLineBreak
var IsOctalDigit = stringutil.IsOctalDigit
var IsWhiteSpaceLike = stringutil.IsWhiteSpaceLike
var IsWhiteSpaceSingleLine = stringutil.IsWhiteSpaceSingleLine
var LowerFirstChar = stringutil.LowerFirstChar
var RemoveByteOrderMark = stringutil.RemoveByteOrderMark
var SplitLines = stringutil.SplitLines
var StripQuotes = stringutil.StripQuotes
var UnquoteString = stringutil.UnquoteString

// Constants
const ComparisonEqual = stringutil.ComparisonEqual
const ComparisonGreaterThan = stringutil.ComparisonGreaterThan
const ComparisonLessThan = stringutil.ComparisonLessThan

