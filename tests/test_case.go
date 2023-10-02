package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/go-dyn/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
