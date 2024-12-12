package main

import (
	"fmt"

	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/gamemanagement"
	"github.com/OrbitalCodeRock/simple_go_cmdl_math_game/iomanagement"
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
		var option int = iomanagement.CollectIntegerInRange(1, 3)
		switch option {
		case 1:
			runGame()
		case 2:
			changeSettings(&gameManager)
		case 3:
			keepRunning = false
		}

	}
}

// Private method for running the game.
func runGame() {

}

// private method for changing the settings of the game.
func changeSettings(gameManager *gamemanagement.GameManager) {
	var keepSetting bool = true
	for keepSetting {
		iomanagement.ListSettingOptions()
		var option int = iomanagement.CollectIntegerInRange(1, 4)
		switch option {
		case 1:
			fmt.Println("Please enter the # of minutes you would like to play, from 1 to 10")
			var minutes int = iomanagement.CollectIntegerInRange(1, 10)
			gameManager.Settings.GameLength = minutes
		case 2:
		case 3:
		case 4:
		}
	}
}
