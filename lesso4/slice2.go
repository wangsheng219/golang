package main

import "fmt"

func main() {
	names := [4]string{
		"a",
		"b",
		"c",
		"d",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]

	b[0] = "xxx"
	fmt.Println(a,b)
	fmt.Println(names)



	var p [2]*string
	p[0] = &names[0]
	p[1] = &names[1]
	*p[0] = "AAAA"

	
}