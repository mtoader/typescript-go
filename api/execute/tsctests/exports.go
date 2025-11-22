package tsctests

import "github.com/microsoft/typescript-go/internal/execute/tsctests"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type FileMap = tsctests.FileMap
type TestClock = tsctests.TestClock
type TestSys = tsctests.TestSys

// Functions (exported as variables)
var GetFileMapWithBuild = tsctests.GetFileMapWithBuild
var NewTscSystem = tsctests.NewTscSystem

