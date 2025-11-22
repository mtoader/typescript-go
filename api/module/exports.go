package module

import "github.com/microsoft/typescript-go/internal/module"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type LookupLocations = module.LookupLocations
type ModeAwareCacheKey = module.ModeAwareCacheKey
type NodeResolutionFeatures = module.NodeResolutionFeatures
type PackageId = module.PackageId
type ParsedPatterns = module.ParsedPatterns
type ResolutionHost = module.ResolutionHost
type ResolvedModule = module.ResolvedModule
type ResolvedProjectReference = module.ResolvedProjectReference
type ResolvedTypeReferenceDirective = module.ResolvedTypeReferenceDirective
type Resolver = module.Resolver

// Functions (exported as variables)
var ComparePatternKeys = module.ComparePatternKeys
var GetAutomaticTypeDirectiveNames = module.GetAutomaticTypeDirectiveNames
var GetCompilerOptionsWithRedirect = module.GetCompilerOptionsWithRedirect
var GetConditions = module.GetConditions
var GetResolutionDiagnostic = module.GetResolutionDiagnostic
var GetTypesPackageName = module.GetTypesPackageName
var MangleScopedPackageName = module.MangleScopedPackageName
var MatchPatternOrExact = module.MatchPatternOrExact
var NewResolver = module.NewResolver
var ParseNodeModuleFromPath = module.ParseNodeModuleFromPath
var ParsePackageName = module.ParsePackageName
var ResolveConfig = module.ResolveConfig
var TryParsePatterns = module.TryParsePatterns
var UnmangleScopedPackageName = module.UnmangleScopedPackageName

// Constants
const InferredTypesContainingFile = module.InferredTypesContainingFile
const NodeResolutionFeaturesAll = module.NodeResolutionFeaturesAll
const NodeResolutionFeaturesBundlerDefault = module.NodeResolutionFeaturesBundlerDefault
const NodeResolutionFeaturesExports = module.NodeResolutionFeaturesExports
const NodeResolutionFeaturesExportsPatternTrailers = module.NodeResolutionFeaturesExportsPatternTrailers
const NodeResolutionFeaturesImports = module.NodeResolutionFeaturesImports
const NodeResolutionFeaturesNode16Default = module.NodeResolutionFeaturesNode16Default
const NodeResolutionFeaturesNodeNextDefault = module.NodeResolutionFeaturesNodeNextDefault
const NodeResolutionFeaturesNone = module.NodeResolutionFeaturesNone
const NodeResolutionFeaturesSelfName = module.NodeResolutionFeaturesSelfName

