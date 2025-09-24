package inter

import (
	"bytes"
	"net/http"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/plainkit/fonts"
	"github.com/plainkit/fonts/inter/assets"
	"github.com/plainkit/html"
)

// Variant represents a specific Inter weight/style combination.
type Variant string

const (
	// Family is the public font-family name for Inter.
	Family = "Inter"

	Regular         Variant = "regular"
	Italic          Variant = "italic"
	Medium          Variant = "medium"
	MediumItalic    Variant = "medium-italic"
	SemiBold        Variant = "semi-bold"
	SemiBoldItalic  Variant = "semi-bold-italic"
	Bold            Variant = "bold"
	BoldItalic      Variant = "bold-italic"
	ExtraBold       Variant = "extra-bold"
	ExtraBoldItalic Variant = "extra-bold-italic"
)

var defaultVariants = []Variant{Regular, Italic, Bold, BoldItalic}

var modTime = time.Unix(0, 0)

// Available lists all registered variants. Blank-import the variant packages
// you need (for example, `_ "github.com/plainkit/fonts/inter/basic"`).
func Available() []Variant {
	assetsList := assets.All()
	sort.Slice(assetsList, func(i, j int) bool { return assetsList[i].Name < assetsList[j].Name })
	variants := make([]Variant, 0, len(assetsList))
	for _, asset := range assetsList {
		variants = append(variants, Variant(asset.Name))
	}
	return variants
}

// Bytes returns the embedded font data for the requested variant if it has been
// registered via a variant subpackage.
func Bytes(v Variant) ([]byte, bool) {
	a, ok := assets.Get(string(v))
	if !ok {
		return nil, false
	}
	return a.Bytes, true
}

// File returns the filename for the given variant if registered.
func File(v Variant) (string, bool) {
	a, ok := assets.Get(string(v))
	if !ok {
		return "", false
	}
	return a.File, true
}

// Preload emits a preload <link> for the given href. Variants must be
// registered separately; this helper simply applies sensible defaults.
func Preload(href string, extras ...html.LinkArg) html.LinkComponent {
	return fonts.Preload(href, extras...)
}

// PreloadVariant emits a preload <link> element for a specific variant if it is
// registered. The second return value reports whether a link was produced.
func PreloadVariant(v Variant, prefix string, extras ...html.LinkArg) (html.HeadArg, bool) {
	href, ok := hrefFor(prefix, v)
	if !ok {
		return nil, false
	}
	return fonts.Preload(href, extras...), true
}

// HeadComponents returns recommended components for loading the provided
// variants. If no variants are supplied, Regular, Italic, Bold, and BoldItalic
// are used. Only variants that have been registered via subpackage imports are
// emitted. Regular is preloaded when present.
func HeadComponents(prefix string, variants ...Variant) []html.HeadArg {
	if len(variants) == 0 {
		variants = defaultVariants
	}

	prefix = strings.TrimSuffix(prefix, "/")

	components := make([]html.HeadArg, 0, len(variants)+1)

	if containsVariant(variants, Regular) {
		if preload, ok := PreloadVariant(Regular, prefix, fonts.FetchPriority("high")); ok {
			components = append(components, preload)
		}
	}

	for _, v := range variants {
		if style, ok := fontFaceStyle(prefix, v); ok {
			components = append(components, style)
		}
	}

	return components
}

func fontFaceStyle(prefix string, v Variant) (html.HeadArg, bool) {
	a, ok := assets.Get(string(v))
	if !ok {
		return nil, false
	}

	href, _ := hrefFor(prefix, v)

	css := `@font-face {
	font-family: '` + Family + `';
	font-style: ` + a.Style + `;
	font-weight: ` + a.Weight + `;
	font-display: swap;
	src: url('` + href + `') format('woff2');
}
`

	return html.HeadStyle(html.UnsafeText(css)), true
}

func containsVariant(list []Variant, v Variant) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func hrefFor(prefix string, v Variant) (string, bool) {
	a, ok := assets.Get(string(v))
	if !ok {
		return "", false
	}
	if prefix == "" {
		return a.File, true
	}
	return prefix + "/" + a.File, true
}

// StaticHandler serves registered Inter font files. Mount with http.StripPrefix
// when serving from a sub-path.
func StaticHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		for _, asset := range sortedAssets() {
			if name == asset.File {
				serveFont(w, r, asset)
				return
			}
		}
		http.NotFound(w, r)
	})
}

// RegisterStatic mounts StaticHandler on an http.ServeMux using the provided
// prefix. The prefix should include a trailing slash, e.g. "/assets/fonts/".
func RegisterStatic(mux *http.ServeMux, prefix string) {
	if mux == nil {
		return
	}
	if prefix == "" {
		prefix = "/"
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	mux.Handle(prefix, http.StripPrefix(prefix, StaticHandler()))
}

func serveFont(w http.ResponseWriter, r *http.Request, asset assets.VariantAsset) {
	w.Header().Set("Content-Type", fonts.MIMETypeWOFF2)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	http.ServeContent(w, r, asset.File, modTime, bytes.NewReader(asset.Bytes))
}

func sortedAssets() []assets.VariantAsset {
	list := assets.All()
	sort.Slice(list, func(i, j int) bool { return list[i].File < list[j].File })
	return list
}
