package build

import "github.com/microsoft/typescript-go/internal/execute/build"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type BuildTask = build.BuildTask
type Options = build.Options
type Orchestrator = build.Orchestrator

// Functions (exported as variables)
var NewOrchestrator = build.NewOrchestrator

