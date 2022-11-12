package routers

// // https://www.youtube.com/watch?v=lvvPlJCMuFw&list=PL9aKtVrF05DyEwK5kdvzrYXFdpZfj1dsG&index=3
// // 13 min
// import (
// 	// "errors"
// 	"fmt"
// 	"testing"
// )

// type MissingParamError struct {
// 	name string
// }

// func NewMissingParamError(params string) (resp string) {
// 	e := MissingParamError{}
// 	e.name = "MissingParamError"

// 	resp = fmt.Sprintf("%s: %s", e.name, params)

// 	return
// }

// type HttpResponse struct {
// 	StatusCode int
// 	body       MissingParamError
// }

// func (h *HttpResponse) BedRequest(paramName string) {
// 	h.StatusCode = 400
// 	// h.body = NewMissingParamError(paramName)
// }

// func (h *HttpResponse) ServerError() {
// 	h.StatusCode = 500
// }

// type body struct {
// 	password string
// 	email    string
// }

// type HttpRequest struct {
// 	Body body
// 	// StatusCode int
// }

// func TestLoginRouter(t *testing.T) {

// 	sut := LoginRoute{}

// 	httpRequest := HttpRequest{
// 		Body: body{
// 			password: "123456",
// 		},
// 	}

// 	httpResponse, err := sut.route(httpRequest)

// 	if httpResponse.StatusCode != 400 {
// 		t.Errorf("Should return 400 if no email is provided: %s ", err)
// 	}
// }
// func TestPasswordIsProvided(t *testing.T) {

// 	sut := LoginRoute{}

// 	httpRequest := HttpRequest{
// 		Body: body{
// 			email: "test@test.com",
// 		},
// 	}

// 	httpResponse, err := sut.route(httpRequest)

// 	if httpResponse.StatusCode != 400 {
// 		t.Errorf("Should return 400 if no password is provided: %s ", err)
// 	}

// }

// func TestLoginRoute(t *testing.T) {

// 	sut := LoginRoute{}

// 	emptyRequest := HttpRequest{}

// 	httpResponse, err := sut.route(emptyRequest)

// 	if httpResponse.StatusCode != 500 {
// 		t.Errorf("Should return 500 if no  httpRequest is provided: %s ", err)
// 	}

// }

// // func TestLoginRouteNoBody(t *testing.T) {
// // 	// Erro do programador, o Go não compila com objeto inválido
// // 	sut := LoginRoute{}

// // 	httpResponse := sut.route(interface{})

// // 	if httpResponse.StatusCode != 500 {
// // 		t.Error("Should return 500 if no  httpRequest has no body")
// // 	}

// // }
