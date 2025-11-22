package tspath

import "github.com/microsoft/typescript-go/internal/tspath"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ComparePathsOptions = tspath.ComparePathsOptions
type Path = tspath.Path

// Functions (exported as variables)
var ChangeExtension = tspath.ChangeExtension
var ChangeFullExtension = tspath.ChangeFullExtension
var CombinePaths = tspath.CombinePaths
var CompareNumberOfDirectorySeparators = tspath.CompareNumberOfDirectorySeparators
var ComparePaths = tspath.ComparePaths
var ComparePathsCaseInsensitive = tspath.ComparePathsCaseInsensitive
var ComparePathsCaseSensitive = tspath.ComparePathsCaseSensitive
var ContainsIgnoredPath = tspath.ContainsIgnoredPath
var ContainsPath = tspath.ContainsPath
var ConvertToRelativePath = tspath.ConvertToRelativePath
var EnsurePathIsNonModuleName = tspath.EnsurePathIsNonModuleName
var EnsureTrailingDirectorySeparator = tspath.EnsureTrailingDirectorySeparator
var ExtensionIsOneOf = tspath.ExtensionIsOneOf
var ExtensionIsTs = tspath.ExtensionIsTs
var FileExtensionIs = tspath.FileExtensionIs
var FileExtensionIsOneOf = tspath.FileExtensionIsOneOf
var GetAnyExtensionFromPath = tspath.GetAnyExtensionFromPath
var GetBaseFileName = tspath.GetBaseFileName
var GetCanonicalFileName = tspath.GetCanonicalFileName
var GetCommonParents = tspath.GetCommonParents
var GetDeclarationEmitExtensionForPath = tspath.GetDeclarationEmitExtensionForPath
var GetDeclarationFileExtension = tspath.GetDeclarationFileExtension
var GetDirectoryPath = tspath.GetDirectoryPath
var GetEncodedRootLength = tspath.GetEncodedRootLength
var GetNormalizedAbsolutePath = tspath.GetNormalizedAbsolutePath
var GetNormalizedPathComponents = tspath.GetNormalizedPathComponents
var GetPathComponents = tspath.GetPathComponents
var GetPathComponentsRelativeTo = tspath.GetPathComponentsRelativeTo
var GetPathFromPathComponents = tspath.GetPathFromPathComponents
var GetRelativePathFromDirectory = tspath.GetRelativePathFromDirectory
var GetRelativePathFromFile = tspath.GetRelativePathFromFile
var GetRelativePathToDirectoryOrUrl = tspath.GetRelativePathToDirectoryOrUrl
var GetRootLength = tspath.GetRootLength
var HasExtension = tspath.HasExtension
var HasImplementationTSFileExtension = tspath.HasImplementationTSFileExtension
var HasJSFileExtension = tspath.HasJSFileExtension
var HasJSONFileExtension = tspath.HasJSONFileExtension
var HasTSFileExtension = tspath.HasTSFileExtension
var HasTrailingDirectorySeparator = tspath.HasTrailingDirectorySeparator
var IsDeclarationFileName = tspath.IsDeclarationFileName
var IsDiskPathRoot = tspath.IsDiskPathRoot
var IsExternalModuleNameRelative = tspath.IsExternalModuleNameRelative
var IsRootedDiskPath = tspath.IsRootedDiskPath
var IsUrl = tspath.IsUrl
var IsVolumeCharacter = tspath.IsVolumeCharacter
var NormalizePath = tspath.NormalizePath
var NormalizeSlashes = tspath.NormalizeSlashes
var PathIsAbsolute = tspath.PathIsAbsolute
var PathIsRelative = tspath.PathIsRelative
var RemoveExtension = tspath.RemoveExtension
var RemoveFileExtension = tspath.RemoveFileExtension
var RemoveTrailingDirectorySeparator = tspath.RemoveTrailingDirectorySeparator
var RemoveTrailingDirectorySeparators = tspath.RemoveTrailingDirectorySeparators
var ResolvePath = tspath.ResolvePath
var ResolveTripleslashReference = tspath.ResolveTripleslashReference
var SplitVolumePath = tspath.SplitVolumePath
var ToFileNameLowerCase = tspath.ToFileNameLowerCase
var ToPath = tspath.ToPath
var TryExtractTSExtension = tspath.TryExtractTSExtension
var TryGetExtensionFromPath = tspath.TryGetExtensionFromPath

// Constants
const DirectorySeparator = tspath.DirectorySeparator
const ExtensionCjs = tspath.ExtensionCjs
const ExtensionCts = tspath.ExtensionCts
const ExtensionDcts = tspath.ExtensionDcts
const ExtensionDmts = tspath.ExtensionDmts
const ExtensionDts = tspath.ExtensionDts
const ExtensionJs = tspath.ExtensionJs
const ExtensionJson = tspath.ExtensionJson
const ExtensionJsx = tspath.ExtensionJsx
const ExtensionMjs = tspath.ExtensionMjs
const ExtensionMts = tspath.ExtensionMts
const ExtensionTs = tspath.ExtensionTs
const ExtensionTsBuildInfo = tspath.ExtensionTsBuildInfo
const ExtensionTsx = tspath.ExtensionTsx

// Variables
var AllSupportedExtensions = tspath.AllSupportedExtensions
var AllSupportedExtensionsWithJson = tspath.AllSupportedExtensionsWithJson
var ExtensionsNotSupportingExtensionlessResolution = tspath.ExtensionsNotSupportingExtensionlessResolution
var SupportedJSExtensions = tspath.SupportedJSExtensions
var SupportedJSExtensionsFlat = tspath.SupportedJSExtensionsFlat
var SupportedTSExtensions = tspath.SupportedTSExtensions
var SupportedTSExtensionsFlat = tspath.SupportedTSExtensionsFlat
var SupportedTSExtensionsWithJson = tspath.SupportedTSExtensionsWithJson
var SupportedTSExtensionsWithJsonFlat = tspath.SupportedTSExtensionsWithJsonFlat

