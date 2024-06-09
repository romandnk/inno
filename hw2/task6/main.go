package main

import (
	"fmt"
	"github.com/rodaine/table"
	"golang.org/x/exp/constraints"
	readdz "inno/hw2/readDz"
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

	studentsById := make(map[int]readdz.Students, len(data.Students))
	for _, student := range data.Students {
		studentsById[student.Id] = student
	}

	// key is a subject id, value is a map which contains grade as a key and assessments as values
	objGrades := make(map[int]map[int][]int, len(data.Objects))
	for _, res := range data.Results {
		if _, ok := objGrades[res.ObjectId]; !ok {
			objGrades[res.ObjectId] = make(map[int][]int)
		}

		student := studentsById[res.StudentId]
		objGrades[res.ObjectId][student.Grade] = append(objGrades[res.ObjectId][student.Grade], res.Result)
	}

	for _, obj := range data.Objects {
		tbl := table.New(obj.Name, "Mean")
		tbl.WithHeaderSeparatorRow('-')

		resultsByGrade, ok := objGrades[obj.Id]
		if !ok {
			continue
		}

		var totalSum float64
		var totalCount int
		for grade, marks := range resultsByGrade {
			markSum := numberArrayAction(marks, 0, func(a, b int) int {
				return a + b
			})
			totalSum += float64(markSum)
			totalCount += len(marks)

			meanByGrade := float64(markSum) / float64(len(marks))
			tbl.AddRow(grade, meanByGrade)
		}

		totalMeanByObj := totalSum / float64(totalCount)
		tbl.AddRow("mean", totalMeanByObj)

		tbl.Print()
	}
}

type numbers interface {
	constraints.Integer | constraints.Float
}

// reduce
func numberArrayAction[T numbers](arr []T, init T, f func(T, T) T) T {
	value := init
	for _, v := range arr {
		value = f(value, v)
	}
	return value
}
