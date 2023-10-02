package factories

import "github.com/brianvoe/gofakeit/v6"

type UserFactory struct {
}

// Definition Define the model's default state.
func (f *UserFactory) Definition() map[string]any {
	return map[string]any{
		"Name":       gofakeit.Name(),
		"Email":      gofakeit.Email(),
		"Identifier": gofakeit.Email(),
		"UserName":   gofakeit.Email(),
		"password":   "password",
	}
}
