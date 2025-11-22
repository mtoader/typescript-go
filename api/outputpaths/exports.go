package outputpaths

import "github.com/microsoft/typescript-go/internal/outputpaths"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type OutputPaths = outputpaths.OutputPaths
type OutputPathsHost = outputpaths.OutputPathsHost

// Functions (exported as variables)
var ForEachEmittedFile = outputpaths.ForEachEmittedFile
var GetBuildInfoFileName = outputpaths.GetBuildInfoFileName
var GetCommonSourceDirectory = outputpaths.GetCommonSourceDirectory
var GetDeclarationEmitOutputFilePath = outputpaths.GetDeclarationEmitOutputFilePath
var GetOutputDeclarationFileNameWorker = outputpaths.GetOutputDeclarationFileNameWorker
var GetOutputExtension = outputpaths.GetOutputExtension
var GetOutputJSFileName = outputpaths.GetOutputJSFileName
var GetOutputJSFileNameWorker = outputpaths.GetOutputJSFileNameWorker
var GetOutputPathsFor = outputpaths.GetOutputPathsFor
var GetSourceFilePathInNewDir = outputpaths.GetSourceFilePathInNewDir
var GetSourceFilePathInNewDirWorker = outputpaths.GetSourceFilePathInNewDirWorker
var GetSourceMapFilePath = outputpaths.GetSourceMapFilePath

