package main


import (
	"fmt"
	"os"
	//"log"
	"io/ioutil"
	"crypto/md5"
)



func main(){
	file1 := os.Args[1]
	
	//f,err := os.Open(file1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	
	//fmt.Printf("%X\n",f)
	//&{2081C2150}
	filebuf, _ := ioutil.ReadFile(file1)
	md5sum := md5.Sum(filebuf)
	fmt.Printf("%X\n",md5sum)


}
