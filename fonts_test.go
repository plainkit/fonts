package fonts_test

import (
	"strings"
	"testing"

	"github.com/plainkit/fonts"
	"github.com/plainkit/html"
)

func TestPreloadDefaults(t *testing.T) {
	link := fonts.Preload("/fonts/inter.woff2")
	got := html.Render(link)

	// Check for required attributes (order may vary)
	requiredAttrs := []string{
		`as="font"`,
		`href="/fonts/inter.woff2"`,
		`rel="preload"`,
		`type="font/woff2"`,
		`crossorigin="anonymous"`,
	}
	for _, attr := range requiredAttrs {
		if !strings.Contains(got, attr) {
			t.Fatalf("missing required attribute %s in: %s", attr, got)
		}
	}
}

func TestPreloadOverrides(t *testing.T) {
	link := fonts.Preload(
		"/fonts/inter.woff2",
		html.ARel("stylesheet"),
		fonts.FetchPriority("high"),
		html.AMedia("screen"),
	)

	got := html.Render(link)

	// Check for required attributes (order may vary)
	requiredAttrs := []string{
		`as="font"`,
		`fetchpriority="high"`,
		`href="/fonts/inter.woff2"`,
		`rel="preload stylesheet"`, // rel attributes combine
		`type="font/woff2"`,
		`media="screen"`,
		`crossorigin="anonymous"`,
	}
	for _, attr := range requiredAttrs {
		if !strings.Contains(got, attr) {
			t.Fatalf("missing required attribute %s in: %s", attr, got)
		}
	}
}

func TestNoCrossorigin(t *testing.T) {
	link := fonts.Preload("/fonts/inter.woff2", fonts.NoCrossorigin())
	got := html.Render(link)
	want := `<link as="font" href="/fonts/inter.woff2" rel="preload" type="font/woff2"/>`
	if got != want {
		t.Fatalf("unexpected render\nwant: %s\n got: %s", want, got)
	}
}
