package entities

// ServiceProvider struct is used to store the details of bus service provider.
type ServiceProvider struct {
	// Buses          Buses  `gorm:"foreignKey:ProviderID;references:ProviderID"`
	ProviderID     uint   `json:"providerid" gorm:"primaryKey; autoIncrement"`
	Email          string `json:"email" gorm:"unique" validate:"required"`
	CompanyName    string `json:"company" gorm:"not null" validate:"required"`
	Password       string `json:"password" gorm:"not null" validate:"required"`
	Role           string `json:"role" gorm:"default: 'provider'"`
	PhoneNumber    string `json:"phone" gorm:"not null" validate:"required"`
	BusCount       uint   `json:"bus_count"`
	Address        string `json:"address" gorm:"not null" validate:"required"`
	IsLocked       bool   `json:"is_account_locked" gorm:"default: true"`
	ProviderWallet int    `json:"provider_wallet"`
}
