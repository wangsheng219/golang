package main

import "fmt"

func main(){

	s := []int{2,3,5,7,11,13}
	printSlice(s)

	s = s[:0]
	printSlice(s)


	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	var s1 []int

	if s1 == nil {
			fmt.Println("nil")
	}
	fmt.Println(len(s1))


}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}