package main


import "fmt"

func main() {
	
	var s string
	var n int

	// /var line string

	for{
			fmt.Print(">   ")
			fmt.Scan(&s,&n)
			if s == "stop" {
				break
			}	
			fmt.Println(s,n)
	}




}