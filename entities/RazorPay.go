package entities

// RazorPay Model
type RazorPay struct {
	BookID          uint    `JSON:"bookID"`
	RazorPaymentID  string  `JSON:"razorpaymentid" gorm:"primaryKey"`
	RazorPayOrderID string  `JSON:"razorpayorderid"`
	Signature       string  `JSON:"signature"`
	AmountPaid      float64 `JSON:"amountpaid"`
}
