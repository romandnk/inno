package main

import (
	"errors"
	"fmt"
)

func EvalSequence(matrix [][]int, userAnswer []int) (int, error) {
	// validate matrix
	err := validateMatrix(matrix)
	if err != nil {
		return 0, err
	}
	// validate user answers
	err = ValidateUserAnswers(matrix, userAnswer)
	if err != nil {
		return 0, err
	}

	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	// max grade cannot be less than zero because we have checked it in validateMatrix function
	percent := userGrade * 100 / maxGrade

	return percent, nil
}

func validateMatrix(matrix [][]int) error {
	n := len(matrix)
	if n == 0 {
		return errors.New("matrix is empty")
	}

	isPathsEmpty := false

	for i := range n {
		if len(matrix[i]) != n {
			return errors.New("matrix is not square")
		}
		if matrix[i][i] != 0 {
			return errors.New("matrix has loop")
		}
		for j := 0; j < n; j++ {
			value := matrix[i][j]
			if value < 0 {
				return errors.New("matrix has a negative value")
			}
			if value > 0 {
				isPathsEmpty = true
			}
		}
	}

	if !isPathsEmpty {
		return errors.New("paths are empty")
	}

	return nil
}

func ValidateUserAnswers(matrix [][]int, userAnswer []int) error {
	n := len(userAnswer)
	m := len(matrix)
	if n == 0 {
		return nil
	}

	err := validateAnswerRange(userAnswer[0], m)
	if err != nil {
		return err
	}

	exist := make(map[int]struct{})
	exist[userAnswer[0]] = struct{}{}

	for i := 1; i < n; i++ {
		if _, ok := exist[userAnswer[i]]; ok {
			return fmt.Errorf("answer is duplicated: %d", userAnswer[i])
		}
		err := validateAnswerRange(userAnswer[i], m)
		if err != nil {
			return err
		}
		exist[userAnswer[i]] = struct{}{}
	}

	return nil
}

func validateAnswerRange(answer, maxAnswer int) error {
	if answer >= maxAnswer || answer < 0 {
		return fmt.Errorf("invalid answer: %d", answer)
	}
	return nil
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	if len(userAnswer) == 0 {
		return 0
	}

	userGrade := 0
	for i := 0; i < len(userAnswer)-1; i++ {
		fromVert := userAnswer[i]
		toVert := userAnswer[i+1]

		userGrade += matrix[fromVert][toVert]
	}

	return userGrade
}

func calMaxGrade(matrix [][]int) int {
	visited := make([]bool, len(matrix))
	maxWeight := 0

	for i := range matrix {
		dfs(matrix, i, visited, 0, &maxWeight)
	}

	return maxWeight
}

func dfs(matrix [][]int, vertex int, visited []bool, currentWeight int, maxWeight *int) {
	visited[vertex] = true

	for i := range matrix[vertex] {
		if matrix[vertex][i] != 0 && !visited[i] {
			dfs(matrix, i, visited, currentWeight+matrix[vertex][i], maxWeight)
		}
	}

	if currentWeight > *maxWeight {
		*maxWeight = currentWeight
	}

	visited[vertex] = false
}
