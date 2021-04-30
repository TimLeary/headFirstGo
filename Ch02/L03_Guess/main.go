// guess challenges players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var tryings int = 10
var min int = 1
var max int = 100

func main()  {
	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	var target int = rand.Intn(max) + min

	for count := 0; count < tryings; count++ {
		guessed := getNumber(count)
		if guessed > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if guessed < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry. You didn't guessed it.")
	}

}

func getNumber(tried int) int {
	fmt.Println("Guess a number between", min, "and", max)
	fmt.Println("You have", tryings - tried, "guess.")
	reader := bufio.NewReader(os.Stdin)
	// Return everything the user has typed, up to where they pressed the Enter key.
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
