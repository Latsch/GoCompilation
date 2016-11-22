//Collatz Conjecture - Start with a number n > 1.
//Find the number of steps it takes to reach one using the following process: If n is even, divide it by 2.
//If n is odd, multiply it by 3 and add 1.

package main

import "fmt"

var count int = 0
var n int

func main(){

	fmt.Println("Input a positive number: ")
	
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	collatz(n)
}

func collatz(n int){

	fmt.Printf("Collatz Number of %d is : ", n)
	fmt.Printf("%d", collatzCalc(n))
}

func collatzCalc(n int) int{
	
	
	if n > 1 {
		if n%2 == 0 {
			n = n /2
			count++
			collatzCalc(n)
		} else {
			n = n *3 + 1
			count++
			collatzCalc(n)	
		}		
	} else if count == 0 {
		fmt.Printf("Zahl muss groesser 1 sein\n")
	}
	
	return count
}

