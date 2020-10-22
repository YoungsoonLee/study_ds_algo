package main

import "fmt"

func firstRepeatedChar(str string) byte {
	n := len(str)
	counters := [256]int{}

	//init
	for i := 0; i < 256; i++ {
		counters[i] = 0
	}

	for i := 0; i < n; i++ {
		if counters[str[i]] == 1 {
			return str[i]
		}
		counters[str[i]]++
	}

	return byte(0)

}

func main() {
	fmt.Printf("%c", firstRepeatedChar("abacd"))
}
