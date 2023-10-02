package seeders

import (
	"github.com/go-dyn/app/models"
	"github.com/goravel/framework/facades"
)

type ExampleSeeder struct {
	len int
}

// Signature The name and signature of the seeder.
func (s *ExampleSeeder) Signature() string {
	return "ExampleSeeder"
}

// Run executes the seeder logic.
func (s *ExampleSeeder) Run() error {
	var examples []models.Example
	facades.Orm().Factory().Count(s.len).Create(&examples)
	return nil
}
