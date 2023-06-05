package utils

import (
	"bytes"
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
	_, err := DecodeBody(req.Body)
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
	_, err := DecodeBodyForGroup(req.Body)
	if err != nil {
		t.Errorf("Test not successful. Expected: ConfigurationGroup, but got: %s", err.Error())
	}
}
