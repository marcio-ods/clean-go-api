package routers

import (
	"github.com/marcio-ods/clean-go-api/domain/usecases"
	"github.com/marcio-ods/clean-go-api/presentation/helpers"
)

type LoginRouter struct {
	Email          string
	Password       string
	AuthUseCase    usecases.AuthUseCase
	EmailValidator interface{}
}

func NewLoginRouter(auth usecases.AuthUseCase) (l LoginRouter) {
	l.AuthUseCase = auth
	return
}

func (l *LoginRouter) Route(req helpers.HttpLoginRequest) (res helpers.HttpResponse) {
	email, password := req.Email, req.Password
	if email == "" || password == "" {
		res.BadRequest("email and password required")
		// err = errors.New(NewMissingParamError("email"))
		return
	}

	accessToken, err := l.AuthUseCase.Auth(&email, &password)
	if err != nil {
		res.UnauthorizedError("invalid credentials")
		// err = errors.New(NewMissingParamError("email"))
		return
	}
	println(accessToken)
	return
}
