package entities

type Coupons struct {
	CouponId   uint   `json:"coupon_id" gorm:"primaryKey; autoIncrement"`
	CouponCode string `json:"coupon_code" gorm:"unique"`
	ValidFrom  string `json:"valid_from" gorm:"not null"`
	ValidUpto  string `json:"valid_upto" gorm:"not null"`
	Discount   int    `json:"discount" gorm:"default: 10"`
	IsActive   bool   `json:"is_active" gorm:"default: false"`
}
