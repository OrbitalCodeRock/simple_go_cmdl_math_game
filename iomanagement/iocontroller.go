package iomanagement

import (
	"fmt"
	"math"
	"strconv"

	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/gamemanagement"
)

// Function for displaying main menu options
func ListMainMenuOptions() {
	var menuString string = `Please enter an integer corresponding to the following options:
	1) Start Game
	2) Change Settings
	3) Exit Game`

	fmt.Println(menuString)
}

// Function for listing setting options
func ListSettingOptions() {
	var settingsString string = `Please enter an integer corresponding to the following options:
	1) Set game length
	2) Set game difficulty
	3) Set math types
	4) Back`

	fmt.Println(settingsString)
}

/* There are different functions for reading in answers because this allows us to parse the user
input differently depending on the type of problem. */

// For now, these are essentially identical.

// Helper function for parsing an int32 from a string. This function does not handle errors, but instead, returns the error
// information to be handeled by the caller.
func parseInt32() (int32, *strconv.NumError) {
	// Read the answer from standard input.
	// Answer should be a single integer value.
	var answer string
	var answerInt int32
	fmt.Scan(&answer)

	// Attempt to parse the answer to an Integer
	ans, err := strconv.ParseInt(answer, 0, 32)

	answerInt = int32(ans)

	numErr := err.(*strconv.NumError)

	return answerInt, numErr
}

func parseInt32Base2() (int32, *strconv.NumError) {
	// Read the answer from standard input.
	// Answer should be a single integer value.
	var answer string
	var answerInt int32
	fmt.Scan(&answer)

	// Attempt to parse the answer to an Integer
	ans, err := strconv.ParseInt(answer, 2, 32)

	answerInt = int32(ans)

	numErr := err.(*strconv.NumError)

	return answerInt, numErr
}

// Returns true if an error is not present, false otherwise. Allows you to print error messages for each error case.
func handleNumError(numErr strconv.NumError, syntaxErrorMessage string, rangeErrorMessage string) bool {
	switch numErr.Err {
	case strconv.ErrSyntax:
		fmt.Println(syntaxErrorMessage)
		return false
	case strconv.ErrRange:
		fmt.Println(rangeErrorMessage)
		return false
	default:
		// In this case, there was no error.
		return true
	}
}

// Helper function for repeatedly asking for user input until a valid int32 value is accessed.
// The start message is printed once before the start of the loop
func readAnswerAsInt32(startMessage string, syntaxErrorMessage string, rangeErrorMessage string) int32 {
	fmt.Println(startMessage)
	for {
		answerInt, numErr := parseInt32()
		if handleNumError(*numErr, syntaxErrorMessage, rangeErrorMessage) {
			return answerInt
		}
	}
}

// Helper function for repeatedly asking for user input until a valid int32 value is accessed.
// The value will be interpreted in base 2.
// The start message is printed once before the start of the loop
func readAnswerAsInt32Base2(startMessage string, syntaxErrorMessage string, rangeErrorMessage string) int32 {
	fmt.Println(startMessage)
	for {
		answerInt, numErr := parseInt32Base2()
		if handleNumError(*numErr, syntaxErrorMessage, rangeErrorMessage) {
			return answerInt
		}
	}
}

var startMessage string = "Please enter an integer value: "
var syntaxErrorMessage string = "The answer format was invalid. It was either empty or contaned invalid characters. Please try again."
var rangeErrorMessage string = "The answer was too large or too small. Please type an integer between -2,147,483,648 and 2,147,483,647"

// Function for gathering an answer for an addition problem from the user.
func ReadAdditionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

func ReadSubtractionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

func ReadMultiplicationAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

func ReadDivisionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

func PrintGameOver(stats gamemanagement.GameStats) {
	var gameOverString string = fmt.Sprintf("Game Over!\nYou solved %d problems correctly\nYou answered %d problems incorrectly\n", stats.ProblemsCorrect, stats.ProblemsIncorrect)
	fmt.Print(gameOverString)
}

// Function for collecting an integer value in a specified range.
func CollectIntegerInRange(lowerBound int32, upperBound int32) int32 {
	var startMessage string = fmt.Sprintf("Please type an integer from %d to %d.\n", lowerBound, upperBound)
	var rangeErrorMessage string = fmt.Sprintf("The answer was too large or too small. Please type an integer between %d and %d\n", lowerBound, upperBound)
	answerInt := readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)

	for answerInt < lowerBound || answerInt > upperBound {
		fmt.Println(rangeErrorMessage)
		answerInt = readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
	}

	return answerInt

}

// Function for collecting a bitmap from user input.
func CollectBitmap(options int32) int32 {

	var answerLimit int32 = int32(math.Pow(2, float64(options)) - 1)
	var answerInt int32 = 0
	var startMessage string = ""
	var rangeErrorMessage string = fmt.Sprintf("The answer was too large or too small. Please type an integer between %b and %b\n", 0, answerLimit)
	answerInt = readAnswerAsInt32Base2(startMessage, syntaxErrorMessage, rangeErrorMessage)

	for answerInt <= 0 || answerInt > answerLimit {
		fmt.Println(rangeErrorMessage)
		answerInt = readAnswerAsInt32Base2(startMessage, syntaxErrorMessage, rangeErrorMessage)
	}

	return answerInt
}
