package change

import "github.com/microsoft/typescript-go/internal/ls/change"

// This file re-exports all public types, functions, constants, and variables from the internal package.

// Types
type LeadingTriviaOption = change.LeadingTriviaOption
type NodeOptions = change.NodeOptions
type Tracker = change.Tracker
type TrailingTriviaOption = change.TrailingTriviaOption

// Functions (exported as variables)
var NewTracker = change.NewTracker

// Constants
const LeadingTriviaOptionExclude = change.LeadingTriviaOptionExclude
const LeadingTriviaOptionIncludeAll = change.LeadingTriviaOptionIncludeAll
const LeadingTriviaOptionJSDoc = change.LeadingTriviaOptionJSDoc
const LeadingTriviaOptionNone = change.LeadingTriviaOptionNone
const LeadingTriviaOptionStartLine = change.LeadingTriviaOptionStartLine
const TrailingTriviaOptionExclude = change.TrailingTriviaOptionExclude
const TrailingTriviaOptionExcludeWhitespace = change.TrailingTriviaOptionExcludeWhitespace
const TrailingTriviaOptionInclude = change.TrailingTriviaOptionInclude
const TrailingTriviaOptionNone = change.TrailingTriviaOptionNone

