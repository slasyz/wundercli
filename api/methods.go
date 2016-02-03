package api

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Struct for JSON list object.
type List struct {
	ID       int
	Title    string
	Revision int
}

// Struct for JSON task object.
type Task struct {
	ID       int
	List_ID  int
	Title    string
	Revision int
}

func getMonthNumber(name string) int {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	for i, el := range months {
		if el == name {
			return i + 1
		}
	}

	return -1
}

// Splits string like "Do something Oct 6" into "Do something" and "Oct 6".
func parseDate(origTaskText string) (taskText string, date string) {
	res := regexp.MustCompile(`^(.*[^\s])\s*((Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s(\d{1,2}))\s*$`).FindStringSubmatch(origTaskText)

	if res == nil {
		return origTaskText, ""
	} else {
		taskText = res[1]

		month := getMonthNumber(res[3])
		day, _ := strconv.Atoi(res[4])
		date = fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), month, day)

		return
	}
}

// Returns list of all user's lists.
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

// Creates task with specified title in specified list.
func DoTaskCreate(list List, taskText string) (err error) {
	taskText, date := parseDate(taskText)
	params := map[string]interface{}{
		"list_id": list.ID,
		"title":   taskText,
	}

	if date != "" {
		params["due_date"] = date
	}

	err = DoRequest("POST", "https://a.wunderlist.com/api/v1/tasks", params, nil)

	return
}

// Marks a task as checked.
func DoTaskCheck(task Task) (err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/tasks/%d", task.ID)
	err = DoRequest("PATCH", url, map[string]interface{}{
		"revision":  task.Revision,
		"completed": true,
	}, nil)

	return
}

// Edits a task title.
func DoTaskEdit(task Task, taskText string) (err error) {
	url := fmt.Sprintf("https://a.wunderlist.com/api/v1/tasks/%d", task.ID)
	err = DoRequest("PATCH", url, map[string]interface{}{
		"revision": task.Revision,
		"title":    taskText,
	}, nil)

	return
}
