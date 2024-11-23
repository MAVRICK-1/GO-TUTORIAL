package main

import (
	"time";
	"fmt"
)

func timeNow(){
	i:= time.Now()

	print(&i)

	fmt.Println(i)
	
}