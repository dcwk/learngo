package patterns

// Factory method
func NewReservation(vertical, reservationDate string) Reservation {
	switch vertical {
	case "flight":
		return FlightReservationImpl{reservationDate}
	case "hotel":
		return HotelReservationImpl{reservationDate}
	default:
		return nil
	}
}

hotelReservation := NewReservation("hotel", "20180101")
hotelReservation.GetReservationDate()
