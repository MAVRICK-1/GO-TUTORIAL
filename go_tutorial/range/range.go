package main

import "fmt"

func main() {

	// nums:= [3]int{78, 79 ,80}


	// for i:=0; i< len(nums); i++ {
	// 	fmt.Println(nums[i])
	// } // 78 79 80

	// for i,num := range nums {
	// 	fmt.Println(i,num) 
	// }

	// Output:
	// 0 78
	// 1 79
	// 2 80

	dict := map[string]string{"a": "apple", "b": "banana"}

	for k,v := range dict {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k,v := range "Golang"{
		fmt.Printf("%d -> %c\n", k, v)
		fmt.Println( k, v) // 0 71 1 111 2 108 3 97 4 110 5 103
	}
	// 0 -> G
	// 1 -> o
	// 2 -> l
	// 3 -> a
	// 4 -> n
	// 5 -> g




	nums := []int{2, 3, 4}
	sum := 0

	// range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // index: 1
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	// range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// a -> apple
	// b -> banana

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// key: a
	// key: b

	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	// 0 103
	// 1 111
}