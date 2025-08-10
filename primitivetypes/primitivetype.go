package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// key word declaring
	var myVariable string
	fmt.Println(myVariable)

	var myOtherVariable = "value for my other varible"
	fmt.Println(myOtherVariable)

	// short declation

	myShortDeclationVariable := "value for my another varible"

	fmt.Println(myShortDeclationVariable)

	// interger types

	//signed type
	myInt := -1 // default system - ame as int32 or int64 depending on architecture
	fmt.Println(myInt)
	myInt8 := int8(10) //	range -128 to 127 - here we need a down cast
	fmt.Printf("myInt8 value: %d - that's not much \n", myInt8)
	myInt16 := int16(-32768) //	range -32768 to 32767 - here we need a downcast also
	fmt.Printf("myInt16 value: %d - that's better, but steel not much\n", myInt16)
	myInt32 := int32(2147483647) //	range -2147483648 to 2147483647 - here to we need a downcast
	fmt.Printf("myInt32 value: %d - that's a lot\n", myInt32)
	myInt64 := int64(9223372036854775807) //	range -9223372036854775808 to 9223372036854775807 - here to we need a downcast
	fmt.Printf("myInt64 value: %d - that's more than enough\n", myInt64)

	//unsigned types

	myUint := uint(1) // default system - ame as int32 or int64 depending on architecture
	fmt.Println(myUint)
	myUint8 := uint8(255) //	range 0 to 255 - here we need a down cast
	fmt.Printf("my unsigned Int8 value: %d - that's not much \n", myUint8)
	myUint16 := uint16(65535) //	range 0 to 65535 - here we need a downcast also
	fmt.Printf("my unsigned Int16 value: %d - that's better, but steel not much\n", myUint16)
	myUint32 := uint32(4294967295) //	range 0 to 4294967295 - here to we need a downcast
	fmt.Printf("my unsigned Int32 value: %d - that's a lot\n", myUint32)
	myUint64 := uint64(18446744073709551615) //	range 0 to 18446744073709551615 - here to we need a downcast
	fmt.Printf("my unsigned Int64 value: %d - that's more than enough\n", myUint64)

	//spscial types

	myRune := 'ç'
	fmt.Printf("myRune caracter: %c - myRune integer (%d)\n", myRune, myRune)

	myCaracter := byte(65)
	fmt.Printf("myCaracter caracter: %c - myCaracter integer (%d)\n", myCaracter, myCaracter)

	ptr := uintptr(unsafe.Pointer(&myUint))

	fmt.Println("uintptr (endereço de x):", ptr)

}
