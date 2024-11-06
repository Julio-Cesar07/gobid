package dtos

import (
	"context"
	"fmt"
	"regexp"
)

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) Evaluator {
	var eval Evaluator

	// valid stuff
	eval.CheckField(NotBlank(req.Username), "username", "this field cannot be empty")
	eval.CheckField(MaxChars(req.Bio, 255), "username", "this field must have a maximum length of 255")
	eval.CheckField(NotBlank(req.Email), "email", "this field cannot be empty")
	fmt.Println(req.Email)
	eval.CheckField(Matches(req.Email, EmailRx), "email", "this field must be a valid email")
	eval.CheckField(
		MinChars(req.Password, 8) && MaxChars(req.Password, 24),
		"password", "this field must have a length between 8 and 24")
	eval.CheckField(
		Matches(req.Password, regexp.MustCompile("[A-Z]")),
		"password", "this field with at least a capital letter")
	eval.CheckField(
		Matches(req.Password, regexp.MustCompile(`\d`)),
		"password", "this field with at least a number")
	eval.CheckField(
		Matches(req.Password, regexp.MustCompile(`[@$!%*?&]`)),
		"password", "this field with at least a special character")

	return eval
}
