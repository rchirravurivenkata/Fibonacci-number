

package main


import (
    "fmt"
)
func calc_fib(n int) int{

	if n <= 1 {
		return n
	} else {
		return calc_fib(n-1) + calc_fib(n-2)
	}
}
func main() {

	var num int
    	fmt.Println("For what integer you need to calculate Fibonacci number?")
    	fmt.Scan(&num)

	fmt.Println(calc_fib(num))
}
