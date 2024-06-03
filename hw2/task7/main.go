package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
)

/* 7. Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.
*/

func main() {
	data, err := readdz.Read()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	studentsByID := make(map[int]readdz.Students, len(data.Students))
	for _, student := range data.Students {
		studentsByID[student.Id] = student
	}

	objectsByID := make(map[int]readdz.Objects, len(data.Objects))
	for _, object := range data.Objects {
		objectsByID[object.Id] = object
	}

	tbl := table.New("Student name", "Grade", "Object", "Result")
	tbl.WithHeaderSeparatorRow('-')

	results := excellentStudent(data.Results, func(r readdz.Results) bool {
		if r.Result == 5 {
			return true
		}
		return false
	})

	for _, result := range results {
		row := []any{
			studentsByID[result.StudentId].Name,
			studentsByID[result.StudentId].Grade,
			objectsByID[result.ObjectId].Name,
			result.Result,
		}
		tbl.AddRow(row...)
	}
	tbl.Print()
}

func excellentStudent(results []readdz.Results, f func(r readdz.Results) bool) []readdz.Results {
	var out []readdz.Results
	for _, res := range results {
		if f(res) {
			out = append(out, res)
		}
	}
	return out
}
