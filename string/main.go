package main

import (
	"fmt"
)

func strStr(T, P string) int {
	if len(P) == 0 {
		return 0
	}

	i, j := 0, 0
	for i < len(T) {
		if T[i] == P[0] {
			for j = 1; j < len(P); j++ {
				if i+j >= len(T) || T[i+j] == P[j] {
					break
				}
			}

			if j == len(P) {
				return i
			}
		}
		i++
	}
	return -1
}

func strStr2(T, P string) int {
	if len(P) == 0 {
		return 0
	}

	length := len(P)
	for i := 0; i < len(T)-len(P)+1; i++ {
		if T[i:i+length] == P {
			return i
		}
	}
	return -1
}

const base = 16777619

/*
func robin_karp(T string, P string) int {
	n, m := len(T), len(P)
	if n < m || len(P) == 0 {
		return 0
	}

	var mult uint32 = 1
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hash(P)
	h := hash(T[:m])
	for i := 0; i < n-m+1; i++ {
		if h == hp {
			return i
		}
		if i > 0 {
			h = h - mult*uint32(T[i-1])
			h = h*base + uint32(T[i+m-1])
		}
	}
	return -1
}
*/

func KMP_PrefixTable(P string) (F []int) {
	F = make([]int, len(P))
	pos, cnd := 2, 0
	F[0], F[1] = -1, 0
	for pos < len(P) {
		if P[pos-1] == P[cnd] {
			cnd++
			F[pos] = cnd
			pos++
		} else if cnd > 0 {
			cnd = F[cnd]
		} else {
			F[pos] = 0
			pos++
		}
	}
	return F
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseXOR(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i] ^= r[j]
		r[j] ^= r[i]
		r[i] ^= r[j]
	}
	return string(r)
}

func isMatch(text, pattern string) bool {
	prev, now := make([]bool, len(pattern)+1), make([]bool, len(pattern)+1)
	for i := 0; i <= len(text); i++ {
		now, prev = prev, now
		now[0] = i == 0
		for j := 1; j <= len(pattern); j++ {
			if pattern[j-1] == '*' {
				now[j] = prev[j] || prev[j-1] || now[j-1]
			} else {
				now[j] = prev[j-1] && (text[i-1] == pattern[j-1] || pattern[j-1] == '?')
			}
		}
	}
	return now[len(pattern)]
}

/*
func reverseWords(sentence string) string {
	var words []string
	var start, end int

	for end < len(sentence) {
		for start < len(sentence) && sentence[start] == '' {
			start++
		}
		if start == len(sentence) {
			break
		}
		end = start + 1
		for end < len(sentence) && sentence[end] != '' {
			end++
		}
		words = append(words, sentence[start:end])
		start = end + 1
	}
	reverse(words)
	return strings.Join(words, " ")
}
*/

func permute(str []rune, i int) {
	if i == len(str) {
		fmt.Println(string(str))
	} else {
		for j := i; j < len(str); j++ {
			str[i], str[j] = str[j], str[i]
			permute(str, i+1)
			str[i], str[j] = str[j], str[i]
		}
	}
}

func combination(set []rune) (subsets []string) {
	length := uint(len(set))

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset string
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = subset + string(set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func combinations(str string) (subsets []string) {
	input := []rune(str)
	return combination(input)
}

func permutations(str string) {
	permute([]rune(str), 0)
}

func main() {
	fmt.Println(combinations("abc"))
	permutations("abc")
	fmt.Println(isMatch("CareerMonk Publications", "*ca?ions"))
}
