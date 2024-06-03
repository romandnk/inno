package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
	"math"
)

/* 6. Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.
Желательно реализуйте эти функции самостоятельно.
*/

func main() {
	data, err := readdz.Read()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	// решение с учетом того, что каждый предмет есть у каждого ученика
	numstudentsByGrade := make(map[int]int, len(data.Students))
	studentsById := make(map[int]readdz.Students, len(data.Students))
	for _, student := range data.Students {
		studentsById[student.Id] = student
		numstudentsByGrade[student.Grade]++
	}

	allResultByObject := make(map[int][]readdz.Results, 3)
	for _, results := range data.Results {
		allResultByObject[results.ObjectId] = append(allResultByObject[results.ObjectId], results)
	}

	for _, obj := range data.Objects {
		tbl := table.New(obj.Name, "Mean")
		tbl.WithHeaderSeparatorRow('-')

		sumByGrade := make(map[int]float64, len(data.Objects)*len(data.Students))

		resultsByObject := allResultByObject[obj.Id]

		for _, result := range resultsByObject {
			student := studentsById[result.StudentId]
			sumByGrade[student.Grade] += float64(result.Result)
		}

		res := calculateMean(sumByGrade, func(grade int, sum float64) float64 {
			return sum / float64(numstudentsByGrade[grade])
		})

		var overallMean float64
		for _, r := range res {
			overallMean += r.Mean
			tbl.AddRow(r.Grade, r.Mean, r.Sum)
		}

		tbl.AddRow("mean", math.Round(overallMean/3*100)/100)

		tbl.Print()
		fmt.Println("-----------")
	}
}

type Mean struct {
	Grade int
	Mean  float64
	Sum   float64
}

func calculateMean(sumByGrade map[int]float64, f func(grade int, sum float64) float64) []Mean {
	var result []Mean
	for grade, sum := range sumByGrade {
		mean := f(grade, sum)
		result = append(result, Mean{
			Grade: grade,
			Sum:   sum,
			Mean:  mean,
		})
	}
	return result
}
