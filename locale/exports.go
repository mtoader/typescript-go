package locale

import (
	"github.com/microsoft/typescript-go/internal/locale"
)

type Locale = locale.Locale

var Default = locale.Default

var (
	WithLocale  = locale.WithLocale
	FromContext = locale.FromContext
	Parse       = locale.Parse
)
