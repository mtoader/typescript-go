package declarations

import "github.com/microsoft/typescript-go/internal/transformers/declarations"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type DeclarationEmitHost = declarations.DeclarationEmitHost
type DeclarationTransformer = declarations.DeclarationTransformer
type GetSymbolAccessibilityDiagnostic = declarations.GetSymbolAccessibilityDiagnostic
type OutputPaths = declarations.OutputPaths
type ReferencedFilePair = declarations.ReferencedFilePair
type SymbolAccessibilityDiagnostic = declarations.SymbolAccessibilityDiagnostic
type SymbolTrackerImpl = declarations.SymbolTrackerImpl
type SymbolTrackerSharedState = declarations.SymbolTrackerSharedState

// Functions (exported as variables)
var NewDeclarationTransformer = declarations.NewDeclarationTransformer
var NewSymbolTracker = declarations.NewSymbolTracker

