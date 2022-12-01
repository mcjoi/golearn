package main



import (
	"reflect"
	"fmt"
)


func main() {
	var a = "initial"
	fmt.Println(a)
	
	var b, c int = 1,2
	fmt.Println(b, c)
	
	var d = true
	fmt.Println(d)
	
	var e int
	fmt.Println(e) // auto init : int null value is 0
	
	var e2 string
	fmt.Println(e2) // auto init : string null value is ""
	
	f := "apple"
	fmt.Println(f) // The := syntax is shorthand for declaring and initializing a variable
	
	h, i := "myname", "yourname"
	fmt.Println(h, i)
	
	j, k := 1235, 486.0
	fmt.Println(j, k)
	fmt.Println(reflect.TypeOf(j), reflect.TypeOf(k)) // check the type of variables
	
	
	
}