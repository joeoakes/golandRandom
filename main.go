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
	fm.Println(time.Now().UnixNano())
	fm.Println(rand.Intn(100))
	fm.Println(rand.Intn(10))
}
