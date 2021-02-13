package catalogue

// Catalogue haolds all of the items available
// represented as Stock structures
type Catalogue interface {
	Stock() []Stock
}

type catalogue struct{}

// NewCatalogue returns the default implementation
// for the Catalogue interface
func NewCatalogue() Catalogue {
	return &catalogue{}
}

func (c *catalogue) Stock() []Stock {
	return []Stock{
		Stock{
			SKU:             "A",
			Name:            "Apples",
			SpecialPrice:    130,
			SpecialQuantity: 3,
			UnitPrice:       50,
		},
		Stock{
			SKU:             "B",
			Name:            "Bananas",
			SpecialPrice:    45,
			SpecialQuantity: 2,
			UnitPrice:       30,
		},
		Stock{
			SKU:       "C",
			Name:      "Cherries",
			UnitPrice: 20,
		},
		Stock{
			SKU:       "D",
			Name:      "Dates",
			UnitPrice: 15,
		},
	}
}
