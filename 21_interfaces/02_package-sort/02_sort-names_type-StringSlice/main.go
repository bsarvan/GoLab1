package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny", "Bharat"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	r := sort.StringSlice(s).Search("Cat")

	fmt.Println(r)
	//sort.Sort(sort.StringSlice(s))
	//r := sort.StringSlice(s)
	//fmt.Println(r)

}

// https://golang.org/pkg/sort/#Sort
