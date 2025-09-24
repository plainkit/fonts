package semibold

import _ "embed"

import "github.com/plainkit/fonts/inter/assets"

//go:embed Inter-SemiBold.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "semi-bold",
		File:   "Inter-SemiBold.woff2",
		Bytes:  data,
		Style:  "normal",
		Weight: "600",
	})
}

func Bytes() []byte { return data }
