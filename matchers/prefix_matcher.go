package matchers

import (
	"fmt"
	"strings"

	"github.com/guicassolato/route-matchers/utils"
)

const DefaultLabelSeparator = '.'

type PrefixMatcher struct{
	Separator      rune
	ComplementSign rune
  Wildcard
}

// Contains returns true if a value has a given prefix.
// Prefixes and values are composed of one or more "labels", separated by a fixed rune character (default: '.').
// The last label of a prefix can be the special wildcard rune (default: '*').
// Wildcards match any sequence of remaining labels in the value.
func (m *PrefixMatcher) Contains(prefix, value string) bool {
	prefixParts := strings.Split(prefix, string(m.separator()))
	valueParts := strings.Split(value, string(m.separator()))
	if len(valueParts) < len(prefixParts) {
		return false
	}
	for i, p := range prefixParts {
		if p == string(m.wildcard()) {
			return true
		}
		if p != valueParts[i] {
			return false
		}
	}
	return prefixParts[len(prefixParts)-1] == valueParts[len(valueParts)-1]
}

func (m *PrefixMatcher) Union(a, b []string) []string {
	return utils.Union(a, b, m.Contains)
}

func (m *PrefixMatcher) Intersection(a, b []string) []string {
	return utils.Intersection(a, b, m.Contains)
}

func (m *PrefixMatcher) Zip(s []string) string {
	return strings.Join(s, string(DefaultSetSeparator))
}

func (m *PrefixMatcher) Complement(s string) string {
	return fmt.Sprintf("%s%s", string(m.complementSign()), s)
}

func (m *PrefixMatcher) separator() rune {
	if m.Separator != 0 {
		return m.Separator
	}
	return DefaultLabelSeparator
}

func (m *PrefixMatcher) complementSign() rune {
	if m.ComplementSign != 0 {
		return m.ComplementSign
	}
	return DefaultComplementSign
}
