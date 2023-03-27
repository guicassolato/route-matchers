package utils

func Partition[T any](slice []T, belongsToP1 func(T) bool) (p1, p2 []T) {
	for _, v := range slice {
		if belongsToP1(v) {
			p1 = append(p1, v)
		} else {
			p2 = append(p2, v)
		}
	}
	return
}

func SliceContains[T comparable](slice []T, val T) bool {
	return SliceContainsFunc(slice)(val)
}

func SliceContainsFunc[T comparable](slice []T) func (T) bool {
	return func(val T) bool {
		for _, v := range slice {
			if v == val {
				return true
			}
		}
		return false
	}
}

func SliceToMap[T comparable, U any](slice []T, mapper func(T) (T, U)) map[T]U {
	m := make(map[T]U)
	for _, v := range slice {
		key, val := mapper(v)
		m[key] = val
	}
	return m
}

func SliceToSet[T comparable](slice []T) map[T]any {
  return SliceToMap(slice, func(v T) (T, any) { return v, nil })
}

func MapSlice[T, U any](slice []T, mapper func(T) U) []U {
	var s2 []U
	for _, v := range slice {
		s2 = append(s2, mapper(v))
	}
	return s2
}

func FilterSlice[T any](slice []T, filter func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func CopySlice[T any](s1 []T) []T {
	s2 := make([]T, len(s1))
	copy(s2, s1)
	return s2
}

func FlattenSlice[T, U any](slice []U, flatten func(U) []T) []T {
	var f []T
	for _, v := range slice {
		f = append(f, flatten(v)...)
	}
	return f
}

func AppendToSet[T comparable](slice []T, values ...T) []T {
	m := SliceToSet(slice)
	for _, v := range values {
		m[v] = nil
	}
	return Keys(m)
}

func Unique[T comparable](slice []T) []T {
	return Keys(SliceToSet(slice))
}

// Union returns the union between two sets A and B, given a provided 'contains' function.
// I.e. Intersection(A, B) → A ∪ B
func Union[T comparable](a, b []T, contains func(T, T) bool) []T {
	u := CopySlice(a)
	for _, vb := range b {
		found := false
		for _, vu := range u {
			if contains(vu, vb) {
				found = true
				break
			}
		}
		if !found {
			u = append(u, vb)
		}
	}
	return u
}

// Intersection returns the intersection between two sets A and B, given a provided 'contains' function.
// I.e. Intersection(A, B) → A ∩ B
func Intersection[T comparable](a, b []T, contains func(T, T) bool) []T {
  as, _ := Partition(a, func(va T) bool {
		for _, vb := range b {
			if contains(vb, va) {
				return true
			}
		}
		return false
	})
  bs, _ := Partition(b, func(vb T) bool {
		for _, va := range a {
			if contains(va, vb) {
				return true
			}
		}
		return false
	})
	return Unique(append(as, bs...))
}
