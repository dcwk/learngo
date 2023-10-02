package patterns

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

type FlightReservationImpl struct {
	reservationDate string
}

func (r FlightReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 1.0
}

type HotelReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 1.0
}
