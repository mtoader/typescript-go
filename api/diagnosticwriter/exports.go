package diagnosticwriter

import "github.com/microsoft/typescript-go/internal/diagnosticwriter"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ASTDiagnostic = diagnosticwriter.ASTDiagnostic
type Diagnostic = diagnosticwriter.Diagnostic
type ErrorSummary = diagnosticwriter.ErrorSummary
type FileLike = diagnosticwriter.FileLike
type FormattedWriter = diagnosticwriter.FormattedWriter
type FormattingOptions = diagnosticwriter.FormattingOptions

// Functions (exported as variables)
var CompareASTDiagnostics = diagnosticwriter.CompareASTDiagnostics
var FlattenDiagnosticMessage = diagnosticwriter.FlattenDiagnosticMessage
var FormatDiagnosticWithColorAndContext = diagnosticwriter.FormatDiagnosticWithColorAndContext
var FormatDiagnosticsStatusAndTime = diagnosticwriter.FormatDiagnosticsStatusAndTime
var FormatDiagnosticsStatusWithColorAndTime = diagnosticwriter.FormatDiagnosticsStatusWithColorAndTime
var FormatDiagnosticsWithColorAndContext = diagnosticwriter.FormatDiagnosticsWithColorAndContext
var FromASTDiagnostics = diagnosticwriter.FromASTDiagnostics
var TryClearScreen = diagnosticwriter.TryClearScreen
var WrapASTDiagnostic = diagnosticwriter.WrapASTDiagnostic
var WrapASTDiagnostics = diagnosticwriter.WrapASTDiagnostics
var WriteErrorSummaryText = diagnosticwriter.WriteErrorSummaryText
var WriteFlattenedASTDiagnosticMessage = diagnosticwriter.WriteFlattenedASTDiagnosticMessage
var WriteFlattenedDiagnosticMessage = diagnosticwriter.WriteFlattenedDiagnosticMessage
var WriteFormatDiagnostic = diagnosticwriter.WriteFormatDiagnostic
var WriteFormatDiagnostics = diagnosticwriter.WriteFormatDiagnostics
var WriteLocation = diagnosticwriter.WriteLocation

// Variables
var ScreenStartingCodes = diagnosticwriter.ScreenStartingCodes

