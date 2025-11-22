package astnav

import "github.com/microsoft/typescript-go/internal/astnav"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var FindChildOfKind = astnav.FindChildOfKind
var FindNextToken = astnav.FindNextToken
var FindPrecedingToken = astnav.FindPrecedingToken
var FindPrecedingTokenEx = astnav.FindPrecedingTokenEx
var GetStartOfNode = astnav.GetStartOfNode
var GetTokenAtPosition = astnav.GetTokenAtPosition
var GetTouchingPropertyName = astnav.GetTouchingPropertyName
var GetTouchingToken = astnav.GetTouchingToken
var VisitEachChildAndJSDoc = astnav.VisitEachChildAndJSDoc

