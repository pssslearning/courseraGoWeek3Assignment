// --------------------------------------------------------------
// Source code mantained at Github repository for Learning
// https://github.com/pssslearning/courseraGoWeek2Assignment
// --------------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"strconv"
)

const DASHES = "--------------------------------------------------------------"

func showUSerPrompt() {
	fmt.Println(DASHES)
	fmt.Println("Please enter a string that can be interpreted as INTEGER ... ")
	fmt.Println("      (press <CTRL+C> or 'X' to exit.) ")
	fmt.Println(DASHES)
}

func main() {
	// Show Welcome message at start
	fmt.Println("\nWelcome user to the Assignment program for Week 3,\n")

	var aRuneInInput rune
	var aString string

	// The "empty integer slice of size (length) 3"
	contentSlice := make([]int, 0, 3)

	for {

		aString = ""

		showUSerPrompt()

		for {
			_, err := fmt.Scanf("%c", &aRuneInInput)

			if err != nil {
				fmt.Println("An error has happenned.")
				fmt.Printf("The error detected was [%s]\n", err)
				fmt.Println("The program will be exited.")
				fmt.Println("Please note that may be remaining input not consumed and sent to your input. ")
				os.Exit(1)
			}

			if aRuneInInput == rune('\n') {
				break
			}

			aString += string(aRuneInInput)
		}

		// Exit program condition
		if aString == "X" {
			// Exit loop
			break
		}

		if n, err := strconv.Atoi(aString); err == nil {
			contentSlice = append(contentSlice, n)
			sortAscending(contentSlice)
			showSliceContent(contentSlice)
		} else {
			fmt.Printf("ENTRY REJECTED [%s]. Cannot be parsed as an integer.\n", aString)
		}

	}

	// Exit with return code 0 to OS.
	fmt.Println("\nGoodbye !!!\n")
	os.Exit(0)
}

func showSliceContent(theSlice []int) {
	fmt.Println(DASHES)
	fmt.Println("The ordered list of integers introduced by now are:")
	for i, n := range theSlice {
		fmt.Printf("Position #%d - Value[%d]\n", i, n)
	}
}

func sortAscending(theSlice []int) {

	// To preven an infinite loop is
	// the sorting algorithm contained
	// a bug
	infiniteLoopProtector := 0

	isOrdered := false
	for !isOrdered {

		aSwapOccurred := false
		for i := range theSlice {

			if infiniteLoopProtector++; infiniteLoopProtector > len(theSlice)*100 {
				fmt.Println("INFINITE LOOP HAPPENED. EXITING ...")
				os.Exit(-1)
			}

			if j := i + 1; j < len(theSlice) {
				if theSlice[i] > theSlice[j] {
					theSlice[i], theSlice[j] = theSlice[j], theSlice[i]
					aSwapOccurred = true
				}
			} else {
				if aSwapOccurred {
					isOrdered = false
				} else {
					isOrdered = true
				}
				break
			}
		}
	}
}
