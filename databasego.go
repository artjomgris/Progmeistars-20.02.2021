package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Action struct {
	Action  string `json:"action"`
	ObjName string `json:"object"`
}

type Teacher struct {
	ID        string   `json:"id"`
	Salary    float64  `json:"salary"`
	Subject   string   `json:"subject"`
	Classroom []string `json:"classroom"`
	Person    struct {
		Name         string `json:"name"`
		Surname      string `json:"surname"`
		PersonalCode string `json:"personalCode"`
	} `json:"person"`
}

func (t Teacher) GetCreateAction() DefinedAction {
	return &CreateTeacher{}
}
func (t Teacher) GetUpdateAction() DefinedAction {
	return &UpdateTeacher{}
}
func (t Teacher) GetReadAction() DefinedAction {
	return &ReadTeacher{}
}
func (t Teacher) GetDeleteAction() DefinedAction {
	return &DeleteTeacher{}
}

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	GetCreateAction() DefinedAction
	GetUpdateAction() DefinedAction
	GetReadAction() DefinedAction
	GetDeleteAction() DefinedAction
}

type CreateTeacher struct {
	T Teacher `json:"data"`
}

func (action *CreateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println("60", err)
		return
	}
}
func (action CreateTeacher) Process() {
	//add action.T to slice of data
	fmt.Println(action.T)
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}

func (action *UpdateTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println("76", err)
		return
	}
}
func (action UpdateTeacher) Process() {
	fmt.Println("Teacher Updated")
}

type ReadTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *ReadTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println("93", err)
		return
	}
}
func (action ReadTeacher) Process() {
	fmt.Println("Read techer")
}

type DeleteTeacher struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (action *DeleteTeacher) GetFromJSON(rawData []byte) {
	err := json.Unmarshal(rawData, action)
	if err != nil {
		fmt.Println("110", err)
		return
	}
}
func (action DeleteTeacher) Process() {
	fmt.Println("Teacher deleted")
}

func main() {
	f, err := os.Open("data.json")
	defer f.Close()
	if err != nil {
		fmt.Println("122", err)
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		var act Action
		err = json.Unmarshal(line, &act)
		if err != nil {
			fmt.Println("134", err)
			return
		}

		var obj GeneralObject
		switch act.ObjName {
		case "Teacher":
			obj = &Teacher{}
		}
		var toDo DefinedAction
		switch act.Action {
		case "create":
			toDo = obj.GetCreateAction()
		case "update":
			toDo = obj.GetUpdateAction()
		case "read":
			toDo = obj.GetReadAction()
		case "delete":
			toDo = obj.GetDeleteAction()
		}
		toDo.GetFromJSON(line)
		toDo.Process()
	}
}
