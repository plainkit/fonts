package extrabolditalic

import _ "embed"

import "github.com/plainkit/fonts/inter/assets"

//go:embed Inter-ExtraBoldItalic.woff2
var data []byte

func init() {
	assets.Register(assets.VariantAsset{
		Name:   "extra-bold-italic",
		File:   "Inter-ExtraBoldItalic.woff2",
		Bytes:  data,
		Style:  "italic",
		Weight: "800",
	})
}

func Bytes() []byte { return data }
