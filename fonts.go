package fonts

import "github.com/plainkit/html"

// MIME types for web fonts.
const (
	MIMETypeWOFF  = "font/woff"
	MIMETypeWOFF2 = "font/woff2"
)

// Preload emits a <link rel="preload"> element with sensible defaults for
// font delivery. Defaults: rel="preload", as="font", type="font/woff2",
// crossorigin="anonymous". Additional html.LinkArg values can override these
// attributes or add new ones following PlainKit conventions.
func Preload(href string, extras ...html.LinkArg) html.LinkComponent {
	args := []html.LinkArg{
		html.LinkHref(href),
		html.LinkRel("preload"),
		html.LinkType(MIMETypeWOFF2),
		html.Crossorigin("anonymous"),
		html.Custom("as", "font"),
	}

	args = append(args, extras...)

	return html.Link(args...)
}

// FetchPriority sets the fetchpriority attribute on a link element.
func FetchPriority(value string) html.LinkArg {
	return html.Custom("fetchpriority", value)
}

// NoCrossorigin clears the crossorigin attribute on the produced link element.
func NoCrossorigin() html.LinkArg {
	return html.Crossorigin("")
}
