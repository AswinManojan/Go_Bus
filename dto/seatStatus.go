package dto

// SeatAvailabilityResponse is used to provide the response based on the seat availability as per the busID and date shared by the user.
type SeatAvailabilityResponse struct {
	BusID                 int
	Date                  string
	BusType               string
	BusStatus             string
	SleeperSlotsLeft      int
	SeaterSlotsLeft       int
	AvailableSleeperSlots []string
	AvailableSeaterSlots  []string
}

// SeatAvailabilityRequest is used to accept input from the user inorder to check the seat availability.
type SeatAvailabilityRequest struct {
	BusID int    `json:"bus_id" gorm:"not null" validate:"required"`
	Date  string `json:"date" gorm:"not null" validate:"required"`
}
