package regular

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-Regular.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "regular",
		File:   "Inter-Regular.woff2",
		Bytes:  data,
		Style:  "normal",
		Weight: "400",
	})
}

// Bytes returns the Inter Regular WOFF2 bytes.
func Bytes() []byte { return data }
