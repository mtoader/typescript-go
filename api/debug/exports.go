package debug

import "github.com/microsoft/typescript-go/internal/debug"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var Assert = debug.Assert
var AssertEqual = debug.AssertEqual
var AssertGreaterThan = debug.AssertGreaterThan
var AssertGreaterThanOrEqual = debug.AssertGreaterThanOrEqual
var AssertIsDefined = debug.AssertIsDefined
var AssertLessThan = debug.AssertLessThan
var AssertLessThanOrEqual = debug.AssertLessThanOrEqual
var AssertNever = debug.AssertNever
var AssertNil = debug.AssertNil
var Fail = debug.Fail
var FailBadSyntaxKind = debug.FailBadSyntaxKind

