package api

import "github.com/microsoft/typescript-go/internal/api"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type API = api.API
type APIInit = api.APIInit
type Callback = api.Callback
type ConfigFileResponse = api.ConfigFileResponse
type ConfigureParams = api.ConfigureParams
type GetSourceFileParams = api.GetSourceFileParams
type GetSymbolAtLocationParams = api.GetSymbolAtLocationParams
type GetSymbolAtPositionParams = api.GetSymbolAtPositionParams
type GetSymbolsAtLocationsParams = api.GetSymbolsAtLocationsParams
type GetSymbolsAtPositionsParams = api.GetSymbolsAtPositionsParams
type GetTypeOfSymbolParams = api.GetTypeOfSymbolParams
type GetTypesOfSymbolsParams = api.GetTypesOfSymbolsParams
type LoadProjectParams = api.LoadProjectParams
type MessagePackType = api.MessagePackType
type MessageType = api.MessageType
type Method = api.Method
type ParseConfigFileParams = api.ParseConfigFileParams
type ProjectResponse = api.ProjectResponse
type Server = api.Server
type ServerOptions = api.ServerOptions
type SymbolResponse = api.SymbolResponse
type TypeResponse = api.TypeResponse

// Functions (exported as variables)
var FileHandle = api.FileHandle
var NewAPI = api.NewAPI
var NewProjectResponse = api.NewProjectResponse
var NewServer = api.NewServer
var NewSymbolResponse = api.NewSymbolResponse
var NewTypeData = api.NewTypeData
var NodeHandle = api.NodeHandle
var ProjectHandle = api.ProjectHandle
var SymbolHandle = api.SymbolHandle
var TypeHandle = api.TypeHandle

// Constants
const CallbackDirectoryExists = api.CallbackDirectoryExists
const CallbackFileExists = api.CallbackFileExists
const CallbackGetAccessibleEntries = api.CallbackGetAccessibleEntries
const CallbackReadFile = api.CallbackReadFile
const CallbackRealpath = api.CallbackRealpath
const MessagePackTypeBin16 = api.MessagePackTypeBin16
const MessagePackTypeBin32 = api.MessagePackTypeBin32
const MessagePackTypeBin8 = api.MessagePackTypeBin8
const MessagePackTypeFixedArray3 = api.MessagePackTypeFixedArray3
const MessagePackTypeU8 = api.MessagePackTypeU8
const MessageTypeCall = api.MessageTypeCall
const MessageTypeCallError = api.MessageTypeCallError
const MessageTypeCallResponse = api.MessageTypeCallResponse
const MessageTypeError = api.MessageTypeError
const MessageTypeRequest = api.MessageTypeRequest
const MessageTypeResponse = api.MessageTypeResponse
const MessageTypeUnknown = api.MessageTypeUnknown
const MethodConfigure = api.MethodConfigure
const MethodGetSourceFile = api.MethodGetSourceFile
const MethodGetSymbolAtLocation = api.MethodGetSymbolAtLocation
const MethodGetSymbolAtPosition = api.MethodGetSymbolAtPosition
const MethodGetSymbolsAtLocations = api.MethodGetSymbolsAtLocations
const MethodGetSymbolsAtPositions = api.MethodGetSymbolsAtPositions
const MethodGetTypeOfSymbol = api.MethodGetTypeOfSymbol
const MethodGetTypesOfSymbols = api.MethodGetTypesOfSymbols
const MethodLoadProject = api.MethodLoadProject
const MethodParseConfigFile = api.MethodParseConfigFile
const MethodRelease = api.MethodRelease

// Variables
var ErrClientError = api.ErrClientError
var ErrInvalidRequest = api.ErrInvalidRequest

