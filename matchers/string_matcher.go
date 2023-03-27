package matchers

import (
	"fmt"
	"strings"

	"github.com/guicassolato/route-matchers/utils"
)

type StringMatcher struct {
	ComplementSign rune
}

func (m *StringMatcher) Contains(v1, v2 string) bool {
	return v1 == v2
}

func (m *StringMatcher) Union(a, b []string) []string {
	return utils.Union(a, b, m.Contains)
}

func (m *StringMatcher) Intersection(a, b []string) []string {
	return utils.Intersection(a, b, m.Contains)
}

func (m *StringMatcher) Zip(s []string) string {
	return strings.Join(s, string(DefaultSetSeparator))
}

func (m *StringMatcher) Complement(s string) string {
	return fmt.Sprintf("%s%s", string(m.complementSign()), s)
}

func (m *StringMatcher) complementSign() rune {
	if m.ComplementSign != 0 {
		return m.ComplementSign
	}
	return DefaultComplementSign
}
