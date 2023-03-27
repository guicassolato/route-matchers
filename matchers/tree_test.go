package matchers

import (
	"testing"

	"gotest.tools/assert"
)

func TestTree(t *testing.T) {
	tree := NewTree()
	tree.Add("foo")
	tree.Add("bar")
	tree.Add("bar")
	assert.Equal(t, tree.Size(), 2)
}

func TestStringTree(t *testing.T) {
	tree := NewTree(WithMatcher(&StringMatcher{}))
	tree.Add("foo")
	tree.Add("bar")
	tree.Add("bar")

	items := tree.List()
	assert.Equal(t, len(items), 2)
  assert.Equal(t, items[0], "﹂foo")
  assert.Equal(t, items[1], "﹂bar")
}

func TestPrefixTree(t *testing.T) {
	tree := NewTree(WithMatcher(&PrefixMatcher{}))
	tree.Add("com.acme.*")
	tree.Add("com.acme.toys")
	tree.Add("internal.acme.toys")
	tree.Add("internal.acme.*")
	tree.Add("internal.acme.telemetry.*")
	tree.Add("internal.acme.telemetry.foo")

	items := tree.List()
  assert.Equal(t, len(items), 6)
  assert.Equal(t, items[0], "﹂com.acme.*")
  assert.Equal(t, items[1], "  ﹂com.acme.toys")
  assert.Equal(t, items[2], "﹂internal.acme.*")
  assert.Equal(t, items[3], "  ﹂internal.acme.toys")
  assert.Equal(t, items[4], "  ﹂internal.acme.telemetry.*")
  assert.Equal(t, items[5], "    ﹂internal.acme.telemetry.foo")

 	tree = NewTree(WithMatcher(&PrefixMatcher{Separator: '/'}))
	tree.Add("/*")
	tree.Add("/admin")
	tree.Add("/admin/*")
	tree.Add("/orgs/*")
	tree.Add("/orgs/123")

	items = tree.List()
  assert.Equal(t, len(items), 5)
  assert.Equal(t, items[0], "﹂/*")
  assert.Equal(t, items[1], "  ﹂/admin")
  assert.Equal(t, items[2], "  ﹂/admin/*")
  assert.Equal(t, items[3], "  ﹂/orgs/*")
  assert.Equal(t, items[4], "    ﹂/orgs/123")
}

func TestKeyValueTree(t *testing.T) {
	tree := NewTree(WithMatcher(&KeyValueMatcher{}))
	tree.Add("foo=x")
	tree.Add("bar=*")
	tree.Add("bar=y")
	tree.Add("baz=z")

	items := tree.List()
  assert.Equal(t, len(items), 4)
  assert.Equal(t, items[0], "﹂foo=x")
  assert.Equal(t, items[1], "﹂bar=*")
  assert.Equal(t, items[2], "  ﹂bar=y")
  assert.Equal(t, items[3], "﹂baz=z")
}

func TestSetTree(t *testing.T) {
	tree := NewTree(WithMatcher(&SetMatcher{}))
	tree.Add("a")
	tree.Add("b,c,d")
	tree.Add("c")
	tree.Add("c,d")
	tree.Add("d")
	tree.Add("d,e")
	tree.Add("e")

	items := tree.List()
  assert.Equal(t, len(items), 7)
  assert.Equal(t, items[0], "﹂a")
  assert.Equal(t, items[1], "﹂b,c,d")
  assert.Equal(t, items[2], "  ﹂c,d")
  assert.Equal(t, items[3], "    ﹂c")
  assert.Equal(t, items[4], "    ﹂d")
  assert.Equal(t, items[5], "﹂d,e")
  assert.Equal(t, items[6], "  ﹂e")
}
