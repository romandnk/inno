package main

import (
	"fmt"
	"github.com/rodaine/table"
	readdz "inno/hw2/readDz"
)

/* 5. Перепишите задачу #3 с использованием структуры-дженерик Cache, изученной на семинаре.
Храните в кеше таблицы студентов и предметов.
*/

func main() {
	data, err := readdz.Read()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	studentsByID := NewCache[int, readdz.Students](len(data.Students))
	for _, student := range data.Students {
		studentsByID.Set(student.Id, student)
	}

	objectsByID := NewCache[int, readdz.Objects](len(data.Objects))
	for _, object := range data.Objects {
		objectsByID.Set(object.Id, object)
	}

	tbl := table.New("Student name", "Grade", "Object", "Result")
	tbl.WithHeaderSeparatorRow('-')

	for _, result := range data.Results {
		name := studentsByID.Get(result.StudentId).Name
		grade := studentsByID.Get(result.StudentId).Grade
		obj := objectsByID.Get(result.ObjectId).Name
		row := []any{name, grade, obj, result.Result}
		tbl.AddRow(row...)
	}
	tbl.Print()
}
