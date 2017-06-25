package main

import "fmt"

func main(){
	s := []int{2,3,5,7,11,13}


	fmt.Println(s[1])
	fmt.Println(s[1:])
	//s = revego()
	//printSlice(s)
	abc := []int{2,3,4,5,11,12,13,15}

	///fmt.Println(len(abc))
	var i int
	for i=0;i<len(abc)/2;i++{

		//fmt.Println(i)
		//fmt.Println(abc[i])

		fmt.Println(abc[len(abc)-i-1])
		abc[i],abc[len(abc)-i-1] = abc[len(abc)-i-1],abc[i]

	}
	fmt.Println(abc)
		

}

//func reverse(s []int,n){
//	var a = a[n:]
//	return s 
//}



func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}