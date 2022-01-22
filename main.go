package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "SISKAEEE"
	builder := strings.Builder{}
	result := concateString(plainText, builder)

	// concate with strings.builder
	builder.WriteString(fmt.Sprintf("- the output result with string.builder: %s\n", result))
	output1 := builder.String()

	// concate with fmt.sprintf
	output2 := fmt.Sprintf("- the output result with fmt.sprintf: %s", result)

	fmt.Println(output1)
	fmt.Println(output2)

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
		b.WriteString(fmt.Sprintf("%d%s", len(v), k))
		joinStr = append(joinStr, b.String())
	}
	result = strings.Join(joinStr, "")
	return
}
