package usecases

import (
	"victorLessa/server/application/repositories"
	"victorLessa/server/domain"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Code int	`json:"code"`
}

func (m *ErrorMessage) Error() string {
    return "Error"
}

type AuthenticateUseCases struct {
	UserRepository repositories.UserRepositoryDb
}

func (u *AuthenticateUseCases) Authenticate(username string, password string) (*domain.User, error) {
	res := u.UserRepository.FindByName(username)
	
	verify := domain.VerifyPassword(res.Password, password)
	
	if res.ID == "" || !verify {
		arrayEmpty := &ErrorMessage{
			Message: "Usuário ou senha incorreto",
			Code: 401,
		}
		return nil, arrayEmpty
	}

	token, err := domain.CreateToken(res)

	if err != nil {
		return nil, err
	}

	res.Token = token
	return res, nil
}