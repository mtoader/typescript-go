package vfs

import "github.com/microsoft/typescript-go/internal/vfs"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type DirEntry = vfs.DirEntry
type Entries = vfs.Entries
type FS = vfs.FS
type FileInfo = vfs.FileInfo
type FileMatcherPatterns = vfs.FileMatcherPatterns
type Usage = vfs.Usage
type WalkDirFunc = vfs.WalkDirFunc
type WildcardMatcher = vfs.WildcardMatcher

// Functions (exported as variables)
var GetPatternFromSpec = vfs.GetPatternFromSpec
var GetRegexFromPattern = vfs.GetRegexFromPattern
var GetRegularExpressionForWildcard = vfs.GetRegularExpressionForWildcard
var GetRegularExpressionsForWildcards = vfs.GetRegularExpressionsForWildcards
var IsImplicitGlob = vfs.IsImplicitGlob
var ReadDirectory = vfs.ReadDirectory

// Variables
var ErrClosed = vfs.ErrClosed
var ErrExist = vfs.ErrExist
var ErrInvalid = vfs.ErrInvalid
var ErrNotExist = vfs.ErrNotExist
var ErrPermission = vfs.ErrPermission
var SkipAll = vfs.SkipAll
var SkipDir = vfs.SkipDir

