package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
)

/* 7. Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
...
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

	notExcellentResults := filter(data.Results, func(results readdz.Results) bool {
		return results.Result != 5
	})

	notExcellentStudent := make(map[int]struct{}, len(notExcellentResults))
	for _, result := range notExcellentResults {
		notExcellentStudent[result.StudentId] = struct{}{}
	}

	tbl := table.New("Student name", "Grade", "Object", "Result")
	for _, result := range data.Results {
		student := studentsByID[result.StudentId]
		if _, ok := notExcellentStudent[student.Id]; ok {
			continue
		}

		obj := objectsByID[result.ObjectId]

		tbl.AddRow(student.Name, student.Grade, obj.Name, result.Result)
	}

	tbl.Print()
}

func filter[T comparable](arr []T, f func(T) bool) []T {
	res := make([]T, 0, len(arr))
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// without filter
//func main() {
//	data, err := readdz.Read()
//	if err != nil {
//		fmt.Println("error: ", err.Error())
//		return
//	}
//
//	studentsByID := make(map[int]readdz.Students, len(data.Students))
//	for _, student := range data.Students {
//		studentsByID[student.Id] = student
//	}
//
//	objectsByID := make(map[int]readdz.Objects, len(data.Objects))
//	for _, object := range data.Objects {
//		objectsByID[object.Id] = object
//	}
//
//	excellentStudent := make(map[int]struct{}, len(data.Students))
//	notExcellentStudent := make(map[int]struct{}, len(data.Students))
//	for _, result := range data.Results {
//		if result.Result == 5 {
//			if _, ok := notExcellentStudent[result.StudentId]; ok {
//				continue
//			}
//			excellentStudent[result.StudentId] = struct{}{}
//		} else {
//			notExcellentStudent[result.StudentId] = struct{}{}
//			if _, ok := excellentStudent[result.StudentId]; ok {
//				delete(excellentStudent, result.StudentId)
//			}
//		}
//	}
//
//	tbl := table.New("Student name", "Grade", "Object", "Result")
//	for _, result := range data.Results {
//		student := studentsByID[result.StudentId]
//		if _, ok := excellentStudent[student.Id]; !ok {
//			continue
//		}
//
//		obj := objectsByID[result.ObjectId]
//
//		tbl.AddRow(student.Name, student.Grade, obj.Name, result.Result)
//	}
//
//	tbl.Print()
//}
