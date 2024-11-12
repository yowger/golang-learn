package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	fmt.Println(mybill.format())
}

/*
	Yes, you're exactly right! In Go, this approach is similar to **methods on a class** in object-oriented programming languages like Java or Python.

	### Explanation of Your Code

	In Go, when you define a function with a **receiver** (like `(b bill)` in `func (b bill) format()`), you're associating that function (or method) with a specific **type** (`bill` in this case). This way, you can call the method on any instance of that type, just like you would with a class in other languages.

	Here's a breakdown:

	```go
	func (b bill) format() string {
		// method code
	}
	```

	- **`(b bill)`** is the **receiver** part of the function. It means that `format` is a method associated with the `bill` type.
	- `b` is the variable name for the receiver, similar to `this` or `self` in other languages. You can use `b` to access fields and other methods of `bill`.
	- After defining this method, you can call it on any instance of `bill` as you mentioned, like `myBill.format()`.

	### Example Usage
	Let's see how this works with an example:

	```go
	package main

	import (
		"fmt"
	)

	// Define the bill struct
	type bill struct {
		name  string
		items map[string]float64
		tip   float64
	}

	// Constructor function for a new bill
	func newBill(name string) bill {
		b := bill{
			name:  name,
			items: map[string]float64{"pie": 5.99, "cake": 3.99},
			tip:   0,
		}
		return b
	}

	// Method to format the bill details as a string
	func (b bill) format() string {
		fs := "Bill breakdown:\n"
		var total float64 = 0

		// List items in the bill
		for k, v := range b.items {
			fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
			total += v
		}

		// Add total to the breakdown
		fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
		return fs
	}

	func main() {
		myBill := newBill("Mario's Bill")
		fmt.Println(myBill.format()) // Calling the format method on the myBill instance
	}
	```

	### Output:
	```
	Bill breakdown:
	pie:                     ...$5.99
	cake:                    ...$3.99
	total:                   ...$9.98
	```

	### How This Compares to Classes
	- **Encapsulation**: Just like methods in a class, methods on a struct allow you to encapsulate behavior with data, grouping related functionality together.
	- **Structs vs. Classes**: Go doesn’t have classes, but structs + methods provide a similar experience by grouping fields and behaviors. However, Go doesn’t support inheritance, which is a core concept in many object-oriented languages.
	- **Type-Specific Functions**: By associating methods with specific types, you can define different behaviors for different types without using inheritance. This keeps Go simple and focuses on **composition over inheritance**.

	In summary, `func (b bill) format()` is a method that can be called on instances of `bill`, working similarly to methods on classes in other languages.
*/
