package main
import(
	"fmt"
	"sync"
)

type ChopS struct{sync.Mutex}

type Philo struct {
	leftCS, rightCS *ChopS
	id int
}

func (p Philo) eat(c chan int) {

	for i:=0;i<3;i++ {
		

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("start to eat",p.id+1)
		fmt.Println("finish eating",p.id+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
		c<-1


	}

}




func main() {
	c := make(chan int)

	CSticks := make([]*ChopS,5)

	for i:=0;i<5;i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo,5)
	for i:=0;i<5;i++ {
		philos[i] = &Philo{CSticks[i],CSticks[(i+1)%5],i}
	}


	for i:=0;i<5;i++ {
		go philos[i].eat(c)
	}


	for i:=0;i<15;i++ {
		<-c
	}

	fmt.Println("Great Finalle")

}