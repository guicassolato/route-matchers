package matchers

const (
	DefaultSetSeparator   = ','
	DefaultComplementSign = '^'
)

var DefaultMatcher Matcher = &StringMatcher{}

type Matcher interface {
	Contains(string, string) bool
	Union([]string, []string) []string
	Intersection([]string, []string) []string
	Zip([]string) string
	Complement(string) string
}
