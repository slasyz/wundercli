package api

import (
	"net/http"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
	"github.com/slasyz/wundercli/config"
)

// Checks for API error.
func ErrorCheck(responseData *interface{}) (err error) {
	switch (*responseData).(type) {
	case map[string]interface{}:
		errorVal, thereIsAnError := (*responseData).(map[string]interface{})["error"]
		if thereIsAnError {
			var errorMessage string
			switch errorVal.(type) {
			case map[string]interface{}:
				errorMessage = (errorVal.(map[string]interface{})["message"]).(string)
			case string:
				errorMessage = errorVal.(string)
			default:
				errorMessage = "unknown error"
			}


			return errors.New(errorMessage)
		} else {
			return
		}
	default:
		return
	}
}


// Makes a GET, POST or PATCH request to Wunderlist API.
//
// responseData should be pointer to some struct, which response will be decoded to.
func DoRequest(method string, url string, requestData map[string]interface{}, responseData interface{}) (err error) {
	var requestDataByte []byte
	if requestData == nil {
		requestDataByte = []byte{}
	} else {
		requestDataByte, _ = json.Marshal(requestData)
	}

	//fmt.Println("=======req=======")
	//fmt.Println(string(requestDataByte))
	//fmt.Println("======/req=======")

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestDataByte))
	if err != nil {
		return errors.New("request creation error")
	}
	req.Header.Set("X-Access-Token", config.Config.AccessToken)
	req.Header.Set("X-Client-ID", clientID)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("response getting error")
	}

	//fmt.Println("=======body========")
	//fmt.Println(string(body))
	//fmt.Println("======/body========")

	json.Unmarshal(body, &responseData)
	//fmt.Println("=====decoded====")
	//fmt.Printf("%+v\n", responseData)
	//fmt.Println("=====/decoded====")

	err = ErrorCheck(&responseData)
	if err != nil {
		return fmt.Errorf("API failure (%s)", err.Error())
	}

	return
}
