package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/gamemanagement"
	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/iomanagement"
	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/problem"
)

func main() {
	// Seed random number generator
	var seed = time.Now().UnixNano()
	numSource := rand.NewSource(seed)

	// Set default game settings
	var gameManager gamemanagement.GameManager
	gameManager.Settings = gamemanagement.GameSettings{
		GameLength:     1,
		GameDifficulty: 1,
		MathTypes:      0b1111,
	}

	var keepRunning bool = true
	for keepRunning {
		iomanagement.ListMainMenuOptions()
		var option int32 = iomanagement.CollectIntegerInRange(1, 3)
		switch option {
		case 1:
			runGame(&gameManager, &numSource)
		case 2:
			changeSettings(&gameManager)
		case 3:
			keepRunning = false
		}

	}
}

// Private method for running the game.
// gameManager - References a struct containing game settings and statistics
// numSource - References a random number generator that generates int64 values.
func runGame(gameManager *gamemanagement.GameManager, numSource *rand.Source) {
	timeUp := false
	go startGameTimer(gameManager, 2, &timeUp)
	for !timeUp {
		mathProblem := problem.GenerateProblem(gameManager.Settings.GameDifficulty, gameManager.Settings.MathTypes, numSource)
		fmt.Println(mathProblem)
		var answer int32
		switch mathProblem.(type) {
		case *problem.AdditionProblem:
			answer = iomanagement.ReadAdditionAnswer()
		case *problem.SubtractionProblem:
			answer = iomanagement.ReadSubtractionAnswer()
		case *problem.MultiplicationProblem:
			answer = iomanagement.ReadMultiplicationAnswer()
		case *problem.DivisionProblem:
			answer = iomanagement.ReadDivisionAnswer()
		}
		var answerCorrect bool = mathProblem.CheckAnswer(answer)
		if answerCorrect {
			gameManager.Stats.ProblemsCorrect++
			fmt.Println("That's correct!")
		} else {
			gameManager.Stats.ProblemsIncorrect++
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
	ticker := time.NewTicker(time.Minute)
	var minutes int32 = 1
	for minutesLeft := gameManager.Settings.GameLength; minutesLeft > 0; minutesLeft-- {
		<-ticker.C
		if minutes%warningFrequency == 0 {
			switch minutesLeft {
			case 1:
				fmt.Printf("\n%d minute remains!\n", minutesLeft)
			default:
				fmt.Printf("\n%d minutes remain!\n", minutesLeft)
			}

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
			Examples: "0001" - Addition only, "0101" - Addition and Multiplication`
			fmt.Println(mathTypeOptions)
			gameManager.Settings.MathTypes = iomanagement.CollectBitmap(4)
		case 4:
			keepSetting = false
		}
	}
}
