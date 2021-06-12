package calculator // declare package name

// import another package
import "fmt"

// init this package
func init() {
	fmt.Println("calculator package init")
}

// declare a struct with name [Calculator]
type Calculator struct {
	a int
	b int
}

// implement a method in a struct [Calculator] with name is [Add]
func (c Calculator) Add(a int, b int) int {
	c.a = a
	c.b = b
	result := c.a + c.b
	return result
}
