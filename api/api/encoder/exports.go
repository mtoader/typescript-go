package encoder

import "github.com/microsoft/typescript-go/internal/api/encoder"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Functions (exported as variables)
var EncodeSourceFile = encoder.EncodeSourceFile

// Constants
const HeaderOffsetExtendedData = encoder.HeaderOffsetExtendedData
const HeaderOffsetMetadata = encoder.HeaderOffsetMetadata
const HeaderOffsetNodes = encoder.HeaderOffsetNodes
const HeaderOffsetStringData = encoder.HeaderOffsetStringData
const HeaderOffsetStringOffsets = encoder.HeaderOffsetStringOffsets
const HeaderSize = encoder.HeaderSize
const NodeDataChildMask = encoder.NodeDataChildMask
const NodeDataStringIndexMask = encoder.NodeDataStringIndexMask
const NodeDataTypeChildren = encoder.NodeDataTypeChildren
const NodeDataTypeExtendedData = encoder.NodeDataTypeExtendedData
const NodeDataTypeMask = encoder.NodeDataTypeMask
const NodeDataTypeString = encoder.NodeDataTypeString
const NodeOffsetData = encoder.NodeOffsetData
const NodeOffsetEnd = encoder.NodeOffsetEnd
const NodeOffsetKind = encoder.NodeOffsetKind
const NodeOffsetNext = encoder.NodeOffsetNext
const NodeOffsetParent = encoder.NodeOffsetParent
const NodeOffsetPos = encoder.NodeOffsetPos
const NodeSize = encoder.NodeSize
const ProtocolVersion = encoder.ProtocolVersion
const SyntaxKindNodeList = encoder.SyntaxKindNodeList

