package api

import (
	"testing"
	"encoding/json"
	"github.com/slasyz/wundercli/config"
)

func TestErrorCheckError(t *testing.T) {
	var responseData interface{}
	json.Unmarshal([]byte(`
		{
			"error": {
				"type": "unauthorized",
				"translation_key": "api_error_unauthorized",
				"message": "You are not authorized."
			}
		}
	`), &responseData)

	expectedStr := "You are not authorized."
	var result string

	err := ErrorCheck(&responseData)
	if err != nil {
		result = err.Error()
	} else {
		t.Fatalf("Expected \"%s\", got no error", expectedStr)
	}

	if expectedStr != result {
		t.Fatalf("Expected \"%s\", got %s", expectedStr, result)
	}
}

func TestErrorCheckOkObj(t *testing.T) {
	var responseData interface{}
	json.Unmarshal([]byte(`
		{
			"id": 409233670,
			"assignee_id": 123,
			"created_at": "2013-08-30T08:36:13.273Z",
			"created_by_id": 6234958,
			"due_date": "2013-08-30",
			"list_id": -12345,
			"revision": 1,
			"starred": false,
			"title": "Hello"
		}
	`), &responseData)

	err := ErrorCheck(&responseData)
	if err != nil {
		t.Fatalf("Case: JSON object. Expected no error, got %s", err.Error())
	}
}

func TestErrorCheckOkList(t *testing.T) {
	var responseData interface{}
	json.Unmarshal([]byte(`
		[
			{
				"id": 409233670,
				"assignee_id": 12345,
				"assigner_id": 5432,
				"created_at": "2013-08-30T08:36:13.273Z",
				"created_by_id": 6234958,
				"due_date": "2013-08-30",
				"list_id": 123,
				"revision": 1,
				"starred": true,
				"title": "Hello"
			},
			{
				"id": 540574744,
				"assignee_id": 432,
				"assigner_id": 23424,
				"created_at": "2014-08-30T08:36:13.273Z",
				"created_by_id": 6234950,
				"due_date": "2014-08-30",
				"list_id": 123,
				"revision": 1,
				"starred": true,
				"title": "Hello again"
			},
]
	`), &responseData)

	err := ErrorCheck(&responseData)
	if err != nil {
		t.Fatalf("Case: JSON array. Expected no error, got %s", err.Error())
	}
}

func TestDoRequest(t *testing.T) {
	exists, err := config.OpenConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !exists {
		DoAuth()
		config.SaveConfig()
	}

	if (clientID == "") || (clientSecret == "") {
		t.Skip("Cannot run test because of empty Client ID/Secret.")
	}

	var response []List
	err = DoRequest("GET", "https://a.wunderlist.com/api/v1/lists", nil, &response)
	if err != nil {
		t.Fatalf("DoRequest error: %s", err.Error())
	}

	t.Logf("response length is %d\n", len(response))
}