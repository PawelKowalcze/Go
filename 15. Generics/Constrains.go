package main

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		// this is where the .String() method
		// is used. That's why we need a more specific
		// constraint instead of the any constraint
		result += val.String()
	}
	return result
}
