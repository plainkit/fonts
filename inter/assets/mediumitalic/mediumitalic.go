package mediumitalic

import _ "embed"

import "github.com/plainkit/fonts/inter/assets"

//go:embed Inter-MediumItalic.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "medium-italic",
		File:   "Inter-MediumItalic.woff2",
		Bytes:  data,
		Style:  "italic",
		Weight: "500",
	})
}

func Bytes() []byte { return data }
