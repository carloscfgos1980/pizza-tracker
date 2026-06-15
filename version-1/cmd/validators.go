package main

import (
	"slices"

	"github.com/carloscfgos1980/pizza-tracker/internal/models"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pizza_valid_type", createSliceValidator(models.PizzaTypes))
		v.RegisterValidation("pizza_valid_size", createSliceValidator(models.PizzaSizes))
	}
}

func createSliceValidator(validValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return slices.Contains(validValues, fl.Field().String())
	}
}
