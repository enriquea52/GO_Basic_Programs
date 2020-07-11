/*
This program reads first and last names from a .txt file 
and prints them into the terminal
*/
package main
import(
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
)
func main() {
	type name struct{
		fname string
		lname string 
	}

	var ind name
	list := make([]name, 0)
	split := make([]string, 2)

	file,err := os.Open("names.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		split = strings.Split(scanner.Text()," ")
		ind.fname = split[0]
		ind.lname = split[1]
		
		list = append(list,ind)
	}
 
	file.Close()
 
	for _, eachline := range list {
		fmt.Println(eachline)
	}
	
}