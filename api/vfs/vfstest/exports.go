package vfstest

import "github.com/microsoft/typescript-go/internal/vfs/vfstest"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Clock = vfstest.Clock
type MapFS = vfstest.MapFS

// Functions (exported as variables)
var Symlink = vfstest.Symlink

