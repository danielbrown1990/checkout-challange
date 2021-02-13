package checkout

// Checkout processes items scanned and
// returns the total of all items
type Checkout interface {
	Scan(SKU string)
	Total() int
}
