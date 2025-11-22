package tstransforms

import "github.com/microsoft/typescript-go/internal/transformers/tstransforms"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ImportElisionTransformer = tstransforms.ImportElisionTransformer
type RuntimeSyntaxTransformer = tstransforms.RuntimeSyntaxTransformer
type TypeEraserTransformer = tstransforms.TypeEraserTransformer

// Functions (exported as variables)
var NewImportElisionTransformer = tstransforms.NewImportElisionTransformer
var NewRuntimeSyntaxTransformer = tstransforms.NewRuntimeSyntaxTransformer
var NewTypeEraserTransformer = tstransforms.NewTypeEraserTransformer

