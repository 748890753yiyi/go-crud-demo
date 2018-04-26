package main

import "fmt"

func main()  {
	/*a := [2]int{1,2}
	fmt.Println(a)
	b := [10]int{}
	b[1] = 2
	fmt.Println(b)

	c := new([10]int)
	c[1] = 3
	fmt.Println(c)

	d := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(d)*/
	a := [...]int{1,4,9,2,5,0,13,3}
	num := len(a)
	for i := 0;i<num ;i++  {
		for j:=i+1;j<num ;j++  {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)

}

