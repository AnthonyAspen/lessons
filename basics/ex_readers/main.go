package main

import "golang.org/x/tour/reader"

type MyReader struct{}

//  Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(str []byte)(int,error){
	for i := range str{
		str[i] = 65
	}	
	return len(str),nil
}

func main() {
	reader.Validate(MyReader{})
}

