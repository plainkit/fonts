package fonts_test

import (
	"testing"

	"github.com/plainkit/fonts"
	"github.com/plainkit/html"
)

func TestPreloadDefaults(t *testing.T) {
	link := fonts.Preload("/fonts/inter.woff2")
	got := html.Render(link)
	want := `<link as="font" href="/fonts/inter.woff2" rel="preload" type="font/woff2" crossorigin="anonymous">`
	if got != want {
		t.Fatalf("unexpected render\nwant: %s\n got: %s", want, got)
	}
}

func TestPreloadOverrides(t *testing.T) {
	link := fonts.Preload(
		"/fonts/inter.woff2",
		html.LinkRel("stylesheet"),
		fonts.FetchPriority("high"),
		html.LinkMedia("screen"),
	)

	got := html.Render(link)
	want := `<link as="font" fetchpriority="high" href="/fonts/inter.woff2" rel="stylesheet" type="font/woff2" media="screen" crossorigin="anonymous">`
	if got != want {
		t.Fatalf("unexpected render\nwant: %s\n got: %s", want, got)
	}
}

func TestNoCrossorigin(t *testing.T) {
	link := fonts.Preload("/fonts/inter.woff2", fonts.NoCrossorigin())
	got := html.Render(link)
	want := `<link as="font" href="/fonts/inter.woff2" rel="preload" type="font/woff2">`
	if got != want {
		t.Fatalf("unexpected render\nwant: %s\n got: %s", want, got)
	}
}
