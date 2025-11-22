package lsp

import "github.com/microsoft/typescript-go/internal/lsp"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Reader = lsp.Reader
type Server = lsp.Server
type ServerOptions = lsp.ServerOptions
type Writer = lsp.Writer

// Functions (exported as variables)
var NewServer = lsp.NewServer
var ToReader = lsp.ToReader
var ToWriter = lsp.ToWriter

