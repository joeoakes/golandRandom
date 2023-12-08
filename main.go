package main

//Packages
import (
	fm "fmt"
	"math/rand"
	"time"
)

// Prints a random number zero to 99
func main() {
	rand.Seed(time.Now().UnixNano())
	fm.Println("Unix time now:", time.Now().UnixNano())
	fm.Println("Get a random number 0-99:", rand.Intn(100))
	fm.Println("Get a random number 0-9:", getNumbers(10))

	// Define the range for the random numbers
	min := 1
	max := 10

	// Generate and print 10 random numbers
	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(max-min+1) + min
		fm.Println(randomNum)
	}
}

// camelCasing with parameter_list
func getNumbers(ten int) int {
	return rand.Intn(ten)
}
