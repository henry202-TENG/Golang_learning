package int_try

import "fmt"

func Print_int(input int) int {
	fmt.Println(input)
	return 0
}

func Test_base_int() int {
	var a_bool bool = true
	var a_int int = 655360000
	var a_int8 int8 = 127
	var a_int16 int16 = 32767
	var a_int32 int32 = 2147483647
	var a_int64 int64 = 9223372036854775807

	int_without_type := 2147483647

	fmt.Println("=====int=====")
	fmt.Println("bool: ", a_bool, fmt.Sprintf("%T", a_bool))
	fmt.Println("int: ", a_int, fmt.Sprintf("%T", a_int))
	fmt.Println("int8: ", a_int8, fmt.Sprintf("%T", a_int8))
	fmt.Println("int16: ", a_int16, fmt.Sprintf("%T", a_int16))
	fmt.Println("int32: ", a_int32, fmt.Sprintf("%T", a_int32))
	fmt.Println("int64: ", a_int64, fmt.Sprintf("%T", a_int64))
	fmt.Println("int_without_type:= ", int_without_type, fmt.Sprintf("%T", int_without_type))
	return 0
}
