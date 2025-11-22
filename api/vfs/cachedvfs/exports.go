package cachedvfs

import "github.com/microsoft/typescript-go/internal/vfs/cachedvfs"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type FS = cachedvfs.FS

// Functions (exported as variables)
var From = cachedvfs.From

