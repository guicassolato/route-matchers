package matchers

import (
	"testing"

	"gotest.tools/assert"
)

func TestTable(t *testing.T) {
	table := NewTable()
	table.Add("foo")
	assert.Equal(t, table.Size(), 2)
}

func TestStringTable(t *testing.T) {
	table := NewTable(WithMatcher(&StringMatcher{}))
	table.Add("foo")
	table.Add("bar")
	assert.Equal(t, table.Size(), 4)

	items := table.List()
	assert.Equal(t, items[0], "foo ∩ bar = {}")
	assert.Equal(t, items[1], "^foo ∩ bar = {}") // FIXME: it should be {bar}
	assert.Equal(t, items[2], "foo ∩ ^bar = {}") // FIXME: it should be {foo}
	assert.Equal(t, items[3], "^foo ∩ ^bar = {}")
}

func TestPrefixTable(t *testing.T) {
	table := NewTable(WithMatcher(&PrefixMatcher{}))
	table.Add("com.acme.*")
	table.Add("com.acme.toys")
	assert.Equal(t, table.Size(), 4)

	items := table.List()
	assert.Equal(t, items[0], "com.acme.* ∩ com.acme.toys = {com.acme.toys}")
	assert.Equal(t, items[1], "^com.acme.* ∩ com.acme.toys = {}")
	assert.Equal(t, items[2], "com.acme.* ∩ ^com.acme.toys = {}") // FIXME: it should be {com.acme.* - com.acme.toys}
	assert.Equal(t, items[3], "^com.acme.* ∩ ^com.acme.toys = {^com.acme.toys}") // FIXME: it should be {^com.acme.*}

	// table.Add("com.acme.*")
	// table.Add("com.acme.toys")
	// table.Add("internal.acme.toys")
	// table.Add("internal.acme.*")
	// table.Add("internal.acme.telemetry.*")
	// table.Add("internal.acme.telemetry.foo")
	// assert.Equal(t, table.Size(), 64)

	// items := table.List()
	// assert.Equal(t, items[0], "com.acme.* ∩ com.acme.toys ∩ internal.acme.toys ∩ internal.acme.* ∩ internal.acme.telemetry.* ∩ internal.acme.telemetry.foo = {}")
}
