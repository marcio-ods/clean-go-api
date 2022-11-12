package routers

import (
	"github.com/marcio-ods/clean-go-api/presentation/helpers"
)

type LoginRouter struct {
	Email    string
	Password string
}

func NewLoginRouter() (l LoginRouter) {
	return
}

func (l LoginRouter) Route(req helpers.HttpLoginRequest) (res helpers.HttpResponse) {
	email, password := req.Email, req.Password
	if email == "" || password == "" {
		res.BadRequest("email and password required")
		// res.BedRequest()
		// err = errors.New(NewMissingParamError("email"))
		return
	}
	return
}
