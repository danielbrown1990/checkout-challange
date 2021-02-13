package checkout

import (
	"checkout-challange/catalogue"
	"fmt"
)

// Checkout processes items scanned and
// returns the total of all items
type Checkout interface {
	Scan(SKU string) error
	Total() int
}

type checkout struct {
	catalogue catalogue.Catalogue
	items     map[string]int
}

// NewCheckout is a constructor for the
// default implementation of Checkout
func NewCheckout(cat catalogue.Catalogue) Checkout {
	return &checkout{
		catalogue: cat,
		items:     make(map[string]int),
	}
}

func (c *checkout) Scan(SKU string) error {
	_, ok := c.catalogue.Stock()[SKU]
	if !ok {
		return fmt.Errorf("no stock found for SKU: %v", SKU)
	}

	quantity, ok := c.items[SKU]
	if !ok {
		c.items[SKU] = 1
		return nil
	}

	quantity++
	c.items[SKU] = quantity
	return nil
}

func (c *checkout) Total() int {
	grandTotal := 0
	for SKU, quantity := range c.items {
		stock, ok := c.catalogue.Stock()[SKU]
		if !ok {
			continue
		}
		grandTotal += stock.GetPrice(quantity)
	}
	return grandTotal
}
