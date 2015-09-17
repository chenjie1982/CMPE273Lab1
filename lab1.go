package main

import ("fmt"
	"math"
	"time"
)

/*func f(n int){
	for i:=0; i<10;i++{
		fmt.Println(n,":",i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
}*/


func fib(x  int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	return fib(x-1)+fib(x-2)
}

/*func fibDP(x int) int{
	sum := make([]int, x+1)
	sum[0] = 0
	sum[1] = 1
	for i:= 2; i < x+1; i++ {
		sum[i] = sum[i-1] + sum[i-2]
	}
	return sum[x]
}*/

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l+2*w
}

func (r *Circle) perimeter() float64 {

	return 2*math.Pi*r.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func Sleep(x time.Duration) {
	<-time.After(x)
}

func main(){

	var input int

	fmt.Print("Enter a number: ")
	//var input float64
	fmt.Scanln(&input)
	fmt.Println("The ",input," fibonacci number is ", fib(input))
	fmt.Print("Enter the sleep time(second): ")
	fmt.Scanln(&input)
	Sleep(time.Second * time.Duration(input))
	fmt.Print("Already slept ",input," second!")
	fmt.Print(time.Now())

	/*output := input*2
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	fmt.Println(output)*/
}
