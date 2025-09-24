package bold

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-Bold.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "bold",
		File:   "Inter-Bold.woff2",
		Bytes:  data,
		Style:  "normal",
		Weight: "700",
	})
}

func Bytes() []byte { return data }
