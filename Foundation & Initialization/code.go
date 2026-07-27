package main

import "fmt"

// 1. Package-level variables are evaluated first.
// Go determines the order based on dependencies.
var globalA = initializeA()
var globalB = initializeB()

func initializeA() int {
	fmt.Println("1. Allocating globalA")
	return 10
}

func initializeB() int {
	fmt.Println("2. Allocating globalB")
	return 20
}

// 2. init() functions run after variable initialization but before main().
// You can have multiple init() functions in a single file.
func init() {
	fmt.Println("3. Running first init()")
}

func init() {
	fmt.Println("4. Running second init()")
}

func main() {
	// 3. main() is the final step of the boot process.
	fmt.Println("5. Executing main()")
}
