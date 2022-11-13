package routers

// https://www.youtube.com/watch?v=lvvPlJCMuFw&list=PL9aKtVrF05DyEwK5kdvzrYXFdpZfj1dsG&index=3
// 13 min
import (
	"github.com/marcio-ods/clean-go-api/domain/usecases"
	"github.com/marcio-ods/clean-go-api/presentation/helpers"
	"testing"
)

var makeSut = func() (LoginRouter, usecases.AuthUseCase) {
	authUseCase := usecases.NewAuthUseCase()
	sut := NewLoginRouter(authUseCase)
	return sut, authUseCase
}

// func TestLoginRouter500(t *testing.T) {

// 	sut := NewLoginRouter()

// 	emptyRequest := helpers.NewHttpLoginRequest()

// 	httpResponse := sut.Route(emptyRequest)

//		if httpResponse.StatusCode != 400 {
//			t.Errorf("Should return 500 if no httpRequest is provided: %s ", httpResponse.Detail)
//		}
//	}
func TestLoginRouter(t *testing.T) {

	sut, _ := makeSut()

	emptyRequest := helpers.NewHttpLoginRequest()

	httpResponse := sut.Route(emptyRequest)

	if httpResponse.StatusCode != 400 {
		t.Errorf("Should return 400 if no email is provided: %s ", httpResponse.Detail)
	}
}

func TestPasswordIsProvided(t *testing.T) {

	sut, _ := makeSut()

	httpRequest := helpers.HttpLoginRequest{
		Email: "test@test.com",
	}

	httpResponse := sut.Route(httpRequest)

	if httpResponse.StatusCode != 400 {
		t.Errorf("Should return 400 if no password is provided: %s ", httpResponse.Detail)
	}

}
func TestEmailIsProvided(t *testing.T) {
	sut, _ := makeSut()
	httpRequest := helpers.HttpLoginRequest{
		Password: "test@test.com",
	}
	httpResponse := sut.Route(httpRequest)
	if httpResponse.StatusCode != 400 {
		t.Errorf("Should return 400 if no email is provided: %s ", httpResponse.Detail)
	}
}

func TestCredentials(t *testing.T) {
	sut, _ := makeSut()
	httpRequest := helpers.HttpLoginRequest{
		Password: "test@test.com",
		Email:    "test@test.com",
	}
	httpResponse := sut.Route(httpRequest)
	if httpResponse.StatusCode == 401 {
		t.Error("Should return 401 when invalid credentials are provided")
	}
}

// func TestAuthUseCaseWhitCorrectParams(t *testing.T) {
// 	sut, authUseCase := makeSut()
// 	httpRequest := helpers.HttpLoginRequest{
// 		Password: "123456",
// 		Email:    "test@test.com",
// 	}
// 	sut.Route(httpRequest)

// 	if authUseCase.Email != httpRequest.Email {
// 		t.Error("TestAuthUseCaseWhitCorrectParams")
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

// func TestLoginRouteNoBody(t *testing.T) {
// 	// Erro do programador, o Go não compila com objeto inválido
// 	sut := LoginRoute{}

// 	httpResponse := sut.route(interface{})

// 	if httpResponse.StatusCode != 500 {
// 		t.Error("Should return 500 if no  httpRequest has no body")
// 	}

// }
