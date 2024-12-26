package gamemanagement

type GameSettings struct {
	GameLength     int32
	GameDifficulty int32
	MathTypes      int32
}

type GameStats struct {
	ProblemsCorrect   int32
	ProblemsIncorrect int32
}

type GameManager struct {
	Settings GameSettings
	Stats    GameStats
}

func (manager GameManager) String() {

}
