package testutil

import "github.com/microsoft/typescript-go/internal/testutil"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var AssertPanics = testutil.AssertPanics
var RecoverAndFail = testutil.RecoverAndFail
var TestProgramIsSingleThreaded = testutil.TestProgramIsSingleThreaded

