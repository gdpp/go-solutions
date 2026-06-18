package main

import (
	"fmt"
)

func main() {
	//uint8: es para almacenar valores positivos hasta 255
	//uint16: es para almacenar valores positivos hasta 65535
	//uint32: es para almacenar valores positivos hasta 4294967295
	//uint64: es para almacenar valores positivos hasta 18446744073709551615
	//int8: es para almacenar valores positivos y negativos hasta 127
	//int16: es para almacenar valores positivos y negativos hasta 32767
	//int32: es para almacenar valores positivos y negativos hasta 2147483647
	//int64: es para almacenar valores positivos y negativos hasta 9223372036854775807

	views := 1000
	views2 := 2000
	totalViews := views + views2

	fmt.Println("Total views: ", totalViews)

	count := 10
	count++
	count++

	fmt.Println("Count: ", count)

	count--
	count--
	count--

	fmt.Println("Count: ", count)

	fmt.Println(totalViews / 2)

}
