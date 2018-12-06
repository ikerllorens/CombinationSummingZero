package main

import "fmt"

type Combination struct {
	number1 int
	number2 int
	number3 int
}

func main() {
	array := []int{-3, -2, 0 ,1, 2}
	result := analyzeArrayAndGetCombinations(array)
	fmt.Println(result)
}


func analyzeArrayAndGetCombinations(numbers []int) []Combination {
	combinationsSumZero := make([]Combination, 0)

	resultComb := createCombination(numbers)

	for i:=0; i<len(resultComb);i++ {
		sum := resultComb[i].number1 + resultComb[i].number2 + resultComb[i].number3
		if(sum == 0) {
			combinationsSumZero = append(combinationsSumZero,resultComb[i])
		}
	}

	return combinationsSumZero
}


func createCombination(numbers []int) []Combination {
	combinations := make([]Combination, 0)


	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			for k := j; k < len(numbers); k++ {
				if  i != j && i != k && j != k  {
					var comb Combination

					comb.number1 = numbers[i]
					comb.number2 = numbers[j]
					comb.number3 = numbers[k]

					combinations = append(combinations, comb)
				}
			}
		}
	}

	return combinations
}



/*
Your previous Bash content is preserved below:

Given an Array with numbers, return all combinations of 3 of those numbers that sum to zero

Example:
input: [-3, -2, 0 ,1, 2]
output: [-2, 0, 2], [-3, 1, 2]
 */