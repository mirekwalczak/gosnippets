package main

import (
	"encoding/json"
	"fmt"
)

type CI struct {
	Name        string            `json:"ciname"`
	UniqueName  string            `json:"uniqueciname"`
	Environment string            `json:"environment"`
	ManagedBy   string            `json:"managedby"`
	Type        string            `json:"type"`
	Status      string            `json:"status"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	ParentType  string            `json:"parenttype"`
	Parent      string            `json:"parent"`
}

func main() {
	jsonObj := []byte(`
		[
    			{
        			"ciid": "PL253020",
        			"CIname": "IE Bank IE Network PRD",
        			"uniqueCIName": "6387",
        			"environment": "PRD",
        			"managedBy": "ISP",
        			"type": "appenv",
        			"createdBy": "importappenv",
        			"timeStamp": "30/04/2018 11:01:24",
        			"status": "ACTIVE",
        			"parameters": {}
    			}
		]
	`)
	var CIObjs []*CI
	err := json.Unmarshal(jsonObj, &CIObjs)
	if err != nil {
		fmt.Println(err)
	}
	for _, ci := range CIObjs {
		fmt.Println(ci.UniqueName)
	}
}
