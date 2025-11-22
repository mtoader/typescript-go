package binder

import "github.com/microsoft/typescript-go/internal/binder"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ActiveLabel = binder.ActiveLabel
type Binder = binder.Binder
type ContainerFlags = binder.ContainerFlags
type NameResolver = binder.NameResolver
type ReferenceResolver = binder.ReferenceResolver
type ReferenceResolverHooks = binder.ReferenceResolverHooks

// Functions (exported as variables)
var BindSourceFile = binder.BindSourceFile
var FindUseStrictPrologue = binder.FindUseStrictPrologue
var GetContainerFlags = binder.GetContainerFlags
var GetLocalSymbolForExportDefault = binder.GetLocalSymbolForExportDefault
var GetSymbolNameForPrivateIdentifier = binder.GetSymbolNameForPrivateIdentifier
var NewReferenceResolver = binder.NewReferenceResolver
var SetValueDeclaration = binder.SetValueDeclaration

// Constants
const ContainerFlagsHasLocals = binder.ContainerFlagsHasLocals
const ContainerFlagsIsBlockScopedContainer = binder.ContainerFlagsIsBlockScopedContainer
const ContainerFlagsIsContainer = binder.ContainerFlagsIsContainer
const ContainerFlagsIsControlFlowContainer = binder.ContainerFlagsIsControlFlowContainer
const ContainerFlagsIsFunctionExpression = binder.ContainerFlagsIsFunctionExpression
const ContainerFlagsIsFunctionLike = binder.ContainerFlagsIsFunctionLike
const ContainerFlagsIsInterface = binder.ContainerFlagsIsInterface
const ContainerFlagsIsObjectLiteralOrClassExpressionMethodOrAccessor = binder.ContainerFlagsIsObjectLiteralOrClassExpressionMethodOrAccessor
const ContainerFlagsIsThisContainer = binder.ContainerFlagsIsThisContainer
const ContainerFlagsNone = binder.ContainerFlagsNone

