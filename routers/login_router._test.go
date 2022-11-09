package routers

import (
	"fmt"
	"testing"
)

type LoginRoute struct {
}

func (l LoginRoute) route(req HttpRequest) (r HttpRequest) {
	body := req.Body
	if body.email == "" || body.password == "" {
		return HttpRequest{
			StatusCode: 400,
		}
	}

	fmt.Printf("body.email: %s", body.email)
	return req
}

type body struct {
	password string
	email    string
}

type HttpRequest struct {
	Body       body
	StatusCode int
}

func TestLoginRouter(t *testing.T) {

	sut := LoginRoute{}

	httpRequest := HttpRequest{
		Body: body{
			password: "123456",
		},
	}

	httpResponse := sut.route(httpRequest)

	if httpResponse.StatusCode != 400 {
		t.Error("Should return 400 if no email is provided")
	}

}
func TestPasswordIsProvided(t *testing.T) {

	sut := LoginRoute{}

	httpRequest := HttpRequest{
		Body: body{
			email: "test@test.com",
		},
	}

	httpResponse := sut.route(httpRequest)

	if httpResponse.StatusCode != 400 {
		t.Error("Should return 400 if no password is provided")
	}

}
