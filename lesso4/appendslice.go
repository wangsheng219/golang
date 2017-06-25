package main

import "fmt"

func main(){

	var s []int
	printSlice(s)

	s = append(s,0)
	printSlice(s)

	s = append(s,1)
	printSlice(s)


	s = append(s,2,3,4)
	printSlice(s)

	l := make([]int,0,1)
	_ = append(l,1)
	fmt.Println(l)

	l = append(l,2)
	fmt.Println(l)


	l1 := []int{13,14,15}


	l = append(l,l1...)
	fmt.Println(l)
}

func printSlice(s []int){

	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}