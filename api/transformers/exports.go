package transformers

import "github.com/microsoft/typescript-go/internal/transformers"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type TransformOptions = transformers.TransformOptions
type Transformer = transformers.Transformer
type TransformerFactory = transformers.TransformerFactory

// Functions (exported as variables)
var Chain = transformers.Chain
var ConvertBindingPatternToAssignmentPattern = transformers.ConvertBindingPatternToAssignmentPattern
var ConvertVariableDeclarationToAssignmentExpression = transformers.ConvertVariableDeclarationToAssignmentExpression
var ExtractModifiers = transformers.ExtractModifiers
var IsExportName = transformers.IsExportName
var IsGeneratedIdentifier = transformers.IsGeneratedIdentifier
var IsHelperName = transformers.IsHelperName
var IsIdentifierReference = transformers.IsIdentifierReference
var IsLocalName = transformers.IsLocalName
var IsOriginalNodeSingleLine = transformers.IsOriginalNodeSingleLine
var IsSimpleCopiableExpression = transformers.IsSimpleCopiableExpression
var IsSimpleInlineableExpression = transformers.IsSimpleInlineableExpression
var SingleOrMany = transformers.SingleOrMany

