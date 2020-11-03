package main

import "sort"

func maxEvents(start, fininsh []int) int {
	sort.Slice(start, func(i, j int) bool {
		if fininsh[i] == fininsh[j] {
			return start[i] < start[j]
		}
		return fininsh[i] < fininsh[j]
	})

	m := make(map[int]bool)
	for j := 0; j < len(start); j++ {
		for i := start[j]; i <= fininsh[j]; i++ {
			if _, ok := m[i]; !ok {
				m[i] = true
				break
			}
		}
	}
	return len(m)
}
