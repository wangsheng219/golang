package main

import (
	"fmt"
	"io/ioutil"
	//"os"
	"strings"
)

func main(){

	//f, err := os.Open("/a.txt")
	//if err != nil {
	//	continue
	//}


	//data := ioutil.ReadFile("a.txt")
	data, _ := ioutil.ReadFile("a.txt")

	s := string(data)

	c := strings.Fields(s)


	//fmt.Println(c)
	//fmt.Println("Fields are: %q", strings.Fields(s))

	//for  k,v := range s{
	//	fmt.Prinln(k,v)
	//}
	count_map := make(map[string]int)
	for _,v := range c{
		//fmt.Println(v)
		//if _,ok := c[v];ok{
		//	fmt.Println(1)
		//}

		if _,ok := count_map[v];ok{
			count_map[v] +=1
			}else{
			count_map[v] =1
		}
		

		
	}
}	