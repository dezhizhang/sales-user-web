package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Name string `validate:"min=6,max=10"`
	Age  int    `validate:"min=1,max=100"`
}

func TestValidator(t *testing.T) {
	validate := validator.New()

	u1 := User{Name: "lidajun", Age: 18}
	err := validate.Struct(u1)
	fmt.Println(err)

	u2 := User{Name: "dj", Age: 101}
	err = validate.Struct(u2)
	fmt.Println(err)
}
