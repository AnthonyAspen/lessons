package main

import (
	"fmt"
	"math"
)

        
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
// compute(hypot(x,y) float64) float64{
// return hypot(3,4)
//}
func hypot(x,y float64) float64{
	return math.Sqrt(x*x + y*y)
}

func main() {

	                //25, 144  = 169
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

