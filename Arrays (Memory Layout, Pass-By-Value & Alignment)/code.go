package main

import (
	"fmt"
	"unsafe"
)

// In Go, the size is part of the type.
// [3]int and [4]int are completely different types.
func processArray(arr [3]int) {
	// Modifying a copy. The original array is untouched.
	arr[0] = 999 
	fmt.Printf("Inside processArray - Addr of arr[0]: %p (This is a copy)\n", &arr[0])
}

func processArrayPointer(arr *[3]int) {
	// Dereferencing the pointer to modify the original array.
	arr[0] = 999 
	fmt.Printf("Inside processArrayPointer - Addr of arr[0]: %p\n", &arr[0])
}

func main() {
	// 1. Memory Contiguity and Layout
	nums := [3]int{10, 20, 30}
	
	fmt.Println("--- Memory Layout ---")
	// Array elements are perfectly contiguous in memory.
	fmt.Printf("Addr of nums[0]: %p\n", &nums[0])
	fmt.Printf("Addr of nums[1]: %p\n", &nums[1])
	fmt.Printf("Addr of nums[2]: %p\n", &nums[2])
	fmt.Printf("Size of array: %d bytes (3 * 8 bytes on a 64-bit system)\n\n", unsafe.Sizeof(nums))

	// 2. The Pass-By-Value Trap
	fmt.Println("--- Pass By Value vs Pointers ---")
	fmt.Printf("Before processArray: %v\n", nums)
	processArray(nums) // Passes a complete 24-byte copy to the stack
	fmt.Printf("After processArray: %v (Unchanged)\n\n", nums)

	fmt.Printf("Before processArrayPointer: %v\n", nums)
	processArrayPointer(&nums) // Passes an 8-byte pointer
	fmt.Printf("After processArrayPointer: %v (Modified)\n", nums)
    
    // 3. Type Rigidity
    // var biggerNums [4]int 
    // nums = biggerNums // COMPILE ERROR: Cannot use [4]int as [3]int in assignment
}
