package helpers

type HttpLoginRequest struct {
	Password string
	Email    string
}

func NewHttpLoginRequest() HttpLoginRequest {
	return HttpLoginRequest{}
}
