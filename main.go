package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to my game!\n", name)
	fmt.Printf("Enter your age: \n")
	var age int
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	numOfQuestions := 2

	fmt.Println("Game Started!")
	fmt.Printf("What's better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct")
		score += 1
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct")
		score += 1
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("How many cores does a Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("You scored %v out of %v.", score, numOfQuestions)
	percent := (float64(score) / float64(numOfQuestions)) * 100
	fmt.Printf("You scored: %v%%", percent)
}
