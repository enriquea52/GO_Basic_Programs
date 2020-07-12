/*
This program is about using interfaces in go for programatically accessing values 
of objects created in runtime.
*/
package main
import(
	"fmt"
	"strings"
	"bufio"
	"os"

)

type Animal interface{
	Eat ()
	Move ()
	Speak ()
}

type Cow struct {
	name string
	food string 
	locomotion string 
	noise string 
}

type Bird struct {
	name string
	food string
	locomotion string 
	noise string 
}

type Snake struct {
	name string 
	food string 
	locomotion string 
	noise string 
}

func (x Cow) Eat() {
	fmt.Println("The cow",x.name," eats: ",x.food)
}
func (x Cow) Move() {
	fmt.Println("The cow",x.name," moves: ",x.locomotion)
}
func (x Cow) Speak() {
	fmt.Println("The cow",x.name,"'s noise is: ",x.noise)
}

func (x Bird) Eat() {
	fmt.Println("The bird",x.name," eats: ",x.food)
}
func (x Bird) Move() {
	fmt.Println("The bird",x.name," moves: ",x.locomotion)
}
func (x Bird) Speak() {
	fmt.Println("The bird",x.name,"'s noise is: ",x.noise)
}
func (x Snake) Eat() {
	fmt.Println("The snake",x.name," eats: ",x.food)
}
func (x Snake) Move() {
	fmt.Println("The snake",x.name," moves: ",x.locomotion)
}
func (x Snake) Speak() {
	fmt.Println("The snake",x.name,"'s noise is: ",x.noise)
}

func inform (a Animal, sel string) {
	if sel == "eat" {
		a.Eat()
	} else if sel == "speak" {
		a.Speak()
	} else if sel == "move" {
		a.Move()
	} else {
		fmt.Println("Please introduce a proper command")
	}
}



func main() {

	m := make(map[string]interface{	Eat ()
	Move ()
	Speak ()
})

	

	fmt.Println("Options Menu:")
	fmt.Println("")
	fmt.Println("1.- newanimal name type")
	fmt.Println("2.- query name comand")
	fmt.Println("")
	fmt.Println("Newanimal types are: Cow   Bird   Snake")
	fmt.Println("Query commands are: Eat   Move   Speak")

	for {
		fmt.Print("> ")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		line:= in.Text()
		line = strings.ToLower(line)
		information:=strings.Split(line," ")


		if information[0] == "newanimal" {
			switch information[2]{
				case "cow":
					m[information[1]] = Cow{information[1],"grass","walk","moo"}
					fmt.Println("Created it!","cow named",information[1])
				case "bird":
					m[information[1]] = Bird{information[1],"worms", "fly","peep"}
					fmt.Println("Created it!","bird named",information[1])
				case "snake":
					m[information[1]] = Snake{information[1],"mice","slither","hsss"}
					fmt.Println("Created it!","snake named",information[1])
				default:
					fmt.Println("Please choose a proper animal type from the menu")
					break	
			}
		}


			if information[0] == "query" && len(m) > 0{
				val, ok := m[information[1]]
				_ = val
				if !(ok) {
					fmt.Println("That animal doesn't exist")
					continue
				} 
				inform(m[information[1]],information[2])
			} else if information[0] == "query" && len(m) == 0{
				fmt.Println("The list is empty please introduce a new animal first!! :)")
			}

			if information[0] != "query" && information[0] != "newanimal"{
					fmt.Println("Please choose only between 'neanimal' or 'query' comands")
			}




		}
	



}