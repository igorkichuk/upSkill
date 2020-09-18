package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		func() {
			fmt.Println(v)
		}()
	}
}
