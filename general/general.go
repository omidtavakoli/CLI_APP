package general

import (
	"strings"
)

type StringSlice struct {
	Slice []string
}
type IntegerSlice struct {
	Slice []int64
}

func (s StringSlice) Find(val string) (int, bool) {
	for i, item := range s.Slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (s IntegerSlice) Find(val int64) (int, bool) {
	for i, item := range s.Slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func CleanText(input []string) string {
	return "'" + strings.Join(input, `','`) + `'`
}
