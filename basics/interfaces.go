package main

import "fmt"

// fmt.Println takes an empty interface as an argument,
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
////// structs
type Engineer struct{
	name string
}
type Developer struct{
	name string
}

////// interface 
type EmployeeInterface interface{
	GetName() string
}
func (e *Engineer) GetName() string{
	return e.name
}
func (d *Developer) GetName() string{
	return d.name
}

func main(){
	a := 42
	b := 24.2
	testInterface(a,b)
	//-----------------
	var e EmployeeInterface
	var d EmployeeInterface
	e = &Engineer{"Bob"}
	d = &Developer{"Mark"}
	fmt.Println(e.GetName())
	fmt.Println(d.GetName())
}
