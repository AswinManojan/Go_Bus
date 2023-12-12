package entities

// BusType struct is used to make the table to store bus type related information.
type BusType struct {
	Buses        Buses `gorm:"foreignKey:BusTypeCode;references:BusTypeCode"`
	BusTypeCode  string
	BusTypeName  string
	Manufacturer string
	SeatLayoutID uint
}
