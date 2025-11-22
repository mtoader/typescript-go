package filefixture

import "github.com/microsoft/typescript-go/internal/testutil/filefixture"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Fixture = filefixture.Fixture

// Functions (exported as variables)
var FromFile = filefixture.FromFile
var FromString = filefixture.FromString

