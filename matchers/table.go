package matchers

import (
	"fmt"
	"strings"
)

func NewTable(options ...StoreOption) Store {
	t := &table{}
	for _, option := range options {
		option(t)
	}
	return t
}

type table struct {
	items [][]string
	Matcher
}

// impl:Store

func (t *table) Add(set string) {
	complement := t.matcher().Complement(set)

	if len(t.items) == 0 {
		t.items = append(t.items, []string{set}, []string{complement})
		return
	}

	var newItems [][]string
	for index := range t.items {
		newItems = append(newItems, append(t.items[index], complement))
		t.items[index] = append(t.items[index], set)
	}
	t.items = append(t.items, newItems...)
}

func (t *table) List() []string {
	var nonIntersectingSets []string
	for _, sets := range t.items {
    nonIntersectingSet := sets[0]
		for _, set := range sets[1:] {
			nonIntersectingSet = t.matcher().Zip(t.matcher().Intersection([]string{nonIntersectingSet}, []string{set}))
		}
		nonIntersectingSets = append(nonIntersectingSets, fmt.Sprintf("%s = {%s}", strings.Join(sets, " ∩ "), nonIntersectingSet))
	}
	return nonIntersectingSets
}

func (t *table) Size() int {
	return len(t.items)
}

func (t *table) SetMatcher(matcher Matcher) {
	t.Matcher = matcher
}

func (t *table) matcher() Matcher {
	if t.Matcher != nil {
		return t.Matcher
	}
	return DefaultMatcher
}
