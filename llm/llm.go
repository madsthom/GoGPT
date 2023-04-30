package llm

type Response struct {
	Messages []string
}

type LLM interface {
	Predict(msg string) (Response, bool)
}
