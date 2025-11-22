package tsoptions

import "github.com/microsoft/typescript-go/internal/tsoptions"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type AlternateModeDiagnostics = tsoptions.AlternateModeDiagnostics
type CommandLineOption = tsoptions.CommandLineOption
type CommandLineOptionKind = tsoptions.CommandLineOptionKind
type CommandLineOptionNameMap = tsoptions.CommandLineOptionNameMap
type CompilerOptionsValue = tsoptions.CompilerOptionsValue
type DidYouMeanOptionsDiagnostics = tsoptions.DidYouMeanOptionsDiagnostics
type ExtendedConfigCache = tsoptions.ExtendedConfigCache
type ExtendedConfigCacheEntry = tsoptions.ExtendedConfigCacheEntry
type FileExtensionInfo = tsoptions.FileExtensionInfo
type NameMap = tsoptions.NameMap
type ParseCommandLineWorkerDiagnostics = tsoptions.ParseCommandLineWorkerDiagnostics
type ParseConfigHost = tsoptions.ParseConfigHost
type ParsedBuildCommandLine = tsoptions.ParsedBuildCommandLine
type ParsedCommandLine = tsoptions.ParsedCommandLine
type SourceOutputAndProjectReference = tsoptions.SourceOutputAndProjectReference
type TsConfigSourceFile = tsoptions.TsConfigSourceFile

// Functions (exported as variables)
var CompilerOptionsAffectDeclarationPath = tsoptions.CompilerOptionsAffectDeclarationPath
var CompilerOptionsAffectEmit = tsoptions.CompilerOptionsAffectEmit
var CompilerOptionsAffectSemanticDiagnostics = tsoptions.CompilerOptionsAffectSemanticDiagnostics
var ConvertOptionToAbsolutePath = tsoptions.ConvertOptionToAbsolutePath
var CreateDiagnosticAtReferenceSyntax = tsoptions.CreateDiagnosticAtReferenceSyntax
var CreateDiagnosticForNodeInSourceFile = tsoptions.CreateDiagnosticForNodeInSourceFile
var CreateDiagnosticForNodeInSourceFileOrCompilerDiagnostic = tsoptions.CreateDiagnosticForNodeInSourceFileOrCompilerDiagnostic
var ForEachCompilerOptionValue = tsoptions.ForEachCompilerOptionValue
var GetCallbackForFindingPropertyAssignmentByValue = tsoptions.GetCallbackForFindingPropertyAssignmentByValue
var GetDefaultLibFileName = tsoptions.GetDefaultLibFileName
var GetLibFileName = tsoptions.GetLibFileName
var GetNameMapFromList = tsoptions.GetNameMapFromList
var GetOptionsSyntaxByArrayElementValue = tsoptions.GetOptionsSyntaxByArrayElementValue
var GetParsedCommandLineOfConfigFile = tsoptions.GetParsedCommandLineOfConfigFile
var GetParsedCommandLineOfConfigFilePath = tsoptions.GetParsedCommandLineOfConfigFilePath
var GetSupportedExtensions = tsoptions.GetSupportedExtensions
var GetSupportedExtensionsWithJsonIfResolveJsonModule = tsoptions.GetSupportedExtensionsWithJsonIfResolveJsonModule
var GetTsConfigPropArrayElementValue = tsoptions.GetTsConfigPropArrayElementValue
var NewParsedCommandLine = tsoptions.NewParsedCommandLine
var NewTsconfigSourceFileFromFilePath = tsoptions.NewTsconfigSourceFileFromFilePath
var ParseBuildCommandLine = tsoptions.ParseBuildCommandLine
var ParseBuildOptions = tsoptions.ParseBuildOptions
var ParseCommandLine = tsoptions.ParseCommandLine
var ParseCompilerOptions = tsoptions.ParseCompilerOptions
var ParseConfigFileTextToJson = tsoptions.ParseConfigFileTextToJson
var ParseJsonConfigFileContent = tsoptions.ParseJsonConfigFileContent
var ParseJsonSourceFileConfigFileContent = tsoptions.ParseJsonSourceFileConfigFileContent
var ParseListTypeOption = tsoptions.ParseListTypeOption
var ParseString = tsoptions.ParseString
var ParseStringArray = tsoptions.ParseStringArray
var ParseTristate = tsoptions.ParseTristate
var ParseTypeAcquisition = tsoptions.ParseTypeAcquisition
var ParseWatchOptions = tsoptions.ParseWatchOptions
var TargetToLibMap = tsoptions.TargetToLibMap

// Constants
const CommandLineOptionTypeBoolean = tsoptions.CommandLineOptionTypeBoolean
const CommandLineOptionTypeEnum = tsoptions.CommandLineOptionTypeEnum
const CommandLineOptionTypeList = tsoptions.CommandLineOptionTypeList
const CommandLineOptionTypeListOrElement = tsoptions.CommandLineOptionTypeListOrElement
const CommandLineOptionTypeNumber = tsoptions.CommandLineOptionTypeNumber
const CommandLineOptionTypeObject = tsoptions.CommandLineOptionTypeObject
const CommandLineOptionTypeString = tsoptions.CommandLineOptionTypeString

// Variables
var BuildNameMap = tsoptions.BuildNameMap
var BuildOpts = tsoptions.BuildOpts
var CommandLineCompilerOptionsMap = tsoptions.CommandLineCompilerOptionsMap
var CompilerNameMap = tsoptions.CompilerNameMap
var CompilerOptionsDidYouMeanDiagnostics = tsoptions.CompilerOptionsDidYouMeanDiagnostics
var InverseJsxOptionMap = tsoptions.InverseJsxOptionMap
var LibFilesSet = tsoptions.LibFilesSet
var LibMap = tsoptions.LibMap
var Libs = tsoptions.Libs
var OptionsDeclarations = tsoptions.OptionsDeclarations
var OptionsForBuild = tsoptions.OptionsForBuild
var OptionsForWatch = tsoptions.OptionsForWatch
var TscBuildOption = tsoptions.TscBuildOption
var WatchNameMap = tsoptions.WatchNameMap

