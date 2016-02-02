package main

import (
	"os"
	"github.com/slasyz/wundercli/api"
	"fmt"
)

var (
	listNamePrompt = []string{"List name"}
	taskPrompt = []string{"List name", "Task text"}
)

// Gets token by calling api.Auth(), puts it to config variable and
// saves it to file.
func handleToken(configPath string) {
	api.DoAuth()
	config.AccessToken = api.GetAccessToken()
	saveConfigFile(configPath)
}

// Gets parameters from command-line arguments or asks them from standard input.
func getParams(prompt []string) (result []string) {
	count := len(prompt)
	result = make([]string, count)

	// From command-line arguments
	for i := 0; i < len(os.Args) - 2; i++ {
		result[i] = os.Args[i+2];
	}
	// From standard input
	for i := len(os.Args) - 2; i < count; i++ {
		fmt.Print(prompt[i] + ": ")
		fmt.Scanln(&result[i])
	}

	return result;
}

func main() {
	configPath := getConfigPath()

	if _, err := os.Stat(configPath); err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Config reading error")
			os.Exit(1)
		}

		handleToken(configPath)
	} else {
		err = parseConfigFile(configPath)

		if err != nil {
			fmt.Println("Config parsing error.")
			os.Exit(1)
		}
	}

	// Just a test
	//var response []struct {
	//	ID int
	//	Title string
	//}
	//err := api.DoRequest("GET", "https://a.wunderlist.com/api/v1/lists", nil, &response)
	//if err != nil {
	//	fmt.Println(response)
	//}

	// Parse command-line arguments.
	if len(os.Args) == 1 {
		cmdHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		cmdList()
	case "show":
		params := getParams(listNamePrompt)
		cmdShow(params[0])
	case "create":
		params := getParams(listNamePrompt)
		cmdCreate(params[0])
	case "remove":
		params := getParams(listNamePrompt)
		cmdRemove(params[0])
	case "append":
		params := getParams(taskPrompt)
		cmdAppend(params[0], params[1])
	case "check":
		params := getParams(listNamePrompt)
		cmdCheck(params[0])
	case "edit":
		params := getParams(listNamePrompt)
		cmdEdit(params[0])
	default:
		cmdHelp()
	}

}
