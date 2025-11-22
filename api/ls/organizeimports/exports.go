package organizeimports

import "github.com/microsoft/typescript-go/internal/ls/organizeimports"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var CompareImportsOrRequireStatements = organizeimports.CompareImportsOrRequireStatements
var GetImportDeclarationInsertIndex = organizeimports.GetImportDeclarationInsertIndex
var GetImportSpecifierInsertionIndex = organizeimports.GetImportSpecifierInsertionIndex
var GetNamedImportSpecifierComparer = organizeimports.GetNamedImportSpecifierComparer
var GetNamedImportSpecifierComparerWithDetection = organizeimports.GetNamedImportSpecifierComparerWithDetection
var GetOrganizeImportsStringComparerWithDetection = organizeimports.GetOrganizeImportsStringComparerWithDetection

