package main

import (
	"fmt"
	"time"

	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/gamemanagement"
	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/iomanagement"
	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/problem"
)

func main() {
	// Set default game settings
	var gameManager gamemanagement.GameManager
	gameManager.Settings = gamemanagement.GameSettings{
		GameLength:     1,
		GameDifficulty: 1,
		MathTypes:      0b1111,
	}

	var keepRunning bool = true
	for keepRunning {
		// List current settings here
		iomanagement.ListMainMenuOptions()
		var option int32 = iomanagement.CollectIntegerInRange(1, 3)
		switch option {
		case 1:
			runGame(&gameManager)
		case 2:
			changeSettings(&gameManager)
		case 3:
			keepRunning = false
		}

	}
}

// Private method for running the game.
func runGame(gameManager *gamemanagement.GameManager) {
	timeUp := false
	go startGameTimer(gameManager, 2, &timeUp)
	for !timeUp {
		mathProblem := problem.GenerateProblem(gameManager.Settings.GameDifficulty)
		fmt.Println(mathProblem)
		var answer int32
		switch mathProblem.(type) {
		case problem.AdditionProblem:
			answer = iomanagement.ReadAdditionAnswer()
		case problem.SubtractionProblem:
			answer = iomanagement.ReadSubtractionAnswer()
		case problem.MultiplicationProblem:
			answer = iomanagement.ReadMultiplicationAnswer()
		case problem.DivisionProblem:
			answer = iomanagement.ReadDivisionAnswer()
		}
		var answerCorrect bool = mathProblem.CheckAnswer(answer)
		if answerCorrect {
			gameManager.Stats.ProblemsCorrect++
			fmt.Println("That's correct!")
		} else {
			gameManager.Stats.ProblemsIncorrect--
			fmt.Println("That answer was incorrect.")
		}
	}
	iomanagement.PrintGameOver(gameManager.Stats)

}

// Method that keeps track of how much time is left in the game and prints warnings for the player.
// gameManager - Game Manager object used to derive the total length of the game.
// warningFrequency - How often, in minutes, that warnings should be outputted to the user.
// timeUp - Pointer to a boolean that is set to true once time has run out.
func startGameTimer(gameManager *gamemanagement.GameManager, warningFrequency int32, timeUp *bool) {
	ticker := time.NewTicker(time.Duration(gameManager.Settings.GameLength) * time.Minute)
	var minutes int32 = 0
	for minutesLeft := gameManager.Settings.GameLength; minutesLeft > 0; minutesLeft-- {
		<-ticker.C
		if minutes%warningFrequency == 0 {
			fmt.Printf("%d minutes remain!", minutesLeft)
		}
		minutes++
	}
	*timeUp = true
}

// private method for changing the settings of the game.
// gameManager - GameManager object pointer that refers to the GameManager holding settings to be changed.
func changeSettings(gameManager *gamemanagement.GameManager) {
	var keepSetting bool = true
	for keepSetting {
		iomanagement.ListSettingOptions()
		var option int32 = iomanagement.CollectIntegerInRange(1, 4)
		switch option {
		case 1:
			fmt.Println("Please enter the # of minutes you would like to play, from 1 to 10.")
			var minutes int64 = int64(iomanagement.CollectIntegerInRange(1, 10))
			gameManager.Settings.GameLength = minutes
		case 2:
			fmt.Println("Please enter a difficulty level, from 1 to 10.")
			var difficulty int32 = iomanagement.CollectIntegerInRange(1, 10)
			gameManager.Settings.GameDifficulty = difficulty
		case 3:
			var mathTypeOptions string = `Please choose the math types of the game. Enter a non-seperated 4-character sequence of 0s and 1s.
			1: Present, 0: Not Present, 1st Digit: Addition, 2nd Digit: Subtraction, 3rd Digit: Multiplication, 4th Digit: Division. 
			Examples: "1000" - Addition only, "1010" - Addition and Multiplication`
			fmt.Println(mathTypeOptions)
			var mathTypes int32 = iomanagement.CollectBitmap(4)
			gameManager.Settings.MathTypes = mathTypes
		case 4:
			keepSetting = false
		}
	}
}
