package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "SISKAEEE"
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
	fmt.Println(strings.Join(joinStr, ""))

}
