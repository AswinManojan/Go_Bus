package entities

type ServiceProvider struct {
	ProviderID  uint   `json:"providerid" gorm:"primaryKey; autoIncrement"`
	Email       string `json:"email" gorm:"unique"`
	CompanyName string `json:"company" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	Role        string `json:"role" gorm:"default: 'user'"`
	PhoneNumber string `json:"phone" gorm:"not null"`
	BusCount    uint   `json:"bus_count" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	IsLocked    bool   `json:"is_account_locked" gorm:"default: true"`
}
