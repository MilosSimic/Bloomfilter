# Bloomfilter
implementation of standard bloom filter algorithm

# Usage:
```go
package main

import (
	"fmt"
)

func main() {
	bf := New(300000, 0.01)
	bf.Add("Dog")
	bf.Add("Cat")

	fmt.Println("Test Dog [true]:", bf.Test("Dog"))
	fmt.Println("Test Cat [true]:", bf.Test("Cat"))
	fmt.Println("Test John[false]:", bf.Test("John"))
	fmt.Println("Test Doe[false]:", bf.Test("Doe"))

	bf.Add("Pig")
	bf.Add("Mice")
	bf.Add("John")

	fmt.Println("Test Pig [true]:", bf.Test("Pig"))
	fmt.Println("Test Mice [true]:", bf.Test("Mice"))
	fmt.Println("Test John[true]:", bf.Test("John"))
	fmt.Println("Test Mick[false]:", bf.Test("Mick"))
}
```
