package gamemanagement

type GameSettings struct {
	GameLength     int
	GameDifficulty int
	MathTypes      int
}

type GameStats struct {
	ProblemsCorrect   int
	ProblemsIncorrect int
}

type GameManager struct {
	Settings GameSettings
	Stats    GameStats
}

func (manager GameManager) String() {

}
