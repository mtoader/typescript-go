package logging

import "github.com/microsoft/typescript-go/internal/project/logging"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type LogCollector = logging.LogCollector
type LogTree = logging.LogTree
type Logger = logging.Logger

// Functions (exported as variables)
var NewLogTree = logging.NewLogTree
var NewLogger = logging.NewLogger
var NewTestLogger = logging.NewTestLogger

