package main
import(
	"fmt"
	"encoding/json"

)

func main(){
	person := make(map[string]string)
	name:=""
	address:=""

	person["name"] = ""
	person["address"] = ""

	fmt.Println("Please Enter Yur name")
	fmt.Scan(&name)
	fmt.Println("Please enter your address")
	fmt.Scan(&address)

	person["name"] = name
	person["address"] = address

	barr,err:=json.Marshal(person)

	if err == nil {
		fmt.Println("")
		fmt.Println("The json object is: ")
		fmt.Println(string(barr))
		//json.Unmarshal(barr,&person)
		//fmt.Println(person)

	} else {
		fmt.Println("Error")
	}

	

}