package parsetestutil

import "github.com/microsoft/typescript-go/internal/testutil/parsetestutil"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var CheckDiagnostics = parsetestutil.CheckDiagnostics
var CheckDiagnosticsMessage = parsetestutil.CheckDiagnosticsMessage
var MarkSyntheticRecursive = parsetestutil.MarkSyntheticRecursive
var ParseTypeScript = parsetestutil.ParseTypeScript

