package main

import "fmt"

func main(){
	sum := 1
	// init and post statements are optional
	for ; sum<10; {
		sum += sum
	}

	fmt.Println(sum)

}
