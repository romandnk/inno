package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
)

/* 3. У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида
_____________________________________
Student name  | Grade | Object    |   Result
____________________________________
Ann			  |     9 | Math	  |  4
Ann 		  |     9 | Biology   |  4
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

	tbl := table.New("Student name", "Grade", "Object", "Result")
	tbl.WithHeaderSeparatorRow('-')

	for _, result := range data.Results {
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
