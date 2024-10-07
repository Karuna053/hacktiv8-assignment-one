package main

import (
	"assignment-one/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {
	var programLocation string = os.Args[0]
	var firstArgument string = "[No First Argument Entered]"

	if len(os.Args) > 1 {
		firstArgument = os.Args[1]

		var person = helpers.Person{}
		if firstArgument == "Thomas" {
			person.ID = 0
			person.Name = "Thomas"
			person.Address = "Kota A"
			person.Job = "progrmamer"
			person.Reason = "Alasan Thomas"
		}

		fmt.Println(programLocation, firstArgument)
		fmt.Println(strings.Repeat("#", 30))

		fmt.Println("ID: ", person.ID)
		fmt.Println("Name: ", person.Name)
		fmt.Println("Address: ", person.Address)
		fmt.Println("Job: ", person.Job)
		fmt.Println("Reason: ", person.Reason)
	}

}
