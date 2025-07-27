package main

import "fmt"

func genarator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("Generator Stage Emitting:", n)
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("Square Stage Received:", n)
			out <- n * n
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) int {
	total := 0
	for n := range in {
		fmt.Println("Sum Stage Received:", n)
		total += n
	}
	return total
}

func main() {
	nums := genarator(1, 2, 3, 4, 5)
	squared := square(nums)
	total := sum(squared)

	fmt.Println("Total:", total)
}
