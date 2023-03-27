package matchers

import (
	"fmt"
	"strings"

	"github.com/guicassolato/route-matchers/utils"
)

const DefaultKeyValueSeparator = '='

type KeyValueMatcher struct{
	Separator      rune
	ComplementSign rune
	Wildcard
}

// Contains returns true if two key=value pairs have identical keys and either the values of the two pairs are identical
// or the value of the first pair is the special wildcard rune (default: '*').
func (m *KeyValueMatcher) Contains(kv1, kv2 string) bool {
	key1, value1 := m.kv(kv1)
	key2, value2 := m.kv(kv2)
	return m.contains(key1, value1, key2, value2)
}

func (m *KeyValueMatcher) Union(a, b []string) []string {
	return utils.Union(a, b, m.Contains)
}

// Intersection returns the intersection between two sets A and B, given a provided 'contains' function.
// E.g.:
//   A = { foo=x, bar=*, baz=y }
//   B = { foo=x, bar=w, baz=z }
// Intersection(A, B) → A ∩ B = { foo=x, bar=w }
func (m *KeyValueMatcher) Intersection(a, b []string) []string {
	return utils.Intersection(a, b, m.Contains)
}

func (m *KeyValueMatcher) Zip(s []string) string {
	return strings.Join(s, string(DefaultSetSeparator))
}

func (m *KeyValueMatcher) Complement(kv string) string {
	key, value := m.kv(kv)
	return fmt.Sprintf("%s%s%s%s", key, string(m.separator()), string(m.complementSign()), value)
}

func (m *KeyValueMatcher) contains(key1, value1, key2, value2 string) bool {
	return key1 == key2 && (value1 == string(m.wildcard()) || value1 == value2)
}

func (m *KeyValueMatcher) separator() rune {
	if m.Separator != 0 {
		return m.Separator
	}
	return DefaultKeyValueSeparator
}

func (m *KeyValueMatcher) complementSign() rune {
	if m.ComplementSign != 0 {
		return m.ComplementSign
	}
	return DefaultComplementSign
}

func (m *KeyValueMatcher) kv(s string) (string, string) {
	parts := strings.Split(s, string(m.separator()))
	val := string(m.wildcard())
	if len(parts) > 1 {
		val = parts[1]
	}
	return parts[0], val
}
