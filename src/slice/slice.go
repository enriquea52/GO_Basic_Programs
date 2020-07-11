package main
import(

	"fmt"
	"sort"
	"strings"
	"strconv"
)

func main () {
	var number string
	var slice []int
	n:=1
	for n < 2{
		fmt.Println("Please enter an integer value: ")
		fmt.Scan(&number)
		if strings.ToLower(number) == "x" {
			fmt.Println("Done! ")
			break
		} else {
			i, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Not an integer bro, try again :)")
			} else{
				slice = append(slice,i)
				sort.Ints(slice)
				fmt.Println("Current slice content")
				fmt.Println(slice)}
			}
		
	}



}