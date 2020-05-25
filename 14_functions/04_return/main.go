package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
	fmt.Println(greet("Bharat ", "Sarvan"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
