package estransforms

import "github.com/microsoft/typescript-go/internal/transformers/estransforms"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var GetESTransformer = estransforms.GetESTransformer
var NewUseStrictTransformer = estransforms.NewUseStrictTransformer

// Variables
var NewES2016Transformer = estransforms.NewES2016Transformer
var NewES2017Transformer = estransforms.NewES2017Transformer
var NewES2018Transformer = estransforms.NewES2018Transformer
var NewES2019Transformer = estransforms.NewES2019Transformer
var NewES2020Transformer = estransforms.NewES2020Transformer
var NewES2021Transformer = estransforms.NewES2021Transformer
var NewES2022Transformer = estransforms.NewES2022Transformer
var NewESNextTransformer = estransforms.NewESNextTransformer

