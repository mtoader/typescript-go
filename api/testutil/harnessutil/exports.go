package harnessutil

import "github.com/microsoft/typescript-go/internal/testutil/harnessutil"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CompilationOutput = harnessutil.CompilationOutput
type CompilationResult = harnessutil.CompilationResult
type HarnessOptions = harnessutil.HarnessOptions
type NamedTestConfiguration = harnessutil.NamedTestConfiguration
type OutputRecorderFS = harnessutil.OutputRecorderFS
type SourceFileCacheKey = harnessutil.SourceFileCacheKey
type TestConfiguration = harnessutil.TestConfiguration
type TestFile = harnessutil.TestFile
type TracerForBaselining = harnessutil.TracerForBaselining

// Functions (exported as variables)
var CompileFiles = harnessutil.CompileFiles
var CompileFilesEx = harnessutil.CompileFilesEx
var EnumerateFiles = harnessutil.EnumerateFiles
var GetConfigNameFromFileName = harnessutil.GetConfigNameFromFileName
var GetFileBasedTestConfigurations = harnessutil.GetFileBasedTestConfigurations
var GetSourceFileCacheKey = harnessutil.GetSourceFileCacheKey
var NewOutputRecorderFS = harnessutil.NewOutputRecorderFS
var NewTracerForBaselining = harnessutil.NewTracerForBaselining
var SetCompilerOptionsFromTestConfig = harnessutil.SetCompilerOptionsFromTestConfig

// Constants
const FakeTSVersion = harnessutil.FakeTSVersion

