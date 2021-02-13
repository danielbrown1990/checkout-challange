package catalogue

type testCatalogue struct{}

// NewTestCatalogue returns a tests catalogue
// for in memory testing
func NewTestCatalogue() Catalogue {
	return &testCatalogue{}
}

func (c *testCatalogue) Stock() map[string]Stock {
	return map[string]Stock{
		"A": Stock{
			SKU:             "A",
			Name:            "Apples",
			SpecialPrice:    130,
			SpecialQuantity: 3,
			UnitPrice:       50,
		},
		"B": Stock{
			SKU:             "B",
			Name:            "Bananas",
			SpecialPrice:    45,
			SpecialQuantity: 2,
			UnitPrice:       30,
		},
		"C": Stock{
			SKU:       "C",
			Name:      "Cherries",
			UnitPrice: 20,
		},
		"D": Stock{
			SKU:       "D",
			Name:      "Dates",
			UnitPrice: 15,
		},
	}
}
