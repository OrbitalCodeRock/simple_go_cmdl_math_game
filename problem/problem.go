package problem

func GenerateProblem(difficulty int32) Problem {

}

type Problem interface {
	InitializeProblem() int32
	GetDifficulty() int32
	String() string
	CheckAnswer(answer int32) bool
}

type AdditionProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem AdditionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 + problem.op2
}

func (problem AdditionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem AdditionProblem) String() string {

}

// Assumes that the answer string is of proper format.
// Proper format: A string containing just a 32-bit signed integer value.
func (problem AdditionProblem) CheckAnswer(answer int32) bool {

}

type SubtractionProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem SubtractionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 - problem.op2
}

func (problem SubtractionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem SubtractionProblem) String() string {

}

func (problem SubtractionProblem) CheckAnswer(answer string) bool {

}

type MultiplicationProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem MultiplicationProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 * problem.op2
}

func (problem MultiplicationProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem MultiplicationProblem) String() string {

}

func (problem MultiplicationProblem) CheckAnswer(answer string) bool {

}

type DivisionProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem DivisionProblem) InitializeProblem() {
	// Generate op1 and op2 using difficulty
	problem.answer = problem.op1 / problem.op2
}

func (problem DivisionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem DivisionProblem) String() string {

}

func (problem DivisionProblem) CheckAnswer(answer string) bool {

}
