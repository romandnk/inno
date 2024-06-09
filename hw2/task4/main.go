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

//func main() {
//	data, err := readdz.Read()
//	if err != nil {
//		fmt.Println("error: ", err.Error())
//		return
//	}
//
//	numstudentsByGrade := make(map[int]int, len(data.Students))
//	studentsById := make(map[int]readdz.Students, len(data.Students))
//	for _, student := range data.Students {
//		studentsById[student.Id] = student
//		numstudentsByGrade[student.Grade]++
//	}
//
//	allResultByObject := make(map[int][]readdz.Results, len(data.Objects))
//	for _, results := range data.Results {
//		allResultByObject[results.ObjectId] = append(allResultByObject[results.ObjectId], results)
//	}
//
//
//
//	for _, obj := range data.Objects {
//		tbl := table.New(obj.Name, "Mean")
//		tbl.WithHeaderSeparatorRow('-')
//
//		sumByGrade := make(map[int]float64, len(data.Objects)*len(data.Students))
//
//		resultsByObject := allResultByObject[obj.Id]
//
//		for _, result := range resultsByObject {
//			student := studentsById[result.StudentId]
//			sumByGrade[student.Grade] += float64(result.Result)
//		}
//
//		var overallMean float64
//		for grade, sum := range sumByGrade {
//			mean := sum / float64(numstudentsByGrade[grade])
//			overallMean += mean
//			tbl.AddRow(grade, mean, sum)
//		}
//		tbl.AddRow("mean", math.Round(overallMean/3*100)/100)
//
//		tbl.Print()
//		fmt.Println("-----------")
//	}
//}

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
			var markSum float64
			for _, mark := range marks {
				markSum += float64(mark)
			}
			totalSum += markSum
			totalCount += len(marks)

			meanByGrade := markSum / float64(len(marks))
			tbl.AddRow(grade, meanByGrade)
		}

		totalMeanByObj := totalSum / float64(totalCount)
		tbl.AddRow("mean", totalMeanByObj)

		tbl.Print()
	}
}
