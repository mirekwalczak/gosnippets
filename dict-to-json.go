package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonDict := map[string]interface{}{
		"RcicDocumentId": 1,
		"FireshowerId":   1,
		"RulesStatuses": []interface{}{
			map[string]interface{}{
				"Id":      1,
				"Comment": "something went wrong",
				"Status":  true,
			},
			map[string]interface{}{
				"Id":      2,
				"Comment": "all clear",
				"Status":  false,
			},
		},
	}
	fmt.Println(jsonDict)
	jsonObj, err := json.Marshal(jsonDict)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n", reflect.TypeOf(jsonObj))
	jsonObjFormatted, err := json.MarshalIndent(jsonDict, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n", string(jsonObj))
	fmt.Println(string(jsonObjFormatted))
}
