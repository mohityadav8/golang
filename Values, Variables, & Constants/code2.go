package main

import "fmt"

// iota increments automatically per line in a const block
const (
	ReadPermission  byte = 1 << iota // 1 << 0 == 0001
	WritePermission                  // 1 << 1 == 0010
	ExecutePermission                // 1 << 2 == 0100
	AdminPermission                  // 1 << 3 == 1000
)

// Untyped constants can hold values larger than any 64-bit integer
// as long as they are never explicitly typed or printed directly.
const HugeNumber = 1 << 100 
const ScaledDown = HugeNumber >> 99 // Evaluates back to 2 at compile time

func main() {
	// Bitmasking: Combining permissions using bitwise OR
	var myRole byte = ReadPermission | ExecutePermission // 0001 | 0100 = 0101

	// Checking permissions using bitwise AND
	if myRole&WritePermission != 0 {
		fmt.Println("Has Write Access")
	} else {
		fmt.Println("NO Write Access") 
	}

	// Short declaration (:=) vs var
	// := infers type but can cause accidental "Shadowing" (covered in next section)
	x := ScaledDown
	fmt.Printf("Scaled number: %d, Type: %T\n", x, x)
}
