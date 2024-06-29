package main

// Checkout contains the methods to scan and get total price
type Checkout struct {
	pricingModel PricingModel // pricingModel contains a map of SKU to price scheme
}

// NewCheckout returns an instance of Checkout
func NewCheckout(pricingModel PricingModel) *Checkout {
	return &Checkout{
		pricingModel: pricingModel,
	}
}

// Scan adds an item to the checkout using the SKU, if there's an issue returns an error
func (c *Checkout) Scan(SKU string) (err error) {
	return nil
}

// GetTotalPrice returns the total price after summing up items in checkout, if there's an issue returns an error and 0 price
func (c *Checkout) GetTotalPrice() (totalPrice int, err error) {
	return 0, nil
}
