package sourcemap

import "github.com/microsoft/typescript-go/internal/sourcemap"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type DocumentPosition = sourcemap.DocumentPosition
type DocumentPositionMapper = sourcemap.DocumentPositionMapper
type ECMALineInfo = sourcemap.ECMALineInfo
type Generator = sourcemap.Generator
type Host = sourcemap.Host
type MappedPosition = sourcemap.MappedPosition
type Mapping = sourcemap.Mapping
type MappingsDecoder = sourcemap.MappingsDecoder
type NameIndex = sourcemap.NameIndex
type RawSourceMap = sourcemap.RawSourceMap
type Source = sourcemap.Source
type SourceIndex = sourcemap.SourceIndex
type SourceMappedPosition = sourcemap.SourceMappedPosition

// Functions (exported as variables)
var CreateECMALineInfo = sourcemap.CreateECMALineInfo
var DecodeMappings = sourcemap.DecodeMappings
var GetDocumentPositionMapper = sourcemap.GetDocumentPositionMapper
var NewGenerator = sourcemap.NewGenerator
var TryGetSourceMappingURL = sourcemap.TryGetSourceMappingURL

// Constants
const MissingLineOrColumn = sourcemap.MissingLineOrColumn
const MissingName = sourcemap.MissingName
const MissingSource = sourcemap.MissingSource

