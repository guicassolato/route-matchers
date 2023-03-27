package matchers

import (
	"strings"

	"github.com/guicassolato/route-matchers/utils"
)

type SetMatcher struct {
	Separator rune
	Matcher
}

func (m *SetMatcher) Contains(supersets, subsets string) bool {
  superset := m.splitter()(supersets)
	for _, subset := range m.splitter()(subsets) {
		found := false
		for _, s := range superset {
			if m.matcher().Contains(s, subset) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (m *SetMatcher) Union(a, b []string) []string {
	return utils.Union(utils.FlattenSlice(a, m.splitter()), utils.FlattenSlice(b, m.splitter()), m.Contains)
}

func (m *SetMatcher) Intersection(a, b []string) []string {
	return utils.Intersection(m.Union(a, nil), m.Union(b, nil), m.Contains)
}

func (m *SetMatcher) separator() rune {
	if m.Separator != 0 {
		return m.Separator
	}
	return DefaultSetSeparator
}

func (m *SetMatcher) Zip(s []string) string {
	return strings.Join(s, string(m.separator()))
}

func (m *SetMatcher) Complement(s string) string {
	values := utils.MapSlice(m.splitter()(s), func(value string) string {
		return m.matcher().Complement(value)
	})
	return m.Zip(values)
}

func (m *SetMatcher) splitter() func(string) []string {
	return func(s string) []string {
		return strings.Split(s, string(m.separator()))
	}
}

func (m *SetMatcher) matcher() Matcher {
	if m.Matcher != nil {
		return m.Matcher
	}
	return DefaultMatcher
}
