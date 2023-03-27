package matchers

import (
	"testing"

	"gotest.tools/assert"

	"github.com/guicassolato/route-matchers/utils"
)

func TestPrefixMatcherContains(t *testing.T) {
  matcher := PrefixMatcher{}

	hostnames := []string{
		"com.acme.*",
		"com.acme.toys",
		"internal.acme.*",
		"internal.acme.toys",
		"internal.acme.telemetry.*",
		"internal.acme.telemetry.foo",
	}

	// com.acme.*
	assert.Check(t, matcher.Contains(hostnames[0], hostnames[0]))  // com.acme.*
	assert.Check(t, matcher.Contains(hostnames[0], hostnames[1]))  // com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[0], hostnames[2])) // internal.acme.*
	assert.Check(t, !matcher.Contains(hostnames[0], hostnames[3])) // internal.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[0], hostnames[4])) // internal.acme.telemetry.*
	assert.Check(t, !matcher.Contains(hostnames[0], hostnames[5])) // internal.acme.telemetry.foo

	// com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[1], hostnames[0])) // com.acme.*
	assert.Check(t, matcher.Contains(hostnames[1], hostnames[1]))  // com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[1], hostnames[2])) // internal.acme.*
	assert.Check(t, !matcher.Contains(hostnames[1], hostnames[3])) // internal.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[1], hostnames[4])) // internal.acme.telemetry.*
	assert.Check(t, !matcher.Contains(hostnames[1], hostnames[5])) // internal.acme.telemetry.foo

	// internal.acme.*
	assert.Check(t, !matcher.Contains(hostnames[2], hostnames[0])) // com.acme.*
	assert.Check(t, !matcher.Contains(hostnames[2], hostnames[1])) // com.acme.toys
	assert.Check(t, matcher.Contains(hostnames[2], hostnames[2]))  // internal.acme.*
	assert.Check(t, matcher.Contains(hostnames[2], hostnames[3]))  // internal.acme.toys
	assert.Check(t, matcher.Contains(hostnames[2], hostnames[4]))  // internal.acme.telemetry.*
	assert.Check(t, matcher.Contains(hostnames[2], hostnames[5]))  // internal.acme.telemetry.foo

	// internal.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[3], hostnames[0])) // com.acme.*
	assert.Check(t, !matcher.Contains(hostnames[3], hostnames[1])) // com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[3], hostnames[2])) // internal.acme.*
	assert.Check(t, matcher.Contains(hostnames[3], hostnames[3]))  // internal.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[3], hostnames[4])) // internal.acme.telemetry.*
	assert.Check(t, !matcher.Contains(hostnames[3], hostnames[5])) // internal.acme.telemetry.foo

	// // internal.acme.telemetry.*
	assert.Check(t, !matcher.Contains(hostnames[4], hostnames[0])) // com.acme.*
	assert.Check(t, !matcher.Contains(hostnames[4], hostnames[1])) // com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[4], hostnames[2])) // internal.acme.*
	assert.Check(t, !matcher.Contains(hostnames[4], hostnames[3])) // internal.acme.toys
	assert.Check(t, matcher.Contains(hostnames[4], hostnames[4]))  // internal.acme.telemetry.*
	assert.Check(t, matcher.Contains(hostnames[4], hostnames[5]))  // internal.acme.telemetry.foo

	// internal.acme.telemetry.foo
	assert.Check(t, !matcher.Contains(hostnames[5], hostnames[0])) // com.acme.*
	assert.Check(t, !matcher.Contains(hostnames[5], hostnames[1])) // com.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[5], hostnames[2])) // internal.acme.*
	assert.Check(t, !matcher.Contains(hostnames[5], hostnames[3])) // internal.acme.toys
	assert.Check(t, !matcher.Contains(hostnames[5], hostnames[4])) // internal.acme.telemetry.*
	assert.Check(t, matcher.Contains(hostnames[5], hostnames[5]))  // internal.acme.telemetry.foo

	paths := []string{
		"*",
		"admin",
		"admin/*",
		"orgs/*",
		"orgs/123",
	}

	matcher = PrefixMatcher{Separator: '/'}

	// *
	assert.Check(t, matcher.Contains(paths[0], paths[0]))  // *
	assert.Check(t, matcher.Contains(paths[0], paths[1]))  // admin
	assert.Check(t, matcher.Contains(paths[0], paths[2]))  // admin/*
	assert.Check(t, matcher.Contains(paths[0], paths[3]))  // orgs/*
	assert.Check(t, matcher.Contains(paths[0], paths[4]))  // orgs/123

	// admin
	assert.Check(t, !matcher.Contains(paths[1], paths[0])) // *
	assert.Check(t, matcher.Contains(paths[1], paths[1]))  // admin
	assert.Check(t, !matcher.Contains(paths[1], paths[2])) // admin/*
	assert.Check(t, !matcher.Contains(paths[1], paths[3])) // orgs/*
	assert.Check(t, !matcher.Contains(paths[1], paths[4])) // orgs/123

	// admin/*
	assert.Check(t, !matcher.Contains(paths[2], paths[0])) // *
	assert.Check(t, !matcher.Contains(paths[2], paths[1])) // admin
	assert.Check(t, matcher.Contains(paths[2], paths[2]))  // admin/*
	assert.Check(t, !matcher.Contains(paths[2], paths[3])) // orgs/*
	assert.Check(t, !matcher.Contains(paths[2], paths[4])) // orgs/123

	// orgs/*
	assert.Check(t, !matcher.Contains(paths[3], paths[0])) // *
	assert.Check(t, !matcher.Contains(paths[3], paths[1])) // admin
	assert.Check(t, !matcher.Contains(paths[3], paths[2])) // admin/*
	assert.Check(t, matcher.Contains(paths[3], paths[3]))  // orgs/*
	assert.Check(t, matcher.Contains(paths[3], paths[4]))  // orgs/123

	// orgs/123
	assert.Check(t, !matcher.Contains(paths[4], paths[0])) // *
	assert.Check(t, !matcher.Contains(paths[4], paths[1])) // admin
	assert.Check(t, !matcher.Contains(paths[4], paths[2])) // admin/*
	assert.Check(t, !matcher.Contains(paths[4], paths[3])) // orgs/*
	assert.Check(t, matcher.Contains(paths[4], paths[4]))  // orgs/123
}

func TestPrefixMatcherUnion(t *testing.T) {
	matcher := PrefixMatcher{}

	union := matcher.Union(
		[]string{"com.acme.foo", "org.acme.*"},
		[]string{"com.acme.foo", "com.acme.bar", "org.acme.baz"},
	)

	assert.Equal(t, len(union), 3)
	assert.Check(t, utils.SliceContains(union, "com.acme.foo"))
	assert.Check(t, utils.SliceContains(union, "com.acme.bar"))
	assert.Check(t, utils.SliceContains(union, "org.acme.*"))
}

func TestPrefixMatcherIntersection(t *testing.T) {
	matcher := PrefixMatcher{}

	intersection := matcher.Intersection(
		[]string{"com.acme.foo", "org.acme.*", "io.acme.*", "net.acme.foo"},
		[]string{"com.acme.foo", "com.acme.bar", "org.acme.baz", "net.acme.*"},
	)

	assert.Equal(t, len(intersection), 3)
	assert.Check(t, utils.SliceContains(intersection, "com.acme.foo"))
	assert.Check(t, utils.SliceContains(intersection, "org.acme.baz"))
	assert.Check(t, utils.SliceContains(intersection, "net.acme.foo"))
}
