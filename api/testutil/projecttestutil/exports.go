package projecttestutil

import "github.com/microsoft/typescript-go/internal/testutil/projecttestutil"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ClientMock = projecttestutil.ClientMock
type NpmExecutorMock = projecttestutil.NpmExecutorMock
type SessionUtils = projecttestutil.SessionUtils
type TypingsInstallerOptions = projecttestutil.TypingsInstallerOptions

// Functions (exported as variables)
var GetSessionInitOptions = projecttestutil.GetSessionInitOptions
var Setup = projecttestutil.Setup
var SetupWithOptions = projecttestutil.SetupWithOptions
var SetupWithOptionsAndTypingsInstaller = projecttestutil.SetupWithOptionsAndTypingsInstaller
var SetupWithTypingsInstaller = projecttestutil.SetupWithTypingsInstaller
var TypesRegistryConfig = projecttestutil.TypesRegistryConfig
var TypesRegistryConfigText = projecttestutil.TypesRegistryConfigText
var WithRequestID = projecttestutil.WithRequestID

// Constants
const TestTypingsLocation = projecttestutil.TestTypingsLocation

