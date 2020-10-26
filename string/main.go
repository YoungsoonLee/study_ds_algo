package main

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
