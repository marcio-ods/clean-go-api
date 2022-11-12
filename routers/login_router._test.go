package routers

// https://www.youtube.com/watch?v=lvvPlJCMuFw&list=PL9aKtVrF05DyEwK5kdvzrYXFdpZfj1dsG&index=3
// 13 min
import (
	"fmt"
	"testing"
)

type LoginRoute struct {
}

func (l LoginRoute) route(req HttpRequest) (r HttpRequest) {

	emptyRequest := HttpRequest{}
	if emptyRequest == req {
		return HttpRequest{
			StatusCode: 500,
		}
	}

	email, password := req.Body.email, req.Body.password
	if email == "" || password == "" {
		return HttpRequest{
			StatusCode: 400,
		}
	}

	fmt.Printf("body.email: %s", email)
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

func TestLoginRoute(t *testing.T) {

	sut := LoginRoute{}

	emptyRequest := HttpRequest{}

	httpResponse := sut.route(emptyRequest)

	if httpResponse.StatusCode != 500 {
		t.Error("Should return 500 if no  httpRequest is provided")
	}

}

// func TestLoginRouteNoBody(t *testing.T) {
// 	// Erro do programador, o Go não compila com objeto inválido
// 	sut := LoginRoute{}

// 	httpResponse := sut.route(interface{})

// 	if httpResponse.StatusCode != 500 {
// 		t.Error("Should return 500 if no  httpRequest has no body")
// 	}

// }
