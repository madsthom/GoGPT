package agents

import (
	"GoGPT/llm"
	"GoGPT/promts"
)

type ChatAgent struct {
	llm.LLM
}

func (a ChatAgent) Run(prompt promts.PromptTemplate, input []string) {
	chatMsg := prompt.Format(input)
	resp, err := a.Predict(chatMsg)
	if err {
		println(err)
		return
	}

	println(resp.Messages[0])
}
