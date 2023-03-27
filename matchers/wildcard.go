package matchers

const DefaultWildcard = '*'

type Wildcard rune

func (m Wildcard) wildcard() rune {
	if m != 0 {
		return rune(m)
	}
	return DefaultWildcard
}
