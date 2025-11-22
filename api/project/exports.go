package project

import "github.com/microsoft/typescript-go/internal/project"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type APISnapshotRequest = project.APISnapshotRequest
type ATAStateChange = project.ATAStateChange
type CheckerPool = project.CheckerPool
type Client = project.Client
type ConfigFileRegistry = project.ConfigFileRegistry
type CreateProgramResult = project.CreateProgramResult
type ExtendedConfigCache = project.ExtendedConfigCache
type FileChange = project.FileChange
type FileChangeKind = project.FileChangeKind
type FileChangeSummary = project.FileChangeSummary
type FileContent = project.FileContent
type FileHandle = project.FileHandle
type FileSource = project.FileSource
type Kind = project.Kind
type Overlay = project.Overlay
type ParseCache = project.ParseCache
type PatternsAndIgnored = project.PatternsAndIgnored
type PendingReload = project.PendingReload
type ProgramUpdateKind = project.ProgramUpdateKind
type Project = project.Project
type ProjectCollection = project.ProjectCollection
type ProjectCollectionBuilder = project.ProjectCollectionBuilder
type ProjectTreeRequest = project.ProjectTreeRequest
type ResourceRequest = project.ResourceRequest
type Session = project.Session
type SessionInit = project.SessionInit
type SessionOptions = project.SessionOptions
type Snapshot = project.Snapshot
type SnapshotChange = project.SnapshotChange
type SnapshotFS = project.SnapshotFS
type TestConfigEntry = project.TestConfigEntry
type TestConfigFileNamesEntry = project.TestConfigFileNamesEntry
type UpdateReason = project.UpdateReason
type WatcherID = project.WatcherID

// Functions (exported as variables)
var NewConfiguredProject = project.NewConfiguredProject
var NewInferredProject = project.NewInferredProject
var NewProject = project.NewProject
var NewSession = project.NewSession
var NewSnapshot = project.NewSnapshot

// Constants
const FileChangeKindChange = project.FileChangeKindChange
const FileChangeKindClose = project.FileChangeKindClose
const FileChangeKindOpen = project.FileChangeKindOpen
const FileChangeKindSave = project.FileChangeKindSave
const FileChangeKindWatchChange = project.FileChangeKindWatchChange
const FileChangeKindWatchCreate = project.FileChangeKindWatchCreate
const FileChangeKindWatchDelete = project.FileChangeKindWatchDelete
const KindConfigured = project.KindConfigured
const KindInferred = project.KindInferred
const PendingReloadFileNames = project.PendingReloadFileNames
const PendingReloadFull = project.PendingReloadFull
const PendingReloadNone = project.PendingReloadNone
const ProgramUpdateKindCloned = project.ProgramUpdateKindCloned
const ProgramUpdateKindNewFiles = project.ProgramUpdateKindNewFiles
const ProgramUpdateKindNone = project.ProgramUpdateKindNone
const ProgramUpdateKindSameFileNames = project.ProgramUpdateKindSameFileNames
const UpdateReasonDidChangeCompilerOptionsForInferredProjects = project.UpdateReasonDidChangeCompilerOptionsForInferredProjects
const UpdateReasonDidOpenFile = project.UpdateReasonDidOpenFile
const UpdateReasonRequestedLanguageServiceForFileNotOpen = project.UpdateReasonRequestedLanguageServiceForFileNotOpen
const UpdateReasonRequestedLanguageServicePendingChanges = project.UpdateReasonRequestedLanguageServicePendingChanges
const UpdateReasonRequestedLanguageServiceProjectDirty = project.UpdateReasonRequestedLanguageServiceProjectDirty
const UpdateReasonRequestedLanguageServiceProjectNotLoaded = project.UpdateReasonRequestedLanguageServiceProjectNotLoaded
const UpdateReasonRequestedLoadProjectTree = project.UpdateReasonRequestedLoadProjectTree
const UpdateReasonUnknown = project.UpdateReasonUnknown
