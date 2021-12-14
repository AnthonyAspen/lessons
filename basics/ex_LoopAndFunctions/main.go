package main

import (
	"fmt"
	"math"
)

// Sqrt without handling an error it's a first execrice
// then I added an error handling 
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("Cannot Sqrt negarive number: %f",e)

}
//z -= (z*z - x) / (2*z)
func Sqrt(x float64) (float64,error) {
	if x < 0 {
		return 0,ErrNegativeSqrt(x)
	}
	z:= float64(1)
	t:= float64(0)

	for {
		z -= (z*z - x) / (2*z)
		// comparirng difference between values
		if diff := math.Abs(z-t); diff < 0.0001{
			return z,nil
		}
		//
		t = z
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

