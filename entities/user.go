package entities

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email       string `json:"email" gorm:"unique"`
	UserName    string `json:"username" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	Role        string `json:"role" gorm:"default: 'user'"`
	PhoneNumber string `json:"phone" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	DOB         string `json:"dob" gorm:"not null"`
	IsLocked    bool   `json:"is_account_locked" gorm:"default: false"`
}
