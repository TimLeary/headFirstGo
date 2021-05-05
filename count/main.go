package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/datafile"
	"log"
)

func main()  {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	votes := make(map[string]int)
	for _, line := range lines {
		votes[line]++
	}

	fmt.Println(votes)

	for candidate, vote := range votes {
		fmt.Println("Candidate:", candidate, "votes:", vote)
	}

}
