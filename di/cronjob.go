package di

import (
	"gobus/services/interfaces"
	"time"
)

// CouponValidator is used to validate the existing coupons
func CouponValidator(ps interfaces.ProviderService) {
	coupons, _ := ps.FindCoupon()
	for _, coupon := range coupons {
		parsedTimeUpto, _ := time.Parse("02012006", coupon.ValidUpto)
		parsedTimeFrom, _ := time.Parse("02012006", coupon.ValidFrom)
		if parsedTimeUpto.Before(time.Now()) || parsedTimeFrom.After(time.Now()) {
			ps.DeactivateCoupon(int(coupon.CouponID))
		} else if parsedTimeUpto.After(time.Now()) && parsedTimeFrom.Before(time.Now()) {
			ps.ActivateCoupon(int(coupon.CouponID))
		}
	}
}
