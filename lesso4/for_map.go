package main


import "fmt"

func main(){
	ages := map[string]int{
		"a":1,
		"b":2,
		"c":3,
		"d":4,
	}

	for name,age :=range ages{

		fmt.Println("names",name,"ages",age)
	}
	for name := range ages{
		fmt.Println(name)
	}





}

