package matchers

type Store interface{
	// operations
	Add(string)
	List() []string
	Size() int

	// options
	SetMatcher(Matcher)
}

type StoreOption func(Store)

func WithMatcher(matcher Matcher) StoreOption {
	return func(store Store) {
		store.SetMatcher(matcher)
	}
}
