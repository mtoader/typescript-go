package ata

import "github.com/microsoft/typescript-go/internal/project/ata"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CachedTyping = ata.CachedTyping
type NameValidationResult = ata.NameValidationResult
type NpmExecutor = ata.NpmExecutor
type TypingsInfo = ata.TypingsInfo
type TypingsInstallRequest = ata.TypingsInstallRequest
type TypingsInstallResult = ata.TypingsInstallResult
type TypingsInstaller = ata.TypingsInstaller
type TypingsInstallerHost = ata.TypingsInstallerHost
type TypingsInstallerOptions = ata.TypingsInstallerOptions

// Functions (exported as variables)
var DiscoverTypings = ata.DiscoverTypings
var NewTypingsInstaller = ata.NewTypingsInstaller
var ValidatePackageName = ata.ValidatePackageName

// Constants
const EmptyName = ata.EmptyName
const NameContainsNonURISafeCharacters = ata.NameContainsNonURISafeCharacters
const NameOk = ata.NameOk
const NameStartsWithDot = ata.NameStartsWithDot
const NameStartsWithUnderscore = ata.NameStartsWithUnderscore
const NameTooLong = ata.NameTooLong

