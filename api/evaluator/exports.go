package evaluator

import "github.com/microsoft/typescript-go/internal/evaluator"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Evaluator = evaluator.Evaluator
type Result = evaluator.Result

// Functions (exported as variables)
var AnyToString = evaluator.AnyToString
var IsTruthy = evaluator.IsTruthy
var NewEvaluator = evaluator.NewEvaluator
var NewResult = evaluator.NewResult

