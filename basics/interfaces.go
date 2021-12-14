package main

import "fmt"

// I realised that fmt.Println takes an empty interface as an argument,
//trying to figure out how it works exactly
// https://go.dev/tour/methods/14
func testInterface (a ...interface{}){

	for _,i := range a{
		_,ok := i.(int)
		if ok {
	 fmt.Printf("value:%v , type: %T\n",i,i)
		} 

	// s := i.(int)
	// fmt.Printf("value:%v , type: %T\n",s,s)
	}

}
func main(){

	a := 42
	b := 24.2
	testInterface(a,b)


}
