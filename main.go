package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "SISKAEEE"
	result := concateString(plainText)
	fmt.Println(result)
}

func concateString(plainText string) (result string) {
	textMap := map[string][]string{}
	for _, p := range plainText {
		if _, exist := textMap[string(p)]; !exist {
			textMap[string(p)] = []string{}
		}
		textMap[string(p)] = append(textMap[string(p)], string(p))
	}

	joinStr := []string{}
	for k, v := range textMap {
		length := len(v)
		preOutput := fmt.Sprintf("%d%s", length, k)
		joinStr = append(joinStr, preOutput)
	}
	result = strings.Join(joinStr, "")
	return
}
