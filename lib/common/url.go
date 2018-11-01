package common

import "strings"

func JoinURL(inputs ...string) string {
	trimSep := "/ *-&^$"
	joinSep := "/"
	var outputs []string
	for _, e := range inputs {
		if trimed := strings.Trim(e, trimSep); trimed != "" {
			outputs = append(outputs, trimed)
		}
	}
	return strings.Join(outputs, joinSep)
}
