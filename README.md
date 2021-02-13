# checkout-challange

Implementation of checkout challange

## Testing

All test cases are within ```CheckoutChallange_test.go``` these demonstrate adding a veriety of different items to the checkout, the tests use a ```testCatalogue``` which returns a static hardcoded catalogue for unit testing.

## CLI

Provided is a small CLI which implements retreiving the catalogue from a JSON file from the file system, this file can be modified and the application reloaded to update the catalogue without recomiling, future addtional features may include a ```Reload()``` method to reload the catalogue without restarting the application and fetching the catalogue from an exteral API instead of the file system.