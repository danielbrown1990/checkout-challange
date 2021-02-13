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
