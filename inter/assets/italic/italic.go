package italic

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-Italic.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "italic",
		File:   "Inter-Italic.woff2",
		Bytes:  data,
		Style:  "italic",
		Weight: "400",
	})
}

func Bytes() []byte { return data }
