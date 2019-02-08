package main

import(
	"fmt"
	"time"
	"math/rand"
	"math"
)

type Vertex struct { // normal struct with values
	X, Y int
}

var m = map[string]int{ // Map in go work as a hash
	"one": 1,
	"two": 2,
	"three": 3,
}

type Dog struct{} // struct with functions

func (d *Dog) dogName() {  // Like prototype in js
	fmt.Println("Juancho")	
}


func add(x, y int) int {
	return x + y
}

func name(x, y string) string {
	return x + " " + y
}

// k := 3 // declaring variable
// var k int = 3// also declaring a variable
// const Pi = 3.14 // way to declare a constante
// defer is like a return but used in cycles, return outside of it
// s := []int{2, 3, 5, 7, 11, 13} // declaring Arrays

// for i, v := range pow {
// 	} range act like a for with pow variable limit
func main() {
	v := Vertex{5, 3}
	dog := Dog{}
	fmt.Println("Hey")
	fmt.Println("The time is:", time.Now())
	fmt.Println("My fav number is: ", rand.Intn(20))
	fmt.Println("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(10,30))
	fmt.Println(name("cris", "pala"))
	fmt.Println(Vertex{2,1}) // connect with the struct vertex
	fmt.Println(v.Y) // LIke hash in ruby, printing the value from V in the top
	dog.dogName()
	fmt.Println(m["two"])
}
