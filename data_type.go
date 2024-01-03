package main

import "fmt"

var b bool = true
var last_name string  = ""
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

var nums_slice []int
var nums_map map[string]int

var	user User = User{1, "John", "<EMAIL>"}

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 10
	first_name := "John"
	y := 50

	fmt.Println(x)
	fmt.Println(first_name)

	x = 20
	fmt.Println(x)

	first_name = "Jane"
	fmt.Println(first_name)

	fmt.Println(b)
	fmt.Println(last_name)

	fmt.Println(numbers)
	fmt.Println(numbers[3])

	fmt.Println(add(x, y))

	nums_slice = append(nums_slice, 1)
	fmt.Println(nums_slice)

	nums_map = make(map[string]int)
	nums_map["a"] = 1
	fmt.Println(nums_map)

	fmt.Println(user)
	fmt.Println(user.Email)
}

type User struct {
	ID int
	Name string
	Email string
}