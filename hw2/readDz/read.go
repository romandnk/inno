package readdz

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func Read() (Data, error) {
	f, err := os.Open("../readDz/dz3.json")
	if err != nil {
		return Data{}, fmt.Errorf("error opening file: %s", err.Error())
	}
	defer f.Close()

	d := json.NewDecoder(f)
	var data Data
	if err := d.Decode(&data); err != nil {
		return Data{}, fmt.Errorf("error decoding: %s", err.Error())
	}

	return data, nil
}
