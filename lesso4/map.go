package main

import "fmt"


func main(){

	ages := make(map[string]int)
	ages["a"] = 1
	ages["b"] = 2

	fmt.Println(ages["a"])
	ages["a"] = ages["b"]+2

	c,ok :=ages["c"]
	
	fmt.Println(ages["c"])
	if ok{
		fmt.Println(c)
	}else{
		fmt.Println("not found")
	}
	if c,ok := ages["c"];ok {
		fmt.Println(c)
	}
	fmt.Println(ages)

}