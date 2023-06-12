package main

import (
	/*  import邏輯是import資料夾，每個資料夾屬於一個module，go 只分module不分檔案，
	    同個module內，不同檔案之間的function可以直接用。
	*/
	"fmt"
	"mod_learning/src/variables/int_try"
)

func main() {
	fmt.Println("Hello, World!")

	// int
	int_try.Print_int(12365)
	int_try.Test_base_int()

}
