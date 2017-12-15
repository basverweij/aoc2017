package day10

import (
	"strconv"
	"strings"
)

type list []int

func newList(size int) list {
	l := make(list, size)
	for i := 0; i < size; i++ {
		l[i] = i
	}

	return l
}

func (l list) String() string {
	s := make([]string, len(l))

	for i, v := range l {
		s[i] = strconv.Itoa(v)
	}

	return strings.Join(s, " ")
}

func (l list) Reverse(pos, length int) {
	if length < 2 {
		// nothing to reverse
		return
	}

	n := len(l)
	for i, j := pos, pos+length-1; i < j; i, j = i+1, j-1 {
		l[i%n], l[j%n] = l[j%n], l[i%n]
	}
}
