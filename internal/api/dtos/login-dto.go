package dtos

import "context"

type LoginDto struct {
	Email    string
	Password string
}

func (dto LoginDto) Valid(ctx context.Context) Evaluator {
	var eval Evaluator

	eval.CheckField(NotBlank(dto.Email), "email", "this field cannot be empty")
	eval.CheckField(Matches(dto.Email, EmailRx), "email", "this field must be an valid email")
	eval.CheckField(NotBlank(dto.Password), "password", "this field cannot be empty")
	eval.CheckField(
		MinChars(dto.Password, 8) && MaxChars(dto.Password, 24),
		"password", "this field must have a lenght between 8 and 24",
	)

	return eval
}
