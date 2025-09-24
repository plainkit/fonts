package extrabold

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-ExtraBold.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "extra-bold",
		File:   "Inter-ExtraBold.woff2",
		Bytes:  data,
		Style:  "normal",
		Weight: "800",
	})
}

func Bytes() []byte { return data }
