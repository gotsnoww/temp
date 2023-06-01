package main

import(
	"fmt"
	"math/rand"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

var t int

func myrand(c chan int) {
	for i:= 0; i < 10; i++{
		v := rand.Intn(100)-50
		c <- v
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go myrand(c1)
	go myrand(c2)
	x := adder()
	y := adder()
	for i:= 0; i < 10; i++{
		v1 := <- c1
		v2 := <- c2
		fmt.Printf("1번 스레드의 %v번째 랜덤한 값 : %v\n", i, v1)
		fmt.Printf("2번 스레드의 %v번째 랜덤한 값 : %v\n", i, v2)
	
		fmt.Printf("1번 스레드의 총 합 : %v\n", x(v1))
		fmt.Printf("2번 스레드의 총 합 : %v\n", y(v2))
	}
}
