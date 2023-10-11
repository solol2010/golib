package strs

import "strings"

// SplitLast by xp, get the last string of the splitted slice
func SplitLast(s, sep string) string {
	return Reverse(strings.Split(Reverse(s), Reverse(sep))[0])
}

// SplitFirst by xp, get the fisrt string of the splitted slice
func SplitFirst(s, sep string) string {
	return strings.Split(s, sep)[0]
}

// Reverse by xp, reverse the string
func Reverse(s string) string {
	length := len(s)
	if length > 1 {
		t := make([]byte, length)
		for i := 0; i < length; i++ {
			t[length-i-1] = s[i]
		}
		return string(t)
	} else {
		return s
	}
}

// TrimMultiSpace by xp, get rid of all spaces
func TrimMultiSpace(s string) (r string) {
	t := s
	for {
		r = strings.Trim(t, " ")
		if t == r {
			return
		} else {
			t = r
		}
	}
}
