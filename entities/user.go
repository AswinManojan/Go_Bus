package entities

// User struct is used to store the informations of User data
type User struct {
	// PassengerInfo PassengerInfo `gorm:"foreignKey:UserID;references:ID"`
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email       string `json:"email" gorm:"unique" validate:"required"`
	UserName    string `json:"username" gorm:"not null" validate:"required"`
	Password    string `json:"password" gorm:"not null" validate:"required"`
	Role        string `json:"role" gorm:"default: 'user'"`
	PhoneNumber string `json:"phone" gorm:"not null" validate:"required"`
	Gender      string `json:"gender" gorm:"not null" validate:"required"`
	DOB         string `json:"dob" gorm:"not null" validate:"required"`
	IsLocked    bool   `json:"is_account_locked" gorm:"default: false"`
	UserWallet  int    `json:"user_wallet"`
}
