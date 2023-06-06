package utils

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

func TestDecodeBody(t *testing.T) {
	jsonBody := []byte(`{
		"entries": {
			"1": "2",
			"3": "4"
		},
		"version": "v1"
	}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/", bodyReader)
	_, err := DecodeBody(req.Body, context.Background())
	if err != nil {
		t.Errorf("Test not successful. Expected: Configuration, but got: %s", err.Error())
	}
}

func TestDecodeBodyGroup(t *testing.T) {
	jsonBody := []byte(`{
		"group": [
		  {
			"label": {
			  "key123": "value1",
			  "key213": "value2"
			},
			"entries": {
			  "key1": "value1",
			  "key2": "value2"
			}
		  
		
		  }
		],
		"version": "1.0"
	  }`)
	bodyReader := bytes.NewReader(jsonBody)
	req, _ := http.NewRequest("POST", "/", bodyReader)
	_, err := DecodeBodyForGroup(req.Body, context.Background())
	if err != nil {
		t.Errorf("Test not successful. Expected: ConfigurationGroup, but got: %s", err.Error())
	}
}

func TestGetKeyIndexInfo(t *testing.T) {
	key := "config/groupID/groupVersion/someValue"

	valueLookingFor := "groupID"
	expected := "groupID"
	result := GetKeyIndexInfo(valueLookingFor, key)

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	valueLookingFor = "groupVersion"
	expected = "groupVersion"
	result = GetKeyIndexInfo(valueLookingFor, key)

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	valueLookingFor = "someValue"
	expected = ""
	result = GetKeyIndexInfo(valueLookingFor, key)

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestGetLabelAsStringWithSeparator(t *testing.T) {
	label := map[string]string{
		"key1": "value1",
	}

	expected := "key1:value1"
	result := GetLabelAsStringWithSeparator(label)

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
