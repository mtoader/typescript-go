package lsconv

import "github.com/microsoft/typescript-go/internal/ls/lsconv"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Converters = lsconv.Converters
type LSPLineMap = lsconv.LSPLineMap
type LSPLineStarts = lsconv.LSPLineStarts
type Script = lsconv.Script

// Functions (exported as variables)
var ComputeLSPLineStarts = lsconv.ComputeLSPLineStarts
var DiagnosticToLSPPull = lsconv.DiagnosticToLSPPull
var DiagnosticToLSPPush = lsconv.DiagnosticToLSPPush
var FileNameToDocumentURI = lsconv.FileNameToDocumentURI
var LanguageKindToScriptKind = lsconv.LanguageKindToScriptKind
var NewConverters = lsconv.NewConverters

