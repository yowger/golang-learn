package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
	// or this, go automatically dereference it
	// b.tip = tip
}

// add an item to the bill
// if updating a value pass pointer
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		// panic stops the program and prints the error
		panic(err)
	}

	fmt.Println("Bill saved to file")
}

/*
	In Go, **structs** and **custom types** offer some parallels to **interfaces** and **types** in TypeScript, but they have key differences due to the nature of the two languages.

	### Structs in Go
	In Go, a **struct** is a composite data type that groups together fields to form more complex data structures. This is similar to **TypeScript types or interfaces** that define the shape of an object. However, Go’s structs are more rigid and closer to actual **objects** than TS types.

	For example, in Go:

	```go
	type Person struct {
		Name string
		Age  int
	}
	```

	And in TypeScript:

	```typescript
	type Person = {
		name: string;
		age: number;
	};
	```

	In both cases, you define an object’s structure with specific fields, but with Go structs:
	- You can define methods on them, making them more like classes in TypeScript or JavaScript.
	- Structs are part of Go’s **static, compile-time typing**, which means they are fully enforced by the compiler.

	### Custom Types in Go
	A **custom type** in Go is often created by defining a new type based on an existing type. For example:

	```go
	type ID string
	type Age int
	```

	This is somewhat similar to **type aliases** in TypeScript, but it has some unique qualities:
	- In Go, defining `type ID string` creates a distinct type from `string`, even though `ID` is based on `string`. You can’t use an `ID` in place of a `string` without an explicit conversion.
	- Custom types are useful for adding specific functionality to a base type (e.g., methods).

	In TypeScript:

	```typescript
	type ID = string; // This is an alias, not a new type
	```

	### Interfaces in Go
	In Go, an **interface** defines a set of methods but doesn’t specify any implementation. This is similar to TypeScript interfaces that define the shape of an object but not the implementation. However, Go interfaces differ because they use **implicit implementation**:

	```go
	type Speaker interface {
		Speak() string
	}

	type Person struct {
		Name string
	}

	func (p Person) Speak() string {
		return "Hello, I'm " + p.Name
	}
	```

	In this example:
	- Any type with a `Speak()` method that returns a `string` is considered to implement the `Speaker` interface.
	- Unlike TypeScript, there’s no explicit declaration required to indicate that `Person` implements `Speaker`.

	### Summary
	- **Structs in Go** are similar to **objects or types** in TypeScript, but they’re more rigid and closer to classes in functionality.
	- **Custom types in Go** define distinct types from their base types, while **TypeScript type aliases** are not distinct types.
	- **Interfaces in Go** define behaviors through methods, similar to TypeScript interfaces, but Go uses **implicit implementation** (a type is considered to satisfy an interface as long as it has the required methods).
*/

// what is compiler

/*
	So, in summary, a compiler checks, optimizes, and translates code into machine-readable format. With every change in the source code, the compiler re-runs to catch errors, ensuring that only valid, type-safe code becomes an executable program.
*/
