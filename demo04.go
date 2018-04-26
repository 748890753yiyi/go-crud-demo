package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	skillfolder := `E:\online-bak`
	files, _ := ioutil.ReadDir(skillfolder)
	for _,file := range files {
		if file.IsDir() {
			continue
		}else {
			fmt.Println(file.Name())
		}
	}
}