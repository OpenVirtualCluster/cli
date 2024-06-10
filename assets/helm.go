package assets

import (
	_ "embed"
)

//go:embed helm-linux
var HelmLinux []byte

//go:embed helm-darwin
var HelmDarwin []byte
