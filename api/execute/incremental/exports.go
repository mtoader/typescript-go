package incremental

import "github.com/microsoft/typescript-go/internal/execute/incremental"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type BuildInfo = incremental.BuildInfo
type BuildInfoDiagnostic = incremental.BuildInfoDiagnostic
type BuildInfoDiagnosticsOfFile = incremental.BuildInfoDiagnosticsOfFile
type BuildInfoEmitSignature = incremental.BuildInfoEmitSignature
type BuildInfoFileId = incremental.BuildInfoFileId
type BuildInfoFileIdListId = incremental.BuildInfoFileIdListId
type BuildInfoFileInfo = incremental.BuildInfoFileInfo
type BuildInfoFilePendingEmit = incremental.BuildInfoFilePendingEmit
type BuildInfoReader = incremental.BuildInfoReader
type BuildInfoReferenceMapEntry = incremental.BuildInfoReferenceMapEntry
type BuildInfoResolvedRoot = incremental.BuildInfoResolvedRoot
type BuildInfoRoot = incremental.BuildInfoRoot
type BuildInfoRootInfoReader = incremental.BuildInfoRootInfoReader
type BuildInfoSemanticDiagnostic = incremental.BuildInfoSemanticDiagnostic
type DiagnosticsOrBuildInfoDiagnosticsWithFileName = incremental.DiagnosticsOrBuildInfoDiagnosticsWithFileName
type FileEmitKind = incremental.FileEmitKind
type FileInfo = incremental.FileInfo
type Host = incremental.Host
type Program = incremental.Program
type SignatureUpdateKind = incremental.SignatureUpdateKind
type TestingData = incremental.TestingData

// Functions (exported as variables)
var ComputeHash = incremental.ComputeHash
var CreateHost = incremental.CreateHost
var GetFileEmitKind = incremental.GetFileEmitKind
var GetMTime = incremental.GetMTime
var NewBuildInfoReader = incremental.NewBuildInfoReader
var NewProgram = incremental.NewProgram
var ReadBuildInfoProgram = incremental.ReadBuildInfoProgram

// Constants
const FileEmitKindAll = incremental.FileEmitKindAll
const FileEmitKindAllDts = incremental.FileEmitKindAllDts
const FileEmitKindAllDtsEmit = incremental.FileEmitKindAllDtsEmit
const FileEmitKindAllJs = incremental.FileEmitKindAllJs
const FileEmitKindDts = incremental.FileEmitKindDts
const FileEmitKindDtsEmit = incremental.FileEmitKindDtsEmit
const FileEmitKindDtsErrors = incremental.FileEmitKindDtsErrors
const FileEmitKindDtsMap = incremental.FileEmitKindDtsMap
const FileEmitKindJs = incremental.FileEmitKindJs
const FileEmitKindJsInlineMap = incremental.FileEmitKindJsInlineMap
const FileEmitKindJsMap = incremental.FileEmitKindJsMap
const FileEmitKindNone = incremental.FileEmitKindNone
const SignatureUpdateKindComputedDts = incremental.SignatureUpdateKindComputedDts
const SignatureUpdateKindStoredAtEmit = incremental.SignatureUpdateKindStoredAtEmit
const SignatureUpdateKindUsedVersion = incremental.SignatureUpdateKindUsedVersion

