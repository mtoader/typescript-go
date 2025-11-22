package semver

import "github.com/microsoft/typescript-go/internal/semver"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type SemverParseError = semver.SemverParseError
type Version = semver.Version
type VersionRange = semver.VersionRange

// Functions (exported as variables)
var MustParse = semver.MustParse
var TryParseVersion = semver.TryParseVersion
var TryParseVersionRange = semver.TryParseVersionRange

