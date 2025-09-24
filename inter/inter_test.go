package inter_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/plainkit/fonts/inter/all"

	"github.com/plainkit/fonts"
	"github.com/plainkit/fonts/inter"
	"github.com/plainkit/html"
)

func TestHeadComponentsDefault(t *testing.T) {
	components := inter.HeadComponents("/assets/fonts")
	if len(components) != 5 {
		t.Fatalf("expected 5 components (1 preload + 4 font-face), got %d", len(components))
	}

	link := html.Render(components[0])
	expectedLink := `<link as="font" fetchpriority="high" href="/assets/fonts/Inter-Regular.woff2" rel="preload" type="font/woff2" crossorigin="anonymous">`
	if link != expectedLink {
		t.Fatalf("unexpected preload link\nwant: %s\n got: %s", expectedLink, link)
	}

	checks := []struct {
		index int
		must  []string
	}{
		{index: 1, must: []string{"font-style: normal;", "font-weight: 400;", "url('/assets/fonts/Inter-Regular.woff2')"}},
		{index: 2, must: []string{"font-style: italic;", "font-weight: 400;", "url('/assets/fonts/Inter-Italic.woff2')"}},
		{index: 3, must: []string{"font-style: normal;", "font-weight: 700;", "url('/assets/fonts/Inter-Bold.woff2')"}},
		{index: 4, must: []string{"font-style: italic;", "font-weight: 700;", "url('/assets/fonts/Inter-BoldItalic.woff2')"}},
	}

	for _, check := range checks {
		rendered := html.Render(components[check.index])
		for _, substr := range check.must {
			if !strings.Contains(rendered, substr) {
				t.Fatalf("component %d missing %q in %s", check.index, substr, rendered)
			}
		}
	}
}

func TestHeadComponentsExtendedVariants(t *testing.T) {
	components := inter.HeadComponents("/assets/fonts", inter.Regular, inter.Medium, inter.SemiBoldItalic, inter.ExtraBold)
	if len(components) != 5 {
		t.Fatalf("expected 5 components, got %d", len(components))
	}

	if got := html.Render(components[0]); !strings.Contains(got, `href="/assets/fonts/Inter-Regular.woff2"`) {
		t.Fatalf("preload missing regular href: %s", got)
	}

	if got := html.Render(components[1]); !strings.Contains(got, "font-weight: 400;") {
		t.Fatalf("regular face missing weight: %s", got)
	}

	if got := html.Render(components[2]); !strings.Contains(got, "font-weight: 500;") {
		t.Fatalf("medium face missing weight: %s", got)
	}

	if got := html.Render(components[3]); !strings.Contains(got, "font-weight: 600;") || !strings.Contains(got, "font-style: italic;") {
		t.Fatalf("semi-bold italic face missing attributes: %s", got)
	}

	if got := html.Render(components[4]); !strings.Contains(got, "font-weight: 800;") {
		t.Fatalf("extra-bold face missing weight: %s", got)
	}
}

func TestHeadComponentsWithoutRegular(t *testing.T) {
	components := inter.HeadComponents("/assets/fonts", inter.Medium, inter.ExtraBoldItalic)
	if len(components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(components))
	}

	if got := html.Render(components[0]); strings.Contains(got, "rel=\"preload\"") {
		t.Fatalf("unexpected preload when regular not requested: %s", got)
	}
}

func TestStaticHandlerServesEachVariant(t *testing.T) {
	variants := []struct {
		name    string
		variant inter.Variant
	}{
		{"Inter-Regular.woff2", inter.Regular},
		{"Inter-Italic.woff2", inter.Italic},
		{"Inter-Medium.woff2", inter.Medium},
		{"Inter-MediumItalic.woff2", inter.MediumItalic},
		{"Inter-SemiBold.woff2", inter.SemiBold},
		{"Inter-SemiBoldItalic.woff2", inter.SemiBoldItalic},
		{"Inter-Bold.woff2", inter.Bold},
		{"Inter-BoldItalic.woff2", inter.BoldItalic},
		{"Inter-ExtraBold.woff2", inter.ExtraBold},
		{"Inter-ExtraBoldItalic.woff2", inter.ExtraBoldItalic},
	}

	handler := inter.StaticHandler()

	for _, v := range variants {
		v := v
		t.Run(v.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/"+v.name, nil)
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			res := rec.Result()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("expected status 200, got %d", res.StatusCode)
			}
			if ct := res.Header.Get("Content-Type"); ct != fonts.MIMETypeWOFF2 {
				t.Fatalf("unexpected content type: %s", ct)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			expected := mustBytes(v.variant)
			if len(body) != len(expected) {
				t.Fatalf("body length mismatch: want %d got %d", len(expected), len(body))
			}
			for i := range body {
				if body[i] != expected[i] {
					t.Fatal("body bytes do not match embedded font")
				}
			}
		})
	}
}

func mustBytes(v inter.Variant) []byte {
	data, ok := inter.Bytes(v)
	if !ok {
		panic("missing variant")
	}
	return data
}
