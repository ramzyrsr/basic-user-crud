package helper

import (
	"github.com/go-playground/validator"
	"lemonilo.app/model"
)

var validate *validator.Validate

func SendValidation(req *model.User) string {
	validate = validator.New()

	errValidation := validate.Struct(req)
	var sum string
	if errValidation != nil {
		var keys []string

		for _, errValidation := range errValidation.(validator.ValidationErrors) {
			keys = append(keys, errValidation.StructField())
		}

		for i := 0; i < len(keys); i++ {
			if i == len(keys)-2 || len(keys) == 1 {
				sum += keys[i] + " "
			} else if i == len(keys)-1 {
				sum += "and " + keys[i] + " "
			} else {
				sum += keys[i] + ", "
			}
		}
	}
	return sum
}
