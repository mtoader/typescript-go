package format

import "github.com/microsoft/typescript-go/internal/format"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type EditorSettings = format.EditorSettings
type FormatCodeSettings = format.FormatCodeSettings
type FormatRequestKind = format.FormatRequestKind
type FormattingContext = format.FormattingContext
type IndentStyle = format.IndentStyle
type LineAction = format.LineAction
type RulesPosition = format.RulesPosition
type SemicolonPreference = format.SemicolonPreference
type TextRangeWithKind = format.TextRangeWithKind

// Functions (exported as variables)
var FindFirstNonWhitespaceColumn = format.FindFirstNonWhitespaceColumn
var FormatDocument = format.FormatDocument
var FormatNodeGivenIndentation = format.FormatNodeGivenIndentation
var FormatOnClosingCurly = format.FormatOnClosingCurly
var FormatOnEnter = format.FormatOnEnter
var FormatOnOpeningCurly = format.FormatOnOpeningCurly
var FormatOnSemicolon = format.FormatOnSemicolon
var FormatSelection = format.FormatSelection
var FormatSpan = format.FormatSpan
var GetContainingList = format.GetContainingList
var GetDefaultFormatCodeSettings = format.GetDefaultFormatCodeSettings
var GetFormatCodeSettingsFromContext = format.GetFormatCodeSettingsFromContext
var GetIndentationForNode = format.GetIndentationForNode
var GetLineStartPositionForPosition = format.GetLineStartPositionForPosition
var GetNewLineOrDefaultFromContext = format.GetNewLineOrDefaultFromContext
var NewFormattingContext = format.NewFormattingContext
var NewTextRangeWithKind = format.NewTextRangeWithKind
var NodeWillIndentChild = format.NodeWillIndentChild
var ShouldIndentChildNode = format.ShouldIndentChildNode
var WithFormatCodeSettings = format.WithFormatCodeSettings

// Constants
const FormatRequestKindFormatDocument = format.FormatRequestKindFormatDocument
const FormatRequestKindFormatOnClosingCurlyBrace = format.FormatRequestKindFormatOnClosingCurlyBrace
const FormatRequestKindFormatOnEnter = format.FormatRequestKindFormatOnEnter
const FormatRequestKindFormatOnOpeningCurlyBrace = format.FormatRequestKindFormatOnOpeningCurlyBrace
const FormatRequestKindFormatOnSemicolon = format.FormatRequestKindFormatOnSemicolon
const FormatRequestKindFormatSelection = format.FormatRequestKindFormatSelection
const IndentStyleBlock = format.IndentStyleBlock
const IndentStyleNone = format.IndentStyleNone
const IndentStyleSmart = format.IndentStyleSmart
const LineActionLineAdded = format.LineActionLineAdded
const LineActionLineRemoved = format.LineActionLineRemoved
const LineActionNone = format.LineActionNone
const RulesPositionContextRulesAny = format.RulesPositionContextRulesAny
const RulesPositionContextRulesSpecific = format.RulesPositionContextRulesSpecific
const RulesPositionNoContextRulesAny = format.RulesPositionNoContextRulesAny
const RulesPositionNoContextRulesSpecific = format.RulesPositionNoContextRulesSpecific
const RulesPositionStopRulesAny = format.RulesPositionStopRulesAny
const RulesPositionStopRulesSpecific = format.RulesPositionStopRulesSpecific
const SemicolonPreferenceIgnore = format.SemicolonPreferenceIgnore
const SemicolonPreferenceInsert = format.SemicolonPreferenceInsert
const SemicolonPreferenceRemove = format.SemicolonPreferenceRemove

