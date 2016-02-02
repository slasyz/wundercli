package api

import (
	"fmt"
)

// Struct for JSON list object.
type List struct {
	ID int
	Title string
	Revision int
}

// Struct for JSON task object.
type Task struct {
	ID int
	List_ID int
	Title string
	Revision int
}

func GetLists() (lists []List, err error) {
	err = DoRequest("GET", "https://a.wunderlist.com/api/v1/lists", nil, &lists)

	return
}

// Gets completed tasks from specified list.
func GetListTasks(list List) (tasks []Task, err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/tasks?completed=false&list_id=%d", list.ID)
	err = DoRequest("GET", url, nil, &tasks)

	return
}

// Creates a list with specified title.
func DoListCreate(listTitle string) (err error) {
	err = DoRequest("POST", "https://a.wunderlist.com/api/v1/lists", map[string]interface{}{
		"title": listTitle,
	}, nil)

	return
}

// Removes a list.
func DoListRemove(list List) (err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/lists/%d?revision=%d", list.ID, list.Revision)
	err = DoRequest("DELETE", url, nil, nil)

	return
}

// Creates task with specified title to specified list.
func DoTaskAppend(list List, taskText string) (err error) {
	err = DoRequest("POST", "https://a.wunderlist.com/api/v1/tasks", map[string]interface{}{
		"list_id": list.ID,
		"title": taskText,
	}, nil)

	return
}

// Marks a task as checked.
func DoTaskCheck(task Task) (err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/tasks/%d", task.ID)
	err = DoRequest("PATCH", url, map[string]interface{}{
		"revision": task.Revision,
		"completed": true,
	}, nil)

	return
}

// Edits a task title.
func DoTaskEdit(task Task, taskText string) (err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/tasks/%d", task.ID)
	err = DoRequest("PATCH", url, map[string]interface{}{
		"revision": task.Revision,
		"title": taskText,
	}, nil)

	return
}
