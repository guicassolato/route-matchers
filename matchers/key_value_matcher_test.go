package matchers

import (
	"testing"

	"gotest.tools/assert"

	"github.com/guicassolato/route-matchers/utils"
)

func TestKeyValueMatcherContains(t *testing.T) {
	matcher := KeyValueMatcher{}
	assert.Check(t, matcher.Contains("foo=x", "foo=x"))
	assert.Check(t, !matcher.Contains("foo=x", "foo=y"))
	assert.Check(t, matcher.Contains("foo=*", "foo=y"))
	assert.Check(t, !matcher.Contains("foo=*", "bar=y"))

	matcher = KeyValueMatcher{Separator: ':'}
	assert.Check(t, matcher.Contains("foo:x", "foo:x"))
	assert.Check(t, !matcher.Contains("foo:x", "foo:y"))
	assert.Check(t, matcher.Contains("foo:*", "foo:y"))
	assert.Check(t, !matcher.Contains("foo:*", "bar:y"))
}

func TestKeyValueMatcherUnion(t *testing.T) {
	matcher := KeyValueMatcher{}

	union := matcher.Union(
		[]string{"foo=x", "bar=*", "baz=y"},
		[]string{"foo=x", "bar=w", "baz=z"},
	)

	assert.Equal(t, len(union), 4)
	assert.Check(t, utils.SliceContains(union, "foo=x"))
	assert.Check(t, utils.SliceContains(union, "bar=*"))
	assert.Check(t, utils.SliceContains(union, "baz=y"))
	assert.Check(t, utils.SliceContains(union, "baz=z"))
}

func TestKeyValueMatcherIntersection(t *testing.T) {
	matcher := KeyValueMatcher{}

	intersection := matcher.Intersection(
		[]string{"foo=x", "bar=*", "baz=y"},
		[]string{"foo=x", "bar=w", "baz=z"},
	)

	t.Log("intersection:", intersection)

	assert.Equal(t, len(intersection), 2)
	assert.Check(t, utils.SliceContains(intersection, "foo=x"))
	assert.Check(t, utils.SliceContains(intersection, "bar=w"))
}
