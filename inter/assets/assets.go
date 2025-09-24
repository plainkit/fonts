package assets

import _ "embed"

// VariantAsset exposes a named WOFF2 font asset.
type VariantAsset struct {
	Name string
	File string
	// Bytes contains the embedded font file.
	Bytes  []byte
	Style  string
	Weight string
}

var registry = map[string]VariantAsset{}

// Register adds a VariantAsset to the registry.
func Register(asset VariantAsset) {
	registry[asset.Name] = asset
}

// Get returns a registered VariantAsset by name.
func Get(name string) (VariantAsset, bool) {
	a, ok := registry[name]
	return a, ok
}

// All returns every registered VariantAsset.
func All() []VariantAsset {
	items := make([]VariantAsset, 0, len(registry))
	for _, asset := range registry {
		items = append(items, asset)
	}
	return items
}
