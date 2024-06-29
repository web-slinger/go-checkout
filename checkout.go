package main

import "fmt"

// Checkout contains the methods to scan and get total price
type Checkout struct {
	items        []string     // items is a list of SKUs that have been successfully scanned
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
	_, ok := c.pricingModel[SKU]
	if !ok {
		return errNotFoundSKU(SKU)
	}

	c.items = append(c.items, SKU)

	return nil
}

// GetTotalPrice returns the total price after summing up items in checkout, if there's an issue returns an error and 0 price
func (c *Checkout) GetTotalPrice() (totalPrice int, err error) {
	return 0, nil
}

func errNotFoundSKU(SKU string) error {
	return fmt.Errorf("not found SKU '%s' in pricing model", SKU)
}
