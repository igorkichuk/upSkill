package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := []int{5, 12, 44, 11, 101}
	sort.Slice(slice, func(i, j int) bool {
		iString := strconv.Itoa(slice[i])
		jString := strconv.Itoa(slice[j])
		sumI, sumJ := 0, 0
		for _, dig := range iString {
			digInt, _ := strconv.Atoi(string(dig))
			sumI += digInt
		}
		for _, dig := range jString {
			digInt, _ := strconv.Atoi(string(dig))
			sumJ += digInt
		}
		return sumI < sumJ
	})

	fmt.Println(slice)
}
