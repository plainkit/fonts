package medium

import (
	_ "embed"

	"github.com/plainkit/fonts/inter/assets"
)

//go:embed Inter-Medium.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "medium",
		File:   "Inter-Medium.woff2",
		Bytes:  data,
		Style:  "normal",
		Weight: "500",
	})
}

func Bytes() []byte { return data }
