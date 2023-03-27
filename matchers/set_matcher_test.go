package matchers

import (
	"testing"

	"gotest.tools/assert"

	"github.com/guicassolato/route-matchers/utils"
)

func TestSetMatcherContains(t *testing.T) {
	matcher := SetMatcher{}
  assert.Check(t, matcher.Contains("", ""))
  assert.Check(t, matcher.Contains("a", "a"))
  assert.Check(t, !matcher.Contains("a", "b"))
  assert.Check(t, matcher.Contains("a,b", "a"))
  assert.Check(t, matcher.Contains("a,b", "b"))
  assert.Check(t, matcher.Contains("a,b", "a,b"))
  assert.Check(t, !matcher.Contains("a,b", "c"))
  assert.Check(t, !matcher.Contains("a", "a,b"))

	matcher = SetMatcher{
		Matcher:   &KeyValueMatcher{ Separator: '=' },
		Separator: '&',
	}
	assert.Check(t, matcher.Contains("", ""))
	assert.Check(t, matcher.Contains("foo=x", "foo=x"))
	assert.Check(t, matcher.Contains("foo=x&bar=*", "foo=x"))
	assert.Check(t, matcher.Contains("foo=x&bar=*", "bar=y"))
	assert.Check(t, matcher.Contains("foo=x&bar=*", "foo=x&bar=y"))
	assert.Check(t, !matcher.Contains("foo=x", "foo=*"))
	assert.Check(t, !matcher.Contains("foo=x", "foo=x&bar=y"))
	assert.Check(t, !matcher.Contains("", "foo=x"))

	matcher = SetMatcher{
		Matcher:   &KeyValueMatcher{ Separator: ':' },
		Separator: ',',
	}
	assert.Check(t, matcher.Contains("content-type:application/json,authorization:*", "content-type:application/json"))
	assert.Check(t, matcher.Contains("content-type:application/json,authorization:*", "authorization:Bearer xyz"))
	assert.Check(t, matcher.Contains("content-type:application/json,authorization:*", "content-type:application/json,authorization:Bearer xyz"))
}

func TestSetMatcherUnion(t *testing.T) {
	matcher := SetMatcher{}

	union := matcher.Union(
		[]string{"a,b", "c", "e"},
		[]string{"a,b", "a", "b", "b,c", "d"},
	)

	assert.Equal(t, len(union), 5)
	assert.Check(t, utils.SliceContains(union, "a"))
	assert.Check(t, utils.SliceContains(union, "b"))
	assert.Check(t, utils.SliceContains(union, "c"))
	assert.Check(t, utils.SliceContains(union, "d"))
	assert.Check(t, utils.SliceContains(union, "e"))
}

func TestSetMatcherIntersection(t *testing.T) {
	matcher := SetMatcher{}

	intersection := matcher.Intersection(
		[]string{"a,b", "c", "e"},
		[]string{"a,b", "a", "b", "b,c", "d"},
	)

	assert.Equal(t, len(intersection), 3)
	assert.Check(t, utils.SliceContains(intersection, "a"))
	assert.Check(t, utils.SliceContains(intersection, "b"))
	assert.Check(t, utils.SliceContains(intersection, "c"))

	matcher = SetMatcher{
		Matcher:   &KeyValueMatcher{ Separator: '=' },
		Separator: '&',
	}

	intersection = matcher.Intersection(
		[]string{"foo=x", "foo=x&bar=*"},
		[]string{"foo=x", "bar=w", "baz=z", "bar=v&baz=z"},
	)

	assert.Equal(t, len(intersection), 3)
	assert.Check(t, utils.SliceContains(intersection, "foo=x"))
	assert.Check(t, utils.SliceContains(intersection, "bar=w"))
	assert.Check(t, utils.SliceContains(intersection, "bar=v"))
}
