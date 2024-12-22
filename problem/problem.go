package problem

func GenerateProblem(difficulty int) Problem {

}

type Problem interface {
	InitializeProblem() int
	GetDifficulty() int
	String() string
	CheckAnswer(answer string) bool
}

type AdditionProblem struct {
	op1, op2, answer, difficulty int
}

func (problem AdditionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 + problem.op2
}

func (problem AdditionProblem) GetDifficulty() int {
	return problem.difficulty
}

func (problem AdditionProblem) String() string {

}

func (problem AdditionProblem) CheckAnswer(answer string) bool {

}

type SubtractionProblem struct {
	op1, op2, answer, difficulty int
}

func (problem SubtractionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 - problem.op2
}

func (problem SubtractionProblem) GetDifficulty() int {
	return problem.difficulty
}

func (problem SubtractionProblem) String() string {

}

func (problem SubtractionProblem) CheckAnswer(answer string) bool {

}

type MultiplicationProblem struct {
	op1, op2, answer, difficulty int
}

func (problem MultiplicationProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 * problem.op2
}

func (problem MultiplicationProblem) GetDifficulty() int {
	return problem.difficulty
}

func (problem MultiplicationProblem) String() string {

}

func (problem MultiplicationProblem) CheckAnswer(answer string) bool {

}

type DivisionProblem struct {
	op1, op2, answer, difficulty int
}

func (problem DivisionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 / problem.op2
}

func (problem DivisionProblem) GetDifficulty() int {
	return problem.difficulty
}

func (problem DivisionProblem) String() string {

}

func (problem DivisionProblem) CheckAnswer(answer string) bool {

}
