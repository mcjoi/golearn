package main



import (
	//"reflect" // typeof
	"fmt" 
	// "math"
)



func main() {
	
	
	fmt.Println("#1 example")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	fmt.Println("#2 example")
	for j:=7; j <=9; j++ {
		fmt.Println(j)
	}
	
	fmt.Println("#3 example")
	for n := 3; n <= 10; n++ {
		if n%2 == 0 {
			 continue // or use break
		}
		fmt.Println(n)
	}
	
	
	
}