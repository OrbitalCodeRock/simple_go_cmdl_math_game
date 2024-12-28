package problem

import (
	"fmt"
	"math"
	"math/rand"
)

const maxSourceNumber int64 = (1 << 63) - 1

// Bit strings where each bit represents a math type.
// These specific strings represent a singular type.
// Addition, Subtraction, Multiplication, and Division, respectively.
var typeMaps = [4]int32{0b0001, 0b0010, 0b0100, 0b1000}

// Function for generating a problem that is one of the allowed problem types.
// Randomly chooses the problem type, giving an equal weight to all allowed types.
// Initiallizes the problem for the caller.
// difficulty - Indicates the difficulty of the problem to generate.
// allowedTypes - Bitmap holding all of the allowed math types for the problem to generate.
// numSource - A random number generator for generating integers from 0 to 2^62
func GenerateProblem(difficulty int32, allowedTypes int32, numSource *rand.Source) Problem {

	var numAllowedTypes int32 = 0
	var typeArray [4]Problem

	// Loop for figuring out how many types of math are allowed, based on the allowedTypes bitmap.
	// Also initializes an array that allows us to assign even probabilities of choosing a given allowed type.
	for _, mathType := range typeMaps {
		if allowedTypes&mathType == mathType {
			// Switch should have a case for every bitmap in typeMaps
			switch mathType {
			case 0b0001:
				typeArray[numAllowedTypes] = &AdditionProblem{op1: 0, op2: 0, answer: 0, difficulty: 0}
			case 0b0010:
				typeArray[numAllowedTypes] = &SubtractionProblem{op1: 0, op2: 0, answer: 0, difficulty: 0}
			case 0b0100:
				typeArray[numAllowedTypes] = &MultiplicationProblem{op1: 0, op2: 0, answer: 0, difficulty: 0}
			case 0b1000:
				typeArray[numAllowedTypes] = &DivisionProblem{op1: 0, op2: 0, answer: 0, difficulty: 0}
			}
			numAllowedTypes++
		}
	}

	// Bucket size is the maximum random number that can be generated divided by the number of allowed types.
	// A "bucket" in this case refers to the set of numbers that correspond to an allowed type.
	var bucketSize int64 = maxSourceNumber / int64(numAllowedTypes)

	// Type Decider will be a value ranging from 0 to (numAllowedTypes - 1).
	var typeDecider int64 = ((*numSource).Int63() - 1) / bucketSize

	var generatedProblem Problem = typeArray[typeDecider]
	generatedProblem.InitializeProblem(difficulty, numSource)
	return generatedProblem

}

type Problem interface {
	// Function for initializing Problem properties.
	InitializeProblem(int32, *rand.Source)

	// Function for getting the difficulty of the problem.
	GetDifficulty() int32

	// Function for getting the string representation of the problem.
	String() string

	// Function for determining if a provided answer is correct.
	CheckAnswer(answer int32) bool
}

type AdditionProblem struct {
	op1, op2, answer, difficulty int32
}

// Function for generating a random operand based on a difficulty value.
// The maximum number of operand digits = difficulty/2
// The minimum number of operand digits = difficulty/4
func generateRandomOperand(difficulty int32, numSource *rand.Source) int32 {
	var maxOpDigits int32 = difficulty / 2
	if difficulty < 2 {
		maxOpDigits = 1
	}
	var minOpDigits int32 = maxOpDigits / 2
	if difficulty < 4 {
		minOpDigits = 1
	}

	var maxNum = int32(math.Pow10(int(maxOpDigits))) - 1
	var minNum = int32(math.Pow10(int(minOpDigits-1))) - 1

	randomOp := (*numSource).Int63()
	for randomOp > int64(maxNum) {
		randomOp /= 10
	}

	for randomOp < int64(minNum) {
		randomOp *= 10
	}

	return int32(randomOp)

}

func (problem *AdditionProblem) InitializeProblem(difficulty int32, numSource *rand.Source) {
	problem.difficulty = difficulty
	problem.op1 = generateRandomOperand(difficulty, numSource)
	problem.op2 = generateRandomOperand(difficulty, numSource)
	problem.answer = problem.op1 + problem.op2
}

func (problem *AdditionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem *AdditionProblem) String() string {
	return fmt.Sprintf("Addition Problem: %d + %d = ?\n", problem.op1, problem.op2)
}

// Assumes that the answer string is of proper format.
// Proper format: A string containing just a 32-bit signed integer value.
func (problem *AdditionProblem) CheckAnswer(answer int32) bool {
	return answer == problem.answer
}

type SubtractionProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem *SubtractionProblem) InitializeProblem(difficulty int32, numSource *rand.Source) {
	problem.difficulty = difficulty
	problem.op1 = generateRandomOperand(difficulty, numSource)
	problem.op2 = generateRandomOperand(difficulty, numSource)
	problem.answer = problem.op1 - problem.op2
}

func (problem *SubtractionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem *SubtractionProblem) String() string {
	return fmt.Sprintf("Subtraction Problem: %d - %d = ?\n", problem.op1, problem.op2)
}

func (problem *SubtractionProblem) CheckAnswer(answer int32) bool {
	return answer == problem.answer
}

type MultiplicationProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem *MultiplicationProblem) InitializeProblem(difficulty int32, numSource *rand.Source) {
	problem.difficulty = difficulty
	// Since multiplication problems are significantly more difficult than addition or subtraction problems, we decrease the difficulty for these problems.
	problem.op1 = generateRandomOperand(difficulty-difficulty/2, numSource)
	problem.op2 = generateRandomOperand(difficulty-difficulty/2, numSource)
	problem.answer = problem.op1 * problem.op2
}

func (problem *MultiplicationProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem *MultiplicationProblem) String() string {
	return fmt.Sprintf("Multiplication Problem: %d * %d = ?\n", problem.op1, problem.op2)
}

func (problem *MultiplicationProblem) CheckAnswer(answer int32) bool {
	return answer == problem.answer
}

type DivisionProblem struct {
	op1, op2, answer, difficulty int32
}

func (problem *DivisionProblem) InitializeProblem(difficulty int32, numSource *rand.Source) {
	problem.difficulty = difficulty
	// Since division problems are significantly more difficult than addition or subtraction problems, we decrease the difficulty for these problems.
	problem.op1 = generateRandomOperand(difficulty-difficulty/2, numSource)
	problem.op2 = generateRandomOperand(difficulty-difficulty/2, numSource)
	problem.answer = problem.op1 / problem.op2
}

func (problem *DivisionProblem) GetDifficulty() int32 {
	return problem.difficulty
}

func (problem *DivisionProblem) String() string {
	return fmt.Sprintf("Division Problem: %d / %d = ?. Answer without the remainder, in other words, round to the nearest integer toward 0 (including 0).\n", problem.op1, problem.op2)
}

func (problem *DivisionProblem) CheckAnswer(answer int32) bool {
	return answer == problem.answer
}
