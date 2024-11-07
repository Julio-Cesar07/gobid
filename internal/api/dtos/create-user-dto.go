package dtos

import (
	"context"
	"regexp"
)

type CreateUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (dto CreateUserDto) Valid(ctx context.Context) Evaluator {
	var eval Evaluator

	// valid stuff
	eval.CheckField(NotBlank(dto.Username), "username", "this field cannot be empty")
	eval.CheckField(MaxChars(dto.Bio, 255), "username", "this field must have a maximum length of 255")
	eval.CheckField(NotBlank(dto.Email), "email", "this field cannot be empty")
	eval.CheckField(Matches(dto.Email, EmailRx), "email", "this field must be a valid email")
	eval.CheckField(
		MinChars(dto.Password, 8) && MaxChars(dto.Password, 24),
		"password", "this field must have a length between 8 and 24")
	eval.CheckField(
		Matches(dto.Password, regexp.MustCompile("[A-Z]")),
		"password", "this field with at least a capital letter")
	eval.CheckField(
		Matches(dto.Password, regexp.MustCompile(`\d`)),
		"password", "this field with at least a number")
	eval.CheckField(
		Matches(dto.Password, regexp.MustCompile(`[@$!%*?&]`)),
		"password", "this field with at least a special character")

	return eval
}
