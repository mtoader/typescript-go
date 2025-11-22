package baseline

import "github.com/microsoft/typescript-go/internal/testutil/baseline"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Options = baseline.Options

// Functions (exported as variables)
var DiffText = baseline.DiffText
var Run = baseline.Run
var RunAgainstSubmodule = baseline.RunAgainstSubmodule

// Constants
const NoContent = baseline.NoContent

