package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple": 	5,
		"pear":		6,
		"orange":	9,
	}

	mp["apple"] = 900
	fmt.Println(mp["apple"])

	val, ok := mp["banana"]
	fmt.Println(val, ok)
	fmt.Println(mp)
}