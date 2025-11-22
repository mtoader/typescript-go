package iovfs

import "github.com/microsoft/typescript-go/internal/vfs/iovfs"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type FsWithSys = iovfs.FsWithSys
type RealpathFS = iovfs.RealpathFS
type WritableFS = iovfs.WritableFS

// Functions (exported as variables)
var From = iovfs.From

