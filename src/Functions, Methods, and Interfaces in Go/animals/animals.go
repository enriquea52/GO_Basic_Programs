/*
This program is about using classes in go for programatically accessing values 
of objects created.
*/
package main
import(
	"fmt"
	"strings"
	"bufio"
	"os"

)

type Animal struct {
	name string
	food string
	locomotion string 
	noise string
}

func (x Animal) Eat() {
	fmt.Println("The ",x.name," eats: ",x.food)
}
func (x Animal) Move() {
	fmt.Println("The ",x.name," moves: ",x.locomotion)
}
func (x Animal) Speak() {
	fmt.Println("The ",x.name,"'s noise is: ",x.noise)
}


func main() {

	cow:=Animal{"cow","grass","walk","moo"}
	bird:=Animal{"bird","worms","fly","pepp"}
	snake:=Animal{"snake","mice","slither","hsss"}


	fmt.Println("Please introduce the niaml and the information you would like to get separated by space")
	fmt.Println("Menu")
	fmt.Println("1.- Animal Eat")
	fmt.Println("2.- Animal Move")
	fmt.Println("3.- Animal Speak")
	fmt.Println("")

	for {
		fmt.Print("> ")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		line:= in.Text()
		line = strings.ToLower(line)
		information:=strings.Split(line," ")


		switch information[0] {
	    case "cow":
	        
	    	 switch information[1] {
	    	 case "eat":
	        	cow.Eat()
	   		 case "move":
	        	cow.Move()
	   		 case "noise":
	       		 cow.Speak()
	    	 default:
	    		fmt.Println("Please type a proper info request for cow animal")
	    		continue
    	 	 }


	    case "snake":
	        
	        switch information[1] {
	    	 case "eat":
	        	snake.Eat()
	   		 case "move":
	        	snake.Move()
	   		 case "noise":
	       		 snake.Speak()
	    	 default:
	    		fmt.Println("Please type a proper info request for snake animal")
	    		continue
    	 	 }
	    case "bird":
	        
	    	switch information[1] {
	    	 case "eat":
	        	bird.Eat()
	   		 case "move":
	        	bird.Move()
	   		 case "noise":
	       		 bird.Speak()
	    	 default:
	    		fmt.Println("Please type a proper info request for bird animal")
	    		continue
    	 	 }


	    default:
	    	fmt.Println("Please type a proper name")
	    	continue
    	 }



	}
	



}