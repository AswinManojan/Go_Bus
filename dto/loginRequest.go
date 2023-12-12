package dto

// LoginRequest struct is used to fetch the login input details from the user.
type LoginRequest struct {
	Email    string `json:"email" gorm:"not null" validate:"required"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}
