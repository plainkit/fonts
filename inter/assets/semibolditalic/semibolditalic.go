package semibolditalic

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-SemiBoldItalic.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "semi-bold-italic",
		File:   "Inter-SemiBoldItalic.woff2",
		Bytes:  data,
		Style:  "italic",
		Weight: "600",
	})
}

func Bytes() []byte { return data }
