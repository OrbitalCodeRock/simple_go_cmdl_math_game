package iomanagement

import (
	"fmt"
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

// Helper method for parsing an integer value from user input.
// Returns the parsed integer.
// Returns error information. Indicates whether or not there is no error, if there is a syntax error, or if there is a range error.
// This error information is retrieved from strconv.ParseInt
func parseInt32() (int32, *strconv.NumError) {
	// Read the answer from standard input.
	// Answer should be a single integer value.
	var answer string
	var answerInt int32
	fmt.Scan(&answer)

	// Attempt to parse the answer to an Integer
	ans, err := strconv.ParseInt(answer, 0, 32)

	answerInt = int32(ans)

	var numErr *strconv.NumError = nil
	if err != nil {
		numErr = err.(*strconv.NumError)
	}

	return answerInt, numErr
}

// Helper method for parsing an integer value from user input. The string is intepreted as a base 2 value.
// Returns the parsed integer.
// Returns error information. Indicates whether or not there is no error, if there is a syntax error, or if there is a range error.
// This error information is retrieved from strconv.ParseInt
func parseInt32Base2() (int32, *strconv.NumError) {
	// Read the answer from standard input.
	// Answer should be a single integer value.
	var answer string
	var answerInt int32
	fmt.Scan(&answer)

	// Attempt to parse the answer to an Integer
	ans, err := strconv.ParseInt(answer, 2, 32)

	answerInt = int32(ans)

	var numErr *strconv.NumError = nil
	if err != nil {
		numErr = err.(*strconv.NumError)
	}

	return answerInt, numErr
}

// Returns true if an error is not present, false otherwise. Allows you to print error messages for each error case.
// numErr - Contains information surrounding the type of error detected from strconv.ParseInt. This is nil if there is no error.
// syntaxErrorMessage - A message to be printed when the user input cannot be parsed as an integer.
// rangeErrorMessage - A message to be printed when the user input is parsed as an integer that is too large or too small.
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
// startMessage - A message to be printed before reading any user input. This message is only printed once.
// syntaxErrorMessage - A message to be printed when the user input cannot be parsed as an integer.
// rangeErrorMessage - A message to be printed when the user input is parsed as an integer that is too large or too small.
func readAnswerAsInt32(startMessage string, syntaxErrorMessage string, rangeErrorMessage string) int32 {
	fmt.Println(startMessage)
	for {
		answerInt, numErr := parseInt32()
		// If there isn't an error, return the parsed answer.
		if numErr == nil || handleNumError(*numErr, syntaxErrorMessage, rangeErrorMessage) {
			return answerInt
		}
	}
}

// Helper function for repeatedly asking for user input until a valid int32 value is retrieved.
// The value will be interpreted in base 2.
// startMessage - A message to be printed before reading any user input. This message is only printed once.
// syntaxErrorMessage - A message to be printed when the user input cannot be parsed as an integer.
// rangeErrorMessage - A message to be printed when the user input is parsed as an integer that is too large or too small.
func readAnswerAsInt32Base2(startMessage string, syntaxErrorMessage string, rangeErrorMessage string) int32 {
	fmt.Println(startMessage)
	for {
		answerInt, numErr := parseInt32Base2()
		if numErr == nil || handleNumError(*numErr, syntaxErrorMessage, rangeErrorMessage) {
			return answerInt
		}
	}
}

// Default messages for the functions below.
var startMessage string = "Please enter an integer value: "
var syntaxErrorMessage string = "The answer format was invalid. It was either empty or contaned invalid characters. Please try again."
var rangeErrorMessage string = "The answer was too large or too small. Please type an integer between -2,147,483,648 and 2,147,483,647"

// Function for gathering an answer for an addition problem from the user.
// Returns the read answer as an integer.
func ReadAdditionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

// Function for reading a subtraction problem answer from the user.
// Returns the read answer as an integer.
func ReadSubtractionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

// Function for reading a multiplication problem answer from the user.
// Returns the read answer as an integer.
func ReadMultiplicationAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

// Function for reading a division problem answer from the user.
// Returns the read answer as an integer.
func ReadDivisionAnswer() int32 {
	return readAnswerAsInt32(startMessage, syntaxErrorMessage, rangeErrorMessage)
}

// Prints a game over message tailored to game statistics.
// stats - The stats struct containing the statistics for the game that just ended.
func PrintGameOver(stats gamemanagement.GameStats) {
	var gameOverString string = fmt.Sprintf("Game Over!\nYou solved %d problems correctly\nYou answered %d problems incorrectly\n", stats.ProblemsCorrect, stats.ProblemsIncorrect)
	fmt.Print(gameOverString)
}

// Function for collecting an integer value from the user in a specified range.
// lowerBound - Inclusive lower bound on the integer value to collect.
// upperBound - Inclusive upper bound on the integer value to collect.
// Returns the collected integer.
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
// options - Indicates the number of bits in the bitmap, and therefore, the number of options toggled by the bitmap.
// E.g, 0b1111 - 4 options, 0b1010 - 4 options, 0b10101 - 5 options
// Returns the bitmap as an integer.
func CollectBitmap(options int32) int32 {

	var answerLimit int32 = (1 << options) - 1
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
