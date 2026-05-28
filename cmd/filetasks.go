package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"
	"strconv"
)

var statusOpt = [3]string{"todo", "in-progress", "done"}
var keys = [5]string{"Id", "Task", "Status", "CreatedAt", "UpdatedAt"}
var idTracker int = 102

type Task struct {
	Id          int       `json:"Id"`
	Description string    `json:"Task"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

func HasField(obj interface{}, fieldName string) (bool, string) {

	val := reflect.ValueOf(obj)
	
    if val.Kind() != reflect.Struct {
        return false, "Object passed is not a struct"
    }
    
    return val.FieldByName(fieldName).IsValid(), ""
}

func findObjectInFile(filepath string, searchId int) (*Task, error) {
	
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	
	for decoder.More() {
		var task Task

		if err := decoder.Decode(&task); err != nil {
			return nil, err
		}
		
		if task.Id == searchId {
			ptrToTask := &task
			return ptrToTask, nil
		}
	}

	return nil, fmt.Errorf("Task with id %d not found", searchId)
}

func addTaskToFile(filepath string, tasks []string) (string, error) {
	fmt.Println("Adding task from inside the function")

	file, err := os.OpenFile(filepath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for _, taskDesc := range tasks {
		taskObj := Task{
			Id:          idTracker,
			Description: taskDesc,
			Status:      statusOpt[0],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		
		fmt.Println("Task: ", taskObj)
		taskEncoded, err := json.MarshalIndent(taskObj, "", "\t")
		if err != nil {
			return "", err
		}
		fmt.Println("Encoded Task: ", string(taskEncoded))
		
		bytes, err := file.Write(append(taskEncoded, []byte("\n")...))
		if err != nil {
			fmt.Println("Error: ", err)
			return "", err
		}
		fmt.Println("Bytes: ", bytes)
		idTracker += 1
	}
	result := fmt.Sprintf("File written succesfully")
	return result, nil
}

func listAllTasksFromFile(filepath string) (string, error){

	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	
	for decoder.More() {
    	var t Task
		if err := decoder.Decode(&t); err == nil {
			fmt.Println(t)
		}
	}

	result := "Succesfully listed all tasks"
	return result, nil
}

func deleteTaskFromFile(filepath string, args []string) (string, error) {
	
	intId, err := strconv.Atoi(args[0])
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	
	for decoder.More() {
		var task Task
		if err := decoder.Decode(&task); err != nil {
			return "", err
		}
		if task.Id != intId {
			tasks = append(tasks, task)
		}
	}

	_, errW := file.Seek(0,0)
	if errW != nil {
		return "", errW
	}

	file.Truncate(0)

	for _, taskObj := range tasks {
		taskEncoded, err := json.MarshalIndent(taskObj, "", "\t")
		if err != nil {
			return "", err
		}
		_, err = file.Write(append(taskEncoded, []byte("\n")...))
		if err != nil {
			return "", err
		}
	}

	return "", nil
}


func updateTaskFromFile(filepath string, args []string) (string, error) {

	intId, err := strconv.Atoi(args[0])
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fieldName := args[1]
	var tasks []Task
	decoder := json.NewDecoder(file)
	
	for decoder.More() {
		var task Task
		if err := decoder.Decode(&task); err != nil {
			return "", err
		}
		if task.Id == intId {
			ok, err := HasField(task, fieldName)
			if !ok {
				return "", err
			} else {
				task.fieldName = args[2]
			}
		}
	}

	_, errW := file.Seek(0,0)
	if errW != nil {
		return "", errW
	}

	file.Truncate(0)

	for _, taskObj := range tasks {
		taskEncoded, err := json.MarshalIndent(taskObj, "", "\t")
		if err == nil {
			return "", err
		}
		_, err = file.Write(append(taskEncoded, []byte("\n")...))
		if err != nil {
			return "", err
		}
	}

	return "", nil

}