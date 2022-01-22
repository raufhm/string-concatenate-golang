package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "SISKAEEE"
	builder := strings.Builder{}
	result := concateString(plainText, builder)
	builder.WriteString(fmt.Sprintf("the output result: %s", result))
	msg := builder.String()
	fmt.Println(msg)
}

func concateString(plainText string, builder strings.Builder) (result string) {
	textMap := map[string][]string{}
	for _, p := range plainText {
		if _, exist := textMap[string(p)]; !exist {
			textMap[string(p)] = []string{}
		}
		textMap[string(p)] = append(textMap[string(p)], string(p))
	}

	joinStr := []string{}
	for k, v := range textMap {
		b := builder
		length := len(v)
		b.WriteString(fmt.Sprintf("%d%s", length, k))
		joinStr = append(joinStr, b.String())
	}
	result = strings.Join(joinStr, "")
	return
}
