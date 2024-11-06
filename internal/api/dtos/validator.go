package dtos

import (
	"context"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator interface {
	Valid(context.Context) Evaluator
}

var EmailRx = regexp.MustCompile(`^[\w\.-]+@([\w-]+\.)+[\w-]{2,4}$`)

type Evaluator map[string]string

func (e *Evaluator) AddFieldError(key, message string) {
	if *e == nil {
		*e = make(Evaluator)
	}

	if _, exists := (*e)[key]; !exists {
		(*e)[key] = message
	}
}

func (e *Evaluator) CheckField(ok bool, key, message string) {
	if !ok {
		e.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(s string, n int) bool {
	return utf8.RuneCountInString(s) <= n
}

func MinChars(s string, n int) bool {
	return utf8.RuneCountInString(s) >= n
}

func Matches(s string, rx *regexp.Regexp) bool {
	return rx.MatchString(s)
}
