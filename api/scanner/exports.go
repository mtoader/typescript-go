package scanner

import "github.com/microsoft/typescript-go/internal/scanner"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ErrorCallback = scanner.ErrorCallback
type EscapeSequenceScanningFlags = scanner.EscapeSequenceScanningFlags
type Scanner = scanner.Scanner
type ScannerState = scanner.ScannerState
type SkipTriviaOptions = scanner.SkipTriviaOptions

// Functions (exported as variables)
var ComputeLineOfPosition = scanner.ComputeLineOfPosition
var ComputePositionOfLineAndCharacter = scanner.ComputePositionOfLineAndCharacter
var ComputePositionOfLineAndCharacterEx = scanner.ComputePositionOfLineAndCharacterEx
var DeclarationNameToString = scanner.DeclarationNameToString
var GetECMAEndLinePosition = scanner.GetECMAEndLinePosition
var GetECMALineAndCharacterOfPosition = scanner.GetECMALineAndCharacterOfPosition
var GetECMALineOfPosition = scanner.GetECMALineOfPosition
var GetECMALineStarts = scanner.GetECMALineStarts
var GetECMAPositionOfLineAndCharacter = scanner.GetECMAPositionOfLineAndCharacter
var GetErrorRangeForNode = scanner.GetErrorRangeForNode
var GetIdentifierToken = scanner.GetIdentifierToken
var GetLeadingCommentRanges = scanner.GetLeadingCommentRanges
var GetRangeOfTokenAtPosition = scanner.GetRangeOfTokenAtPosition
var GetScannerForSourceFile = scanner.GetScannerForSourceFile
var GetShebang = scanner.GetShebang
var GetSourceTextOfNodeFromSourceFile = scanner.GetSourceTextOfNodeFromSourceFile
var GetTextOfNode = scanner.GetTextOfNode
var GetTextOfNodeFromSourceText = scanner.GetTextOfNodeFromSourceText
var GetTokenPosOfNode = scanner.GetTokenPosOfNode
var GetTrailingCommentRanges = scanner.GetTrailingCommentRanges
var GetViableKeywordSuggestions = scanner.GetViableKeywordSuggestions
var IdentifierToKeywordKind = scanner.IdentifierToKeywordKind
var IsIdentifierPart = scanner.IsIdentifierPart
var IsIdentifierPartEx = scanner.IsIdentifierPartEx
var IsIdentifierStart = scanner.IsIdentifierStart
var IsIdentifierText = scanner.IsIdentifierText
var IsIntrinsicJsxName = scanner.IsIntrinsicJsxName
var IsValidIdentifier = scanner.IsValidIdentifier
var NewScanner = scanner.NewScanner
var ScanTokenAtPosition = scanner.ScanTokenAtPosition
var SkipTrivia = scanner.SkipTrivia
var SkipTriviaEx = scanner.SkipTriviaEx
var StringToToken = scanner.StringToToken
var TokenToString = scanner.TokenToString

// Constants
const EscapeSequenceScanningFlagsAllowExtendedUnicodeEscape = scanner.EscapeSequenceScanningFlagsAllowExtendedUnicodeEscape
const EscapeSequenceScanningFlagsAnnexB = scanner.EscapeSequenceScanningFlagsAnnexB
const EscapeSequenceScanningFlagsAnyUnicodeMode = scanner.EscapeSequenceScanningFlagsAnyUnicodeMode
const EscapeSequenceScanningFlagsAtomEscape = scanner.EscapeSequenceScanningFlagsAtomEscape
const EscapeSequenceScanningFlagsRegularExpression = scanner.EscapeSequenceScanningFlagsRegularExpression
const EscapeSequenceScanningFlagsReportErrors = scanner.EscapeSequenceScanningFlagsReportErrors
const EscapeSequenceScanningFlagsReportInvalidEscapeErrors = scanner.EscapeSequenceScanningFlagsReportInvalidEscapeErrors
const EscapeSequenceScanningFlagsString = scanner.EscapeSequenceScanningFlagsString

