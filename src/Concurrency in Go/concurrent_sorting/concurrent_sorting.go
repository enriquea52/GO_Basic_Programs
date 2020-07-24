/*
This program gets builds an array with user input and divides it into 4 sub arrays (slices) and sorts each slice concurrently
so at the end the program it merges the sorted slices and sorts the full array.
*/


/*
sample slices for sorting and testing
slice:=[]int{5,2,6,2,1,3}
slice:=[]int{5,2,6,2,1,3,3,8,9}
slice:=[]int{5,2,6,2,1,3,3,8,9,6,5,2,4,6,3,6,3,1,3,2,4,6,7,8,6,4}
slice:=[]int{4,3,2,1,7,5,8,5}
slice:=[]int{5,2,6,2,3,2,4,5,6,7,8,9}
*/


package main
import(
	"fmt"
	"sort"
	"strconv"
)
func main() {

slice:=[]int{}
var number string

for { //Creating slice from user input
	fmt.Println("Please enter an integer to add to your array to be sorted")
	fmt.Println("enter x to sort your array")
	fmt.Scan(&number)
	if number == "x" ||number == "X"{
		break
	}
	s,err:=strconv.Atoi(number)
	_=err
	slice = append(slice,s)
	fmt.Println("Your array is: ",slice," until now")

}


c := make(chan int) //Creating a channel to wait for processes to complete
size := len(slice)
threads:= 4
step:=size/threads
last:=size%threads
fmt.Println("Your slice: ",slice)
sorter :=invoke() //Creating function with identifier m

	if size < threads { //In case there are not enough numbers to make each thread work, exit the program 
		fmt.Println("Please define an array of more than 3 numbers, till next time bro :)")
		return
	}

	//In this section the threads are created depending on the nature (number of elements) of the slice to be sorted
	if last == 0 { 
		for i:=0;i<threads;i++ {
			go sorter(slice,i*step,i*step+step,c)
			//fmt.Println(slice[i*step:i*step+step])
		}
	} else {
		for i:=0;i<threads;i++{
			if i == threads-1{	
				go sorter(slice,i*step,i*step+step+last,c)
				//fmt.Println(slice[i*step:i*step+step+last])
			}else  {
				go sorter(slice,i*step,i*step+step,c)
				//fmt.Println(slice[i*step:i*step+step])
			}

		}
	}


	for i:=0;i<threads;i++{ //Waiting for threads to finish their instructions
		<-c
	}


	fmt.Println("Sub-sorted slice ",slice) //Printing array with each sorted slice but not array not fully sorted
	sort.Ints(slice)
	fmt.Println("Sorted slice ",slice) //Array fully sorted

}

//Sorting function, function is created by another function to keep track of which instance is doing what.
func invoke() func (sli []int,init int,fin int,c chan int) {
	m:=0
	fn:= func(sli []int,init int,fin int,c chan int) {
		
		sort.Ints(sli[init:fin])
		m++
		fmt.Println("Process:",m,"...",sli[init:fin]," sorted")

		c<-1

	}
	return fn

}





