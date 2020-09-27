package main

import (
	"fmt"
	"reflect"
)

type FWLog struct {
	Date          string
	DeviceID      string
	DeviceIP      string
	OperationType string
	User          string
	RuleBefore    string
	RuleAfter     string
	Action        string
}

func (f *FWLog) Fields() []string {
	var fields []string
	s := reflect.ValueOf(f).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fields = append(fields, typeOfT.Field(i).Name)
	}
	return fields
}

func main() {
	f := new(FWLog)
	fmt.Println(f.Fields())

}
