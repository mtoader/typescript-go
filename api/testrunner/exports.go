package testrunner

import "github.com/microsoft/typescript-go/internal/testrunner"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CompilerBaselineRunner = testrunner.CompilerBaselineRunner
type CompilerTestType = testrunner.CompilerTestType
type ParseTestFilesOptions = testrunner.ParseTestFilesOptions
type Runner = testrunner.Runner

// Functions (exported as variables)
var NewCompilerBaselineRunner = testrunner.NewCompilerBaselineRunner

// Constants
const TestTypeConformance = testrunner.TestTypeConformance
const TestTypeRegression = testrunner.TestTypeRegression

