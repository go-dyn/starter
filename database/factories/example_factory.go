package factories

import "github.com/brianvoe/gofakeit/v6"

type ExampleFactory struct {
}

// Definition Define the model's default state.
func (f *ExampleFactory) Definition() map[string]any {
	return map[string]any{
		"Title":   gofakeit.LoremIpsumSentence(1),
		"Content": gofakeit.LoremIpsumSentence(1000),
	}
}
