package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPromptTemplate_SingleArgument(t *testing.T) {
	promptTemplate := PromptTemplate{inputVariables: []string{"product"}, template: "What is a good name for {product}"}

	prompt := promptTemplate.format([]string{"Fancy Shoes"})

	assert.Equal(t, "What is a good name for Fancy Shoes", prompt)
}

func TestPromptTemplate_MultipleArguments(t *testing.T) {
	promptTemplate := PromptTemplate{inputVariables: []string{"sentiment", "product"}, template: "What is a {sentiment} name for {product}"}

	prompt := promptTemplate.format([]string{"good", "Fancy Shoes"})

	assert.Equal(t, "What is a good name for Fancy Shoes", prompt)
}
