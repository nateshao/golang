package main

import (
	"strings"
)

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a code1_slice of the substrings between those separators.
//func Split(s, sep string) (result []string) {
//	i := strings.Index(s, sep)
//
//	for i > -1 {
//		result = append(result, s[:i])
//		s = s[i+1:]
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}

func Split(s, sep string) (result []string) {
	//i := strings.Index(s, sep)
	//
	//for i > -1 {
	//	result = append(result, s[:i])
	//	s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
	//	i = strings.Index(s, sep)
	//}
	//result = append(result, s)
	//return

	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
