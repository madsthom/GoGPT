package agents

import (
	"GoGPT/llm"
	"GoGPT/promts"
)

type Agent interface {
	Run(prompt promts.PromptTemplate, input []string)
	llm.LLM
}
