package pprof

import "github.com/microsoft/typescript-go/internal/pprof"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ProfileSession = pprof.ProfileSession

// Functions (exported as variables)
var BeginProfiling = pprof.BeginProfiling

