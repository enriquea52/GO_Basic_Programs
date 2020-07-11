package main
import (
	"fmt"
)

func main() {
	var number_float float64

	fmt.Println("Please introduce a floating point number")
	fmt.Scan(&number_float)

	fmt.Println(int(number_float))
}