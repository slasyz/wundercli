package main

import (
	"fmt"
	"github.com/slasyz/wundercli/api"
	"github.com/slasyz/wundercli/config"
	"os"
)

// Gets parameters from command-line arguments or set them empty if not present.
func getParams(count int) (result []string) {
	result = make([]string, count)

	// From command-line arguments
	for i := 0; i < len(os.Args)-3; i++ {
		result[i] = os.Args[i+3]
	}
	// Empty parameters
	for i := len(os.Args) - 3; i < count; i++ {
		result[i] = ""
	}
	fmt.Println()

	return result
}

func main() {
	exists, err := config.OpenConfig()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	if !exists {
		api.DoAuth()
		config.SaveConfig()
	}

	// Parse command-line arguments.
	if len(os.Args) <= 2 {
		cmdHelp()
	}

	switch os.Args[1] {
	case "list":
		switch os.Args[2] {
		case "all":
			err = cmdListAll()
		case "show":
			params := getParams(1)
			err = cmdListShow(params[0])
		case "create":
			params := getParams(1)
			err = cmdListCreate(params[0])
		case "remove":
			params := getParams(1)
			err = cmdListRemove(params[0])
		default:
			cmdHelp()
		}
	case "task":
		switch os.Args[2] {
		case "create":
			params := getParams(2)
			err = cmdTaskCreate(params[0], params[1])
		case "check":
			params := getParams(1)
			err = cmdTaskCheck(params[0])
		case "edit":
			params := getParams(1)
			err = cmdTaskEdit(params[0])
		default:
			cmdHelp()
		}
	default:
		cmdHelp()
	}

	if err != nil {
		fmt.Printf("Error: %s\n\n", err.Error())
		os.Exit(1)
	}

}
