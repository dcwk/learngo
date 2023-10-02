package patterns

type AbstractFactory interface {
	CreateReservation() Reservation
	CreateInvoice() Invoice
}

type HotelFactory struct{}

func (f HotelFactory) CreateReservation() Reservation {
	return new(HotelReservationImpl)
}

func (f HotelFactory) CreateInvoice() Invoice {
	return new(HotelInvoice)
}

type FlightFactory struct{}

func (f FlightFactory) CreateReservation() Reservation {
	return new(FlightReservationImpl)
}

func (f FlightFactory) CreateInvoice() Invoice {
	return new(FlightInvoice)
}

func GetFactory(vertical string) AbstractFactory {
	var factory AbstractFactory
	switch vertical {
	case "hotel":
		factory = HotelFactory{}
	case "flight":
		factory = FlightFactory{}
	}

	return factory
}

//hotelFactory = GetFactory("hotel")
//hotelFactory.CreateReservation()
//hotelFactory.CreateInvoice()
