package jsnum

import "github.com/microsoft/typescript-go/internal/jsnum"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type Number = jsnum.Number
type PseudoBigInt = jsnum.PseudoBigInt

// Functions (exported as variables)
var FromString = jsnum.FromString
var Inf = jsnum.Inf
var NaN = jsnum.NaN
var NewPseudoBigInt = jsnum.NewPseudoBigInt
var ParsePseudoBigInt = jsnum.ParsePseudoBigInt
var ParseValidBigInt = jsnum.ParseValidBigInt

// Constants
const MaxSafeInteger = jsnum.MaxSafeInteger
const MinSafeInteger = jsnum.MinSafeInteger

