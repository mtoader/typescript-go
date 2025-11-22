package moduletransforms

import "github.com/microsoft/typescript-go/internal/transformers/moduletransforms"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type CommonJSModuleTransformer = moduletransforms.CommonJSModuleTransformer
type ESModuleTransformer = moduletransforms.ESModuleTransformer
type ImpliedModuleTransformer = moduletransforms.ImpliedModuleTransformer

// Functions (exported as variables)
var NewCommonJSModuleTransformer = moduletransforms.NewCommonJSModuleTransformer
var NewESModuleTransformer = moduletransforms.NewESModuleTransformer
var NewImpliedModuleTransformer = moduletransforms.NewImpliedModuleTransformer

