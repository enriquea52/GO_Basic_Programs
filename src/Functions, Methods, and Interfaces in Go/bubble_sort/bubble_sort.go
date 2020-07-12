/*
This program implements the Bubble Sort algorithm in a function for sorting slices (of length 10) of integers 
receiving input from the user.
*/

package main

import(
	"fmt"
	"strconv"
)

func BubbleSort(numbers []int) {
	fmt.Println("sorting...")

	for i:=0; i < len(numbers) ; i++ {
		for j:=0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers,j)
			}

		}
	}


}

func Swap(sli []int,index int) {
	tmp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = tmp

}


func main() {
	fmt.Println("Please enter 10 integers\n")

	var number string
	integers := []int{}

	for i:=0; i<10;i++{
		fmt.Println("Please enter integer number[",i+1,"]: ")
		fmt.Scan(&number)

		x, err := strconv.Atoi(number)
		_ = err

		integers=append(integers,x)
		fmt.Print("Curent numbers are: ")
		fmt.Println(integers)


	}

	BubbleSort(integers)
	fmt.Println(integers)
	
}