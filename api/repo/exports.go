package repo

import "github.com/microsoft/typescript-go/internal/repo"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type SkippableTest = repo.SkippableTest

// Functions (exported as variables)
var SkipIfNoTypeScriptSubmodule = repo.SkipIfNoTypeScriptSubmodule
var TypeScriptSubmoduleExists = repo.TypeScriptSubmoduleExists

// Variables
var RootPath = repo.RootPath
var TestDataPath = repo.TestDataPath
var TypeScriptSubmodulePath = repo.TypeScriptSubmodulePath

