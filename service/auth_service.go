package service

import (
	"gin-socmed/lib"
	"gin-socmed/model"
	"gin-socmed/repository"
)

type AuthService interface {
	Register(req *model.RegisterRequest) error
	Login(req *model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *model.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &lib.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.ConfirmPassword {
		return &lib.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := lib.HashPassword(req.Password)
	if err != nil {
		return &lib.InternalServerError{Message: err.Error()}
	}

	var user model.User
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Password = passwordHash

	if err := s.repository.Register(&user); err != nil {
		return &lib.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	var data model.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)

	if err != nil {
		return nil, &lib.NotFoundError{Message: err.Error()}
	}

	if err := lib.VerifiyPassword(user.Password, req.Password); err != nil {
		return nil, &lib.NotFoundError{Message: "Wrong Email or Password"}
	}

	token, err := lib.GenerateToken(user)
	if err != nil {
		return nil, &lib.InternalServerError{Message: err.Error()}
	}

	data = model.LoginResponse{
		Token: token,
		Email: user.Email,
	}

	return &data, nil
}
