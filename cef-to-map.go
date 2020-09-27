package main

import (
	"fmt"
	"regexp"
)

func groupmap(s string, r *regexp.Regexp) map[string]string {
	values := r.FindStringSubmatch(s)
	keys := r.SubexpNames()

	d := make(map[string]string)
	for i := 1; i < len(keys); i++ {
		d[keys[i]] = values[i]
	}

	return d
}

func main() {
	s := `time=2017-05-30T19:02:08-05:00 level=info a msg=some log message app=sample size=10 a`
	re := regexp.MustCompile(`^time=(?P<time>.*?)\slevel=(?P<level>.*?)\smsg=(?P<msg>.*?)\sapp=(?P<app>.*?)\ssize=(?P<size>.*?)$`)
	result := make(map[string]string)
	result = groupmap(s, re)
	for k, v := range result {
		fmt.Printf("%s: %s\n", k, v)
	}
}
