package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	// * is the dereference operator so that we can safely update the value
	//  "tifa"  -> "wedge"
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

}

/*

|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|

*/

/*
	In Go, pointers are variables that hold the memory address of another variable. Instead of storing a direct value, a pointer refers to (or "points to") the location in memory where a value is stored.
*/

/*
	package main

	import "fmt"

	func main() {
		a := 10
		var p *int = &a  // p is a pointer to the memory address of a

		fmt.Println("Value of a:", a)
		fmt.Println("Memory address of a:", &a)
		fmt.Println("Value of p:", p)      // p holds the memory address of a
		fmt.Println("Value at p:", *p)     // *p dereferences p, giving the value stored at that address (10)

		*p = 20  // modifies the value at the memory address pointed to by p
		fmt.Println("Updated value of a:", a)  // now a is 20
	}
*/

/*
	Explanation

	1. Pointer Declaration and Initialization:
		var p *int = &a: This line declares p as a pointer to an integer (*int), and assigns it the memory address of a using the & operator.

	2. Dereferencing:
		*p = 20: The * symbol before a pointer variable (in this case, *p) accesses the value at that memory address. Changing *p also changes the value of a.

	3. Result:
		Modifying *p changes a directly since p points to a's memory address.
*/

/*
	Why Pointers are Important in Go

	1. Efficiency: Passing a large structure or slice by value (copying it) in function calls can be expensive. By using a pointer, you pass a reference to the data, making the program faster and less memory-intensive.

	2. Mutability: Go function parameters are passed by value by default, meaning changes to a variable within a function do not affect the original variable. However, if you pass a pointer, the function can modify the original variable directly.

	3. Avoiding Copies in Large Data Structures: For complex data structures, such as slices, maps, or custom structs, pointers can prevent unnecessary data duplication, keeping code cleaner and more performant.
*/

/*
	package main

	import "fmt"

	func increment(value *int) {
		*value += 1  // Modify the original value by dereferencing the pointer
	}

	func main() {
		num := 5
		increment(&num)  // Pass the address of num
		fmt.Println("Incremented value:", num)  // Output: 6
	}
*/

/*
	In this example:

	The increment function receives a pointer, which allows it to directly modify num in the main function.
	Pointers help write efficient, performance-focused code in Go, especially in cases involving large data or when values need to be modified across functions.
*/

// --------------

/*
	When to Use a Pointer to a Slice
	You might want to use a pointer to a slice in the following cases:

	1. Nil Checks: To differentiate between a slice that is empty (len(slice) == 0) and one that is nil.

	2. Memory Efficiency: If you're dealing with a very large slice and creating additional references, using a pointer could avoid some small, additional overhead.

	3. Reslicing: When you want to modify the slice itself (e.g., append or resize the slice within a function and have that change reflected in the caller).
*/

/*
	package main

	import "fmt"

	func modifyAndResizeSlice(s *[]int) {
		*s = append(*s, 6, 7, 8)  // Modify the underlying slice by adding elements
	}

	func main() {
		slice := []int{1, 2, 3, 4, 5}
		modifyAndResizeSlice(&slice)
		fmt.Println("Resized slice:", slice)  // Output: [1 2 3 4 5 6 7 8]
	}
*/

/*
In this example, the function modifies the actual slice structure by appending new elements, which wouldn't have been possible if we passed the slice directly without a pointer.

Summary
	For read-only or simple modifications to the data in a slice, passing the slice directly is efficient and sufficient.

	If you need to resize the slice or handle a nil slice, passing a pointer to the slice might be beneficial.
*/
