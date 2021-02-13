package catalogue

// Stock is a struct to hold information about
// a single item in a catalogue
type Stock struct {
	Name            string
	SKU             string
	UnitPrice       int
	SpecialQuantity int
	SpecialPrice    int
}

// GetPrice will return the total amount for a
// specified quantity accounting for offers
func (s Stock) GetPrice(quantity int) int {
	if s.SpecialQuantity != 0 && quantity >= s.SpecialQuantity { //Special offer applicable
		additional := 0 //any additional items not included in offer
		if quantity%s.SpecialQuantity != 0 {
			additional = quantity % s.SpecialQuantity * s.UnitPrice
		}

		return ((quantity / s.SpecialQuantity) * s.SpecialPrice) + additional
	}

	return s.UnitPrice * quantity
}
