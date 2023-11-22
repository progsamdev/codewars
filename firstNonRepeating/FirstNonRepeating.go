package firstnonrepeating

import (
	"strings"
)

/*
Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.


*/

// this solution uses map, so I must guarantee the order
// func FirstNonRepeating(str string) string {
// 	m := make(map[string]string)
// 	r := ""
// 	var keys []string
// 	for _, k := range str {
// 		keys = append(keys, strings.ToLower(string((k))))
// 	}

// 	for _, c := range str {
// 		k := strings.ToLower(string(c))
// 		_, seen := m[k]
// 		if seen {
// 			m[k] += string(c)
// 		} else {
// 			m[k] = string(c)
// 		}
// 	}
// 	for _, v := range keys {
// 		if len(m[v]) == 1 {
// 			r = m[v]
// 			break
// 		}
// 	}
// 	return r
// }

func FirstNonRepeating(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}
