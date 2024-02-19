package model

type User struct {
	Base
	UserAPI
}

type UserAPI struct {
	Fullname string `json:"fullname,omitempty" gorm:"not null"`
	Username string `json:"username,omitempty" gorm:"type:varchar(191);index:idx_users_username_unique,unique,where:deleted_at is null;not null"`
	Email    string `json:"email,omitempty" gorm:"type:varchar(191);index:idx_users_email_unique,unique,where:deleted_at is null;not null"`
	Mobile   string `json:"mobile,omitempty"`
	Password string `json:"-" gorm:"not null"`
}

type RegisterRequest struct {
	Fullname        string `json:"fullname,omitempty" example:"John doe" validate:"required"`
	Email           string `json:"email,omitempty" example:"john.doe@mail.com" validate:"required,email"`
	Password        string `json:"password,omitempty" example:"@Password123" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password,omitempty" example:"@Password123" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty" example:"john.doe@mail.com" validate:"required,email"`
	Password string `json:"password,omitempty" example:"@Password123" validate:"required,min=8"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
