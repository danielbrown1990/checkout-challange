package catalogue

import (
	"encoding/json"
	"io/ioutil"
)

// Catalogue haolds all of the items available
// represented as Stock structures
type Catalogue interface {
	Stock() map[string]Stock
}

type catalogue struct {
	stock map[string]Stock
}

// NewCatalogue returns the default implementation
// for the Catalogue interface
func NewCatalogue(catalogLocation string) (Catalogue, error) {
	data, err := ioutil.ReadFile(catalogLocation) // Could be modified here to fetch from http endpoint
	if err != nil {
		return nil, err
	}

	var stock map[string]Stock
	err = json.Unmarshal(data, &stock)
	if err != nil {
		return nil, err
	}

	c := &catalogue{
		stock: stock,
	}

	return c, nil
}

func (c *catalogue) Stock() map[string]Stock {
	return c.stock
}
