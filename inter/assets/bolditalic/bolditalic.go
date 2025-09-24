package bolditalic

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-BoldItalic.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "bold-italic",
		File:   "Inter-BoldItalic.woff2",
		Bytes:  data,
		Style:  "italic",
		Weight: "700",
	})
}

func Bytes() []byte { return data }
