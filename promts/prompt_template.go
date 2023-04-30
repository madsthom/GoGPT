package promts

import "strings"

type PromptTemplate struct {
	inputVariables []string
	Template       string
}

func (pt PromptTemplate) Format(input []string) string {
	result := pt.Template
	for i := 0; i < len(pt.inputVariables); i++ {
		result = strings.Replace(result, "{"+pt.inputVariables[i]+"}", input[i], 1)
	}
	return result
}
