package main

import "fmt"

// I realised that fmt.Println takes an empty interface as an argument,
//trying to figure out how it works exactly
func testInterface (a ...interface{}){
	for _,i := range a{
	fmt.Printf("value:%v , type: %T\n",i,i)
	}

}
func main(){

	a := 42
	b := 24.2
	testInterface(a,b)


}
