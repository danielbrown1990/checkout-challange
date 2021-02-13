package checkout

// Checkout processes items scanned and
// returns the total of all items
type Checkout interface {
	Scan(SKU string)
	Total() int
}

type checkout struct{}

// NewCheckout is a constructor for the
// default implementation of Checkout
func NewCheckout() Checkout {
	return &checkout{}
}

func (c *checkout) Scan(SKU string) {

}

func (c *checkout) Total() int {
	return 0
}
