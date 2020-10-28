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

func removeAdjacentPairs(str string) string {
	r := []rune(str)
	n, i, j := len(r), 0, 0
	for i = 0; i < n; i++ {
		if j > 0 && i < n && (r[i] == r[j-1]) {
			j--
		} else {
			r[j] = r[i]
			j++
		}
	}
	return string(r[:j])
}

func minWindow(S string, T string) string {
	rem := 0
	shouldFind := make(map[byte]int)
	for i := range T {
		rem++
		shouldFind[T[i]]++
	}

	if rem > len(S) {
		return ""
	}

	var hasFound = string(make([]byte, len(S)))
	start, end := 0, 0
	for end < len(S) {
		if v, ok := shouldFind[S[end]]; ok {
			if v > 0 {
				rem--
			}
		}
		shouldFind[S[end]]--
	}

	for rem <= 0 {
		if len(hasFound) >= len(S[start:end+1]) {
			hasFound = S[start : end+1]
		}
		if _, ok := shouldFind[S[start]]; ok {
			shouldFind[S[start]]++
			if shouldFind[S[start]] > 0 {
				rem++
			}
		}
		start++
	}

	return ""
}

func findMatch(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	// init 2d array
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	pos := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(visited, board, i, j, word, pos) {
				return true
			}
		}
	}
	return false
}

func helper(visited [][]bool, board [][]byte, r, c int, word string, pos int) bool {
	if pos == len(word) {
		return true
	}

	if r < 0 || r == len(board) || c < 0 || c == len(board[0]) {
		return false
	}

	if board[r][c] != word[pos] || visited[r][c] {
		return false
	}

	visited[r][c] = true
	// check 8 neighbors
	if helper(visited, board, r-1, c, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c, word, pos+1) {
		return true
	}
	if helper(visited, board, r, c-1, word, pos+1) {
		return true
	}
	if helper(visited, board, r, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r+1, c-1, word, pos+1) {
		return true
	}
	if helper(visited, board, r-1, c+1, word, pos+1) {
		return true
	}
	if helper(visited, board, r-1, c-1, word, pos+1) {
		return true
	}

	visited[r][c] = false
	return false
}

func main() {
	fmt.Println(combinations("abc"))
	permutations("abc")
	fmt.Println(isMatch("CareerMonk Publications", "*ca?ions"))
}
