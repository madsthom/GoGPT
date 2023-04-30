package promts

import "strings"

type PromptTemplate struct {
	inputVariables []string
	template       string
}

func (pt PromptTemplate) format(input []string) string {
	result := pt.template
	for i := 0; i < len(pt.inputVariables); i++ {
		result = strings.Replace(result, "{"+pt.inputVariables[i]+"}", input[i], 1)
	}
	return result
}
