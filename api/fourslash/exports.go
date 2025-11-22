package fourslash

import "github.com/microsoft/typescript-go/internal/fourslash"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ApplyCodeActionFromCompletionOptions = fourslash.ApplyCodeActionFromCompletionOptions
type CompletionsExpectedCodeAction = fourslash.CompletionsExpectedCodeAction
type CompletionsExpectedItem = fourslash.CompletionsExpectedItem
type CompletionsExpectedItemDefaults = fourslash.CompletionsExpectedItemDefaults
type CompletionsExpectedItems = fourslash.CompletionsExpectedItems
type CompletionsExpectedList = fourslash.CompletionsExpectedList
type EditRange = fourslash.EditRange
type ExpectedCompletionEditRange = fourslash.ExpectedCompletionEditRange
type FourslashTest = fourslash.FourslashTest
type Ignored = fourslash.Ignored
type Marker = fourslash.Marker
type MarkerInput = fourslash.MarkerInput
type MarkerOrRange = fourslash.MarkerOrRange
type MarkerOrRangeOrName = fourslash.MarkerOrRangeOrName
type RangeMarker = fourslash.RangeMarker
type SignatureHelpCase = fourslash.SignatureHelpCase
type TestData = fourslash.TestData
type TestFileInfo = fourslash.TestFileInfo
type VerifyCompletionsResult = fourslash.VerifyCompletionsResult

// Functions (exported as variables)
var NewFourslash = fourslash.NewFourslash
var ParseTestData = fourslash.ParseTestData

// Variables
var AnyTextEdits = fourslash.AnyTextEdits

