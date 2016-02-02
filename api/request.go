package api

import (
	"net/http"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
)

// TODO: checks for API error, like
//   {
//     "error": {
//       "type": "unauthorized",
//       "translation_key": "api_error_unauthorized",
//       "message": "You are not authorized."
//     }
//   }
func ErrorCheck(responseData interface{}) (errorMessage string, err error) {
	return
}


// Makes a GET, POST or PATCH request to Wunderlist API.
//
// responseData should be pointer to some struct, which response will be decoded to.
func DoRequest(method string, url string, requestData map[string]string, responseData interface{}) (err error) {
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
	req.Header.Set("X-Access-Token", GetAccessToken())
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

	errorMessage, err := ErrorCheck(&responseData)
	if err != nil {
		return fmt.Errorf("API error: {}", errorMessage)
	}

	return
}
