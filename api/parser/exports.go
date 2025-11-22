package parser

import "github.com/microsoft/typescript-go/internal/parser"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type ParseFlags = parser.ParseFlags
type Parser = parser.Parser
type ParserState = parser.ParserState
type ParsingContext = parser.ParsingContext
type ParsingContexts = parser.ParsingContexts

// Functions (exported as variables)
var GetJSDocCommentRanges = parser.GetJSDocCommentRanges
var ParseIsolatedEntityName = parser.ParseIsolatedEntityName
var ParseSourceFile = parser.ParseSourceFile

// Constants
const PCArgumentExpressions = parser.PCArgumentExpressions
const PCArrayBindingElements = parser.PCArrayBindingElements
const PCArrayLiteralMembers = parser.PCArrayLiteralMembers
const PCBlockStatements = parser.PCBlockStatements
const PCClassMembers = parser.PCClassMembers
const PCCount = parser.PCCount
const PCEnumMembers = parser.PCEnumMembers
const PCHeritageClauseElement = parser.PCHeritageClauseElement
const PCHeritageClauses = parser.PCHeritageClauses
const PCImportAttributes = parser.PCImportAttributes
const PCImportOrExportSpecifiers = parser.PCImportOrExportSpecifiers
const PCJSDocComment = parser.PCJSDocComment
const PCJSDocParameters = parser.PCJSDocParameters
const PCJsxAttributes = parser.PCJsxAttributes
const PCJsxChildren = parser.PCJsxChildren
const PCObjectBindingElements = parser.PCObjectBindingElements
const PCObjectLiteralMembers = parser.PCObjectLiteralMembers
const PCParameters = parser.PCParameters
const PCRestProperties = parser.PCRestProperties
const PCSourceElements = parser.PCSourceElements
const PCSwitchClauseStatements = parser.PCSwitchClauseStatements
const PCSwitchClauses = parser.PCSwitchClauses
const PCTupleElementTypes = parser.PCTupleElementTypes
const PCTypeArguments = parser.PCTypeArguments
const PCTypeMembers = parser.PCTypeMembers
const PCTypeParameters = parser.PCTypeParameters
const PCVariableDeclarations = parser.PCVariableDeclarations
const ParseFlagsAwait = parser.ParseFlagsAwait
const ParseFlagsIgnoreMissingOpenBrace = parser.ParseFlagsIgnoreMissingOpenBrace
const ParseFlagsJSDoc = parser.ParseFlagsJSDoc
const ParseFlagsNone = parser.ParseFlagsNone
const ParseFlagsType = parser.ParseFlagsType
const ParseFlagsYield = parser.ParseFlagsYield

