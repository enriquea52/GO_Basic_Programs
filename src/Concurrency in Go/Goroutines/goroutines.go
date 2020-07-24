package main
import(
	"fmt"
	"sync"
)


func task1(num *int){
	*num = *num+1
	*num = *num+1
	*num = *num+1


}

func task2(num *int){
	fmt.Println(*num)

}



func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	var x int


	task1(&x)

	go func() {
		task2(&x)
		wg.Done()
	}()


	wg.Wait()
}

