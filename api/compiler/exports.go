package compiler

import "github.com/microsoft/typescript-go/internal/compiler"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CheckerPool = compiler.CheckerPool
type CompilerHost = compiler.CompilerHost
type EmitHost = compiler.EmitHost
type EmitOnly = compiler.EmitOnly
type EmitOptions = compiler.EmitOptions
type EmitResult = compiler.EmitResult
type FileIncludeReason = compiler.FileIncludeReason
type LibFile = compiler.LibFile
type Program = compiler.Program
type ProgramLike = compiler.ProgramLike
type ProgramOptions = compiler.ProgramOptions
type SourceFileMayBeEmittedHost = compiler.SourceFileMayBeEmittedHost
type SourceMapEmitResult = compiler.SourceMapEmitResult
type WriteFile = compiler.WriteFile
type WriteFileData = compiler.WriteFileData

// Functions (exported as variables)
var CombineEmitResults = compiler.CombineEmitResults
var FilterNoEmitSemanticDiagnostics = compiler.FilterNoEmitSemanticDiagnostics
var GetDiagnosticsOfAnyProgram = compiler.GetDiagnosticsOfAnyProgram
var HandleNoEmitOnError = compiler.HandleNoEmitOnError
var NewCachedFSCompilerHost = compiler.NewCachedFSCompilerHost
var NewCompilerHost = compiler.NewCompilerHost
var NewProgram = compiler.NewProgram
var SortAndDeduplicateDiagnostics = compiler.SortAndDeduplicateDiagnostics

// Constants
const EmitAll = compiler.EmitAll
const EmitOnlyDts = compiler.EmitOnlyDts
const EmitOnlyForcedDts = compiler.EmitOnlyForcedDts
const EmitOnlyJs = compiler.EmitOnlyJs

