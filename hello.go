package main

import (
	"fmt"
	"reflect"

	"time"

	"example.com/hello/calculator"
	"github.com/fatih/color"
	"github.com/leekchan/accounting"
	"rsc.io/quote"
)

// https://www.golangprograms.com/go-language/functions.html
// simple function print a message
func printMessage() {
	fmt.Println("call printMessage")
}

func add(a int, b int) int {
	return a + b
}

func rectangle(l int, b int) (parameters int, area int) {
	parameters = (l + b) * 2
	area = l * b
	return // return statement without variable name
}

// declare an interface
type Animal interface {
	speak()
}

// declare a struct
type Dog struct {
	name string
}

// implement a method for Dog
func (d Dog) Speak() {
	fmt.Printf("%s, name=%s", "Dog", d.name)
}

func main() {
	// fmt.Println("Hello, World!")
	fmt.Println(quote.Go()) //-> "rsc.io/quote"

	today := time.Now()
	switch today.Day() {
	case 11:
		fmt.Println("Today is 11th")
	case 12:
		fmt.Println("today 12th")
	default:
		fmt.Println("default")
	}

	// Print with default helper functions
	color.Cyan("Prints text in cyan.1")
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123)) // "$123,456,789.21"
	fmt.Println(ac.FormatMoney(12345678))         // "$12,345,678.00"
	fmt.Println("----------------")

	var i int
	for {
		fmt.Printf("%s%d", "-> run for i=", i)
		if i >= 10 {
			break
		}
		i++
	}
	fmt.Println("----------------")
	printMessage()

	fmt.Println("----------------")
	fmt.Println(add(2, 3))

	// array
	fmt.Println("Array----------------")
	var myArray = [...]int{1, 2, 3}
	fmt.Println((myArray))
	fmt.Println((myArray[1]))

	fmt.Println("---------------Example 2--------------------")
	for index, element := range myArray {
		fmt.Println(index, "=>", element)
	}

	// map
	fmt.Println("Map----------------")
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	fmt.Println(employee)

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// call func
	var p, a int
	p, a = rectangle(10, 20)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	var strSlice = []string{"India", "Canada", "Japan"}

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	b := make([]int, 2, 5)
	b[0] = 10
	b[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(b), cap(b))

	b = append(b, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", b)
	fmt.Printf("Length is %d Capacity is %d\n", len(b), cap(b))

	c := new(calculator.Calculator)
	fmt.Println(c.Add(2, 3))

	var d Dog = Dog{"Black dog"}
	d.Speak()
}
