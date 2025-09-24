# Fonts for Plain

Embedded, self-hosted web fonts for the Plain component library.

## Installation

```bash
go get github.com/plainkit/html
go get github.com/plainkit/fonts
```

## Available Fonts

### Inter
- Variants: regular, italic, medium, medium italic, semibold, semibold italic, bold, bold italic, extra bold, extra bold italic (all WOFF2)
- Variant modules live under `github.com/plainkit/fonts/inter/assets/...` – import only the weights you need to keep binaries small
- Convenience sets:
  - `_ "github.com/plainkit/fonts/inter/basic"` registers regular, italic, bold, bold italic
  - `_ "github.com/plainkit/fonts/inter/all"` registers every shipped variant
- `inter.HeadComponents(prefix, variants...)` – emit preload + `@font-face` blocks for registered variants
- `inter.PreloadVariant(variant, prefix)` returns `(html.Component, bool)` so you can append the link only when the variant is available; `inter.Preload(href)` is a convenience wrapper when you already know the URL
- `inter.Bytes(variant)` – embedded font bytes for serving
- `inter.StaticHandler()` / `inter.RegisterStatic(mux, prefix)` – serve registered assets

If you omit the `variants...` parameter, `HeadComponents` includes the core
regular + italic + bold + bold italic set, preloading the regular weight (assuming
those variants were registered).

## Registering Variants

```go
import (
    _ "github.com/plainkit/fonts/inter/basic" // or inter/all, or individual assets/...
)
```

Only the variants whose packages are imported (directly or via a convenience set)
are linked into your binary, so you can keep deliverables lean by importing just
what you need.

## Serving the Font Assets

```go
package main

import (
    "net/http"

    _ "github.com/plainkit/fonts/inter/basic"
    "github.com/plainkit/fonts/inter"
)

func main() {
    mux := http.NewServeMux()
    inter.RegisterStatic(mux, "/assets/fonts/")

    _ = http.ListenAndServe(":8080", mux)
}
```

This mounts `/assets/fonts/Inter-*.woff2` with long-lived caching headers for the
registered variants. If you need manual control, call `inter.StaticHandler()` and
wrap it with `http.StripPrefix`.

## Adding to `<head>` with PlainKit

```go
import (
    _ "github.com/plainkit/fonts/inter/basic"

    x "github.com/plainkit/html"
    "github.com/plainkit/fonts"
    "github.com/plainkit/fonts/inter"
)

func head() x.HeadComponent {
    return x.Head(
        inter.HeadComponents("/assets/fonts")..., // preload + @font-face for defaults
        x.HeadStyle(x.Text(`body { font-family: "Inter", sans-serif; }`)),
    )
}
```

To include only certain weights/styles:

```go
head := x.Head(
    inter.HeadComponents(
        "/assets/fonts",
        inter.Regular,
        inter.Medium,
        inter.SemiBoldItalic,
        inter.ExtraBold,
    )...,
)
```

## Tweaking `<link>` output

`fonts.Preload` appends additional `html.LinkArg` values after its defaults, so you
can override anything PlainKit-style:

- `fonts.FetchPriority("high")` – set `fetchpriority`
- `fonts.NoCrossorigin()` – drop the `crossorigin` attribute entirely
- any other `html.LinkArg` (for example `html.LinkRel("stylesheet")`)

## License

Inter is distributed under the SIL Open Font License. The original `OFL.txt`
from Google Fonts is included in `fonts/inter`.
