package dto

// MakePaymentResp is used for responding to the RazorPay API.
type MakePaymentResp struct {
	BookingID      int
	AmountInRupees int
	OrderID        string
	Email          string
	PhoneNumber    string
}
