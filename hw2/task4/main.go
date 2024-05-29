package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
)

/* 4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам в виде:
________________
Math	 | Mean
________________
 9 grade | 4.5
10 grade | 5
11 grade | 3.5
________________
mean     | 4		- среднее значение среди всех учеников
________________
________________
Biology	 | Mean
________________
...
Вводные данные представлены в файле dz3.json
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

	for _, obj := range data.Objects {
		tbl := table.New(obj.Name, "Mean")
		tbl.WithHeaderSeparatorRow('-')

		allResultByObject := make([]readdz.Results, 0, len(data.Objects)*len(data.Students))
		for _, results := range data.Results {
			if results.ObjectId == obj.Id {
				allResultByObject = append(allResultByObject, results)
			}
		}

		sumByGrade := make(map[int]float64, len(data.Objects)*len(data.Students))
		studentsByGrade := make(map[int]int, len(data.Students))
		for _, result := range allResultByObject {
			student := studentsByID[result.StudentId]
			studentsByGrade[student.Grade]++
			sumByGrade[student.Grade] += float64(result.Result)
		}

		for grade, sum := range sumByGrade {
			tbl.AddRow(grade, sum/float64(studentsByGrade[grade]), sum)
		}

		tbl.Print()
		fmt.Println("-----------")
	}
}
