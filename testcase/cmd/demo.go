package main

import "fmt"

func main(){
	buf := make([]byte, 8)
	a := make([]byte, 8)
	buf = append(buf[:0], a[:3]...)

	fmt.Println(buf)
}
