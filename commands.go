package main

import "fmt"

func cmdHelp() {
	fmt.Println(`Usage:

  - Show list of lists:
    $ wundercli list

  - Show tasks from list:
    $ wundercli show [LISTNAME]

  - Create new list:
    $ wundercli create [LISTNAME]

  - Remove the list:
    $ wundercli remove [LISTNAME]

  - Add task to a list:
    $ wundercli append [LISTNAME [TASKTEXT]]

  - Mark task checked:
    $ wundercli check [LISTNAME]

  - Edit task:
    $ wundercli edit [LISTNAME]
`)
}

func cmdList() {

}

func cmdShow(listName string) {

}

func cmdCreate(listName string) {

}

func cmdRemove(listName string) {

}

func cmdAppend(listName string, taskText string) {

}

func cmdCheck(listName string) {

}

func cmdEdit(listName string) {

}
