package service

type LoginService interface {
	LoginUser(email string, password string) bool
	GetStaticUser() (email string, password string)
}

type loginInformation struct {
	email    string
	password string
}

func NewLoginService(email string, password string) LoginService {
	return &loginInformation{
		email: email, password: password,
	}
}

func (s *loginInformation) LoginUser(email string, password string) bool {
	return s.email == email && s.password == password
}

func (s *loginInformation) GetStaticUser() (string, string) {
	return s.email, s.password
}
