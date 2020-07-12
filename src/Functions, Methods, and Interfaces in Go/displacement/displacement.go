/*
This program shows an implementation of the dyanmic use of function by creating functions with different parameters taking
as an example a physics scenario by using a formula to calculate displacement given time, acceleration, initial velocity and 
initial displacement. The initial parameters are entered by the user.
*/
package main
import(
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn:= func(t float64) float64{
		return (0.5*a*t*t)+(v0*t)+s0
	}

	return fn
}

func main() {
	var x string
	param:=[]float64{0,0,0}
	param_names := []string{"Acceleration","Initial Velocity","Initial displacement"}

	for i:=0;i<3;i++{
		fmt.Println("Please Introduce the ",param_names[i],"parameter")
		fmt.Scan(&x)
		k,err := strconv.ParseFloat(x, 64)
		_=err
		param[i] = k
	}
	fmt.Println("Acceleration: ",param[0],"\nInitial Velocity: ",param[1],"\nInitial Displacement: ",param[2],"\n")

	fn := GenDisplaceFn(param[0],param[1],param[2]) 
	fmt.Println("Displacement for time = 1 is: ",fn(1))
	fmt.Println("Displacement for time = 5 is: ",fn(5))


}
