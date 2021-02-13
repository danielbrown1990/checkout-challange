package checkout

// Checkout processes items scanned and
// returns the total of all items
type Checkout interface {
	Scan(SKU string)
	Total() int
}

type checkout struct {
	items map[string]int
}

// NewCheckout is a constructor for the
// default implementation of Checkout
func NewCheckout() Checkout {
	return &checkout{
		items: make(map[string]int),
	}
}

func (c *checkout) Scan(SKU string) {
	count, ok := c.items[SKU]
	if !ok {
		c.items[SKU] = 1
		return
	}

	count++
	c.items[SKU] = count
}

func (c *checkout) Total() int {
	return 0
}
