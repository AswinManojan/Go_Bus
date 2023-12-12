package entities

// Coupons struct is used to store information related to coupons.
type Coupons struct {
	CouponID   uint   `json:"coupon_id" gorm:"primaryKey; autoIncrement"`
	CouponCode string `json:"coupon_code" gorm:"unique" validate:"required"`
	ValidFrom  string `json:"valid_from" gorm:"not null" validate:"required"`
	ValidUpto  string `json:"valid_upto" gorm:"not null" validate:"required"`
	Discount   int    `json:"discount" gorm:"default: 10" validate:"required"`
	IsActive   bool   `json:"is_active" gorm:"default: false"`
}
