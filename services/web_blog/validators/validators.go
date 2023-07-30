package validators

import "github.com/go-playground/validator/v10"

// specify a struct to handle errors
type ErrorResponse struct {
	FailedField string
	Tag         string
}

// add validation for Add, Update and Delete
type AuthorAddPostBody struct {
	Title string `json:"title" validate:"required"`
}

type AuthorDeletePostBody struct {
	Id int `json:"id" validate:"required"`
}

type AuthorUpdatePostBody struct {
	Id    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

var validate = validator.New()

func ValidateStruct(postBody interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(postBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element) // append value pointer (&element) --> referencing, karena variable errors adalah variable pointer slice *ErrorResponse
		}
	}
	return errors
}
