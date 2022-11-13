package usecases

type AuthUseCase struct {
	LoadUserByEmailRepository   interface{}
	UpdateAccessTokenRepository interface{}
	Encrypter                   interface{}
	TokenGenerator              interface{}
}

func NewAuthUseCase() (a AuthUseCase) {
	return
}

func (a AuthUseCase) Auth(email, password *string) (accessToken string, err error) {
	if *email == "" || *password == "" {
		return "", nil
	}

	accessToken = "123"
	// user := a.LoadUserByEmailRepository.load(email)
	//  isValid := user && await a.Encrypter.compare(password, user.password)
	//  if (isValid) {
	// 	 accessToken :=  a.TokenGenerator.generate(user._id)
	// 	a.UpdateAccessTokenRepository.update(user._id, accessToken)
	// 	return accessToken, nil
	//   }

	return
}
