package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte)(int,error){
	n,err := rot.r.Read(b)
	if err == io.EOF{
		return 0,err
	}
	for i :=0;i <n;i++{
		if b[i] >= 'a' && b[i] <= 'z'{
			if b[i] >= 'm'{
				b[i] -= 13
			}else{
				b[i] += 13
			}
		}
		if b[i] >= 'A' && b[i] <= 'Z'{
			if b[i] >= 'M'{
				b[i] -= 13
			}else{
				b[i] += 13
			}
		}
	}

	return len(b),nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

