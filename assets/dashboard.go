package assets

import (
	"embed"
	_ "embed"
)

//go:embed .next
var NextFS embed.FS
