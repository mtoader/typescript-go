package tsc

import "github.com/microsoft/typescript-go/internal/execute/tsc"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CommandLineResult = tsc.CommandLineResult
type CommandLineTesting = tsc.CommandLineTesting
type CompileAndEmitResult = tsc.CompileAndEmitResult
type CompileTimes = tsc.CompileTimes
type DiagnosticReporter = tsc.DiagnosticReporter
type DiagnosticsReporter = tsc.DiagnosticsReporter
type EmitInput = tsc.EmitInput
type ExitStatus = tsc.ExitStatus
type ExtendedConfigCache = tsc.ExtendedConfigCache
type Statistics = tsc.Statistics
type System = tsc.System
type Watcher = tsc.Watcher

// Functions (exported as variables)
var CreateBuilderStatusReporter = tsc.CreateBuilderStatusReporter
var CreateDiagnosticReporter = tsc.CreateDiagnosticReporter
var CreateReportErrorSummary = tsc.CreateReportErrorSummary
var CreateWatchStatusReporter = tsc.CreateWatchStatusReporter
var EmitAndReportStatistics = tsc.EmitAndReportStatistics
var EmitFilesAndReportErrors = tsc.EmitFilesAndReportErrors
var GetTraceWithWriterFromSys = tsc.GetTraceWithWriterFromSys
var PrintBuildHelp = tsc.PrintBuildHelp
var PrintHelp = tsc.PrintHelp
var PrintVersion = tsc.PrintVersion
var QuietDiagnosticReporter = tsc.QuietDiagnosticReporter
var QuietDiagnosticsReporter = tsc.QuietDiagnosticsReporter
var WriteConfigFile = tsc.WriteConfigFile

// Constants
const ExitStatusDiagnosticsPresent_OutputsGenerated = tsc.ExitStatusDiagnosticsPresent_OutputsGenerated
const ExitStatusDiagnosticsPresent_OutputsSkipped = tsc.ExitStatusDiagnosticsPresent_OutputsSkipped
const ExitStatusInvalidProject_OutputsSkipped = tsc.ExitStatusInvalidProject_OutputsSkipped
const ExitStatusNotImplemented = tsc.ExitStatusNotImplemented
const ExitStatusProjectReferenceCycle_OutputsSkipped = tsc.ExitStatusProjectReferenceCycle_OutputsSkipped
const ExitStatusSuccess = tsc.ExitStatusSuccess

