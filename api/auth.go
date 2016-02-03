package api

import (
	"errors"
	"fmt"
	"github.com/slasyz/wundercli/config"
	"os/exec"
	"runtime"
)

// Starts an authentication process:
//   - Opens a browser;
//   - asks for the code;
//   - gets access token in exchange for code.
func DoAuth() (err error) {
	redirectUrl := "http://slasyz.ru/wundercli/"
	url := fmt.Sprintf("https://www.wunderlist.com/oauth/authorize?client_id=%s&redirect_uri=%s&state=", clientID, redirectUrl)

	// Open browser.
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows", "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		return errors.New("Error opening browser.")
	}

	// Ask for the code.
	fmt.Print("Enter the code: ")
	var code string
	fmt.Scanln(&code)

	// Exchange code for access token.
	var data struct {
		Access_Token string
	}
	err = DoRequest("POST", "https://www.wunderlist.com/oauth/access_token", map[string]interface{}{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}, &data)
	if err != nil {
		return errors.New("Authentication error.")
	}

	config.Config.AccessToken = data.Access_Token

	return nil
}
