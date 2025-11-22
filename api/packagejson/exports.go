package packagejson

import "github.com/microsoft/typescript-go/internal/packagejson"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type DependencyFields = packagejson.DependencyFields
type ExportsOrImports = packagejson.ExportsOrImports
type Fields = packagejson.Fields
type HeaderFields = packagejson.HeaderFields
type InfoCache = packagejson.InfoCache
type InfoCacheEntry = packagejson.InfoCacheEntry
type JSONValue = packagejson.JSONValue
type JSONValueType = packagejson.JSONValueType
type PackageJson = packagejson.PackageJson
type PathFields = packagejson.PathFields
type TypeValidatedField = packagejson.TypeValidatedField
type VersionPaths = packagejson.VersionPaths

// Functions (exported as variables)
var NewInfoCache = packagejson.NewInfoCache
var Parse = packagejson.Parse

// Constants
const JSONValueTypeArray = packagejson.JSONValueTypeArray
const JSONValueTypeBoolean = packagejson.JSONValueTypeBoolean
const JSONValueTypeNotPresent = packagejson.JSONValueTypeNotPresent
const JSONValueTypeNull = packagejson.JSONValueTypeNull
const JSONValueTypeNumber = packagejson.JSONValueTypeNumber
const JSONValueTypeObject = packagejson.JSONValueTypeObject
const JSONValueTypeString = packagejson.JSONValueTypeString

