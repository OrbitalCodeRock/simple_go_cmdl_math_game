package gamemanagement

type GameSettings struct {

	// The length of a game in minutes.
	GameLength int64

	// A value indicating the difficulty of the math problems generated.
	// Higher difficulties yield problems involving larger operands.
	GameDifficulty int32

	// A bit map representing all math types for the game.
	// 1st bit - Addition
	// 2nd bit - Subtraction
	// 3rd bit - Multiplication
	// 4th bit - Division
	// 0 digit - not present, 1 digit- present
	// Ex: 0b1111 - All Types present, 0b0001 - Addition only, 0b1000 - Division only
	MathTypes int32
}

type GameStats struct {
	ProblemsCorrect   int32
	ProblemsIncorrect int32
}

type GameManager struct {
	Settings GameSettings
	Stats    GameStats
}
