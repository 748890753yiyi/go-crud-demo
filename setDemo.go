package main

import "fmt"

/*func f() (r int)  {
	r = 5
	defer func() {
		r=r+5
	}()
	return
}*/
func f() (r int)  {
	t := 5
	defer func() {
		t=t+5
		fmt.Printf("t: %d\n",t)
	}()
	r = t
	return
}
func main() {
	fmt.Printf("f(): %d",f())
}