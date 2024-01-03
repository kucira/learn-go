package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(name, age, firstName, lastName)
}

var name string = "World"
var age int = 42
var (
	firstName string = "John"
	lastName  string = "Doe"
	language string = "Go"
)

if(age > 18) {
	fmt.Println("Adult")
}
else {
	fmt.Println("Minor")
}

for i := 0; i < count; i++ {
	fmt.Println(i)
}

switch language {
case "Go":
	fmt.Println("oke")
default:
	fmt.Println("oke juga")
}
