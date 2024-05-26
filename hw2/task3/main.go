package main

import (
	"encoding/json"
	"fmt"
	"github.com/rodaine/table"
	"os"
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

type Data struct {
	Results  []Results  `json:"results"`
	Objects  []Objects  `json:"objects"`
	Students []Students `json:"students"`
}

type Results struct {
	ObjectId  int `json:"object_id"`
	StudentId int `json:"student_id"`
	Result    int `json:"result"`
}

type Objects struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Students struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func main() {
	f, err := os.Open("dz3.json")
	if err != nil {
		fmt.Printf("error opening file: %s\n", err.Error())
		return
	}
	defer f.Close()

	d := json.NewDecoder(f)
	var data Data
	if err := d.Decode(&data); err != nil {
		fmt.Printf("error decoding: %s\n", err.Error())
		return
	}

	studentsByID := make(map[int]Students, len(data.Students))
	for _, student := range data.Students {
		studentsByID[student.Id] = student
	}

	objectsByID := make(map[int]Objects, len(data.Objects))
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
