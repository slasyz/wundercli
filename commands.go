package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/slasyz/wundercli/api"
	"os"
	"strings"
)

// Gets list object by its short name.
// Works like shell tab completion.
func getListByShortName(listName string) (list api.List, err error) {
	lists, err := api.GetLists()
	if err != nil {
		return
	}

	var sel []api.List
	listNameLower := strings.ToLower(listName)
	for _, el := range lists {
		elLower := strings.ToLower(el.Title)
		if strings.HasPrefix(elLower, listNameLower) || listName == "" {
			if elLower == listNameLower {
				sel = []api.List{el}
				break
			}

			sel = append(sel, el)
		}
	}

	if len(sel) == 0 {
		err = errors.New("List not found")
		return
	} else if len(sel) > 1 {
		if listName == "" {
			fmt.Println("Currently available lists:\n")
		} else {
			fmt.Printf("There is several lists starting with \"%s\":\n\n", listName)
		}
		for i, el := range sel {
			fmt.Printf("  [%d] %s\n", i+1, el.Title)
		}
		fmt.Println()

		fmt.Print("What is the number of the list? ")
		var listNo int
		fmt.Scanln(&listNo)
		fmt.Println()

		if listNo > len(sel) || listNo < 1 {
			err = errors.New("Incorrect input.")
			return
		}

		return sel[listNo-1], nil
	} else {
		return sel[0], nil
	}
}

// Outputs list tasks and asks for one of it.
func askForTask(list api.List) (task api.Task, err error) {
	tasks, err := api.GetListTasks(list)
	if err != nil {
		return
	}

	for i, el := range tasks {
		fmt.Printf("  [%d] %s\n", i+1, el.Title)
	}
	fmt.Println()

	fmt.Print("What is the number of the task? ")
	var taskNo int
	fmt.Scanln(&taskNo)
	fmt.Println()

	if taskNo > len(tasks) || taskNo < 1 {
		err = errors.New("Incorrect input.")
		return
	}

	return tasks[taskNo-1], nil
}

func cmdHelp() {
	fmt.Println(`
Usage:

  - Show list of lists:
    $ wundercli list all

  - Show tasks from list:
    $ wundercli list show [LISTNAME]

  - Create new list:
    $ wundercli list create [LISTTITLE]

  - Remove the list:
    $ wundercli list remove [LISTNAME]

  - Add task to a list:
    $ wundercli task create [LISTNAME [TASKTEXT]]

  - Mark task checked:
    $ wundercli task check [LISTNAME]

  - Edit task:
    $ wundercli task edit [LISTNAME]
`)
	os.Exit(0)
}

func cmdListAll() (err error) {
	lists, err := api.GetLists()
	if err != nil {
		return
	}

	fmt.Println("Available lists are:\n")
	for _, el := range lists {
		fmt.Printf("    - %s\n", el.Title)
	}
	fmt.Println()

	return
}

func cmdListShow(listName string) (err error) {
	list, err := getListByShortName(listName)
	if err != nil {
		return
	}

	tasks, err := api.GetListTasks(list)
	if err != nil {
		return
	}

	fmt.Printf("Tasks from \"%s\" list:\n\n", list.Title)
	for _, el := range tasks {
		fmt.Printf("    - %s\n", el.Title)
	}
	fmt.Println()

	return
}

func cmdListCreate(listTitle string) (err error) {
	if listTitle == "" {
		fmt.Print("Enter new list title: ")
		in := bufio.NewReader(os.Stdin)
		listTitle, err = in.ReadString(byte('\n'))
		if err != nil {
			return errors.New("reading from console")
		}
		fmt.Println()

		if listTitle == "" {
			return errors.New("list title cannot be empty")
		}
	}

	err = api.DoListCreate(listTitle)
	if err != nil {
		return
	}

	fmt.Println("List was created successfully.\n")
	return
}

func cmdListRemove(listName string) (err error) {
	list, err := getListByShortName(listName)
	if err != nil {
		return
	}

	err = api.DoListRemove(list)
	if err != nil {
		return
	}

	fmt.Println("List was removed successfully.\n")
	return
}

func cmdTaskCreate(listName string, taskText string) (err error) {
	list, err := getListByShortName(listName)
	if err != nil {
		return
	}

	if taskText == "" {
		fmt.Print("Enter task text: ")
		in := bufio.NewReader(os.Stdin)
		taskText, err = in.ReadString(byte('\n'))
		if err != nil {
			return errors.New("reading from console")
		}

		if taskText == "" {
			return errors.New("task text cannot be empty")
		}
	}
	fmt.Println()

	err = api.DoTaskCreate(list, taskText)
	if err != nil {
		return
	}

	fmt.Println("Task was created successfully.\n")
	return
}

func cmdTaskCheck(listName string) (err error) {
	list, err := getListByShortName(listName)
	if err != nil {
		return
	}

	task, err := askForTask(list)
	if err != nil {
		return
	}

	err = api.DoTaskCheck(task)
	if err != nil {
		return
	}

	fmt.Println("Task was marked checked successfully.\n")
	return
}

func cmdTaskEdit(listName string) (err error) {
	list, err := getListByShortName(listName)
	if err != nil {
		return
	}

	task, err := askForTask(list)
	if err != nil {
		return
	}

	fmt.Print("What is the new text of the task? ")
	var taskText string
	fmt.Scanln(&taskText)
	fmt.Println()

	if taskText == "" {
		return errors.New("task text cannot be empty")
	}

	err = api.DoTaskEdit(task, taskText)
	if err != nil {
		return
	}

	fmt.Println("Task was edited successfully.\n")
	return
}
