package iomanagement

import "fmt"

// Function for displaying main menu options
func ListMainMenuOptions() {

}

// Function for listing setting options
func ListSettingOptions() {

}

/* There are different functions for reading in answers because this allows us to parse the user
input differently depending on the type of problem. */

// Function for gathering an answer for an addition problem from the user.
func ReadAdditionAnswer(problem AdditionProblem) {

}

func ReadSubtractionAnswer(problem SubtractionProblem) {

}

func ReadMultiplicationAnswer(problem MultiplicationProblem) {

}

func ReadDivisionAnswer(problem DivisionProblem) {

}

func PrintGameOver(stats GameStats) {

}

// Function for collecting an integer value in a specified range.
func CollectIntegerInRange(lowerBound int, upperBound int) int {
	fmt.Println("Invalid Input, please type an integer from 1-3.")
}

// Function for collecting a bitmap from user input.
func CollectBitmap(options int) {

}
