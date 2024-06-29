package main

import "fmt"

// Checkout contains the methods to scan and get total price
type Checkout struct {
	items        map[string]int // items is a map of SKU and quantity
	pricingModel PricingModel   // pricingModel contains a map of SKU to price scheme
}

// NewCheckout returns an instance of Checkout
func NewCheckout(pricingModel PricingModel) *Checkout {
	return &Checkout{
		items:        map[string]int{},
		pricingModel: pricingModel,
	}
}

// Scan adds an item to the checkout using the SKU, if there's an issue returns an error
func (c *Checkout) Scan(SKU string) (err error) {
	_, ok := c.pricingModel[SKU]
	if !ok {
		return errNotFoundSKU(SKU)
	}

	_, found := c.items[SKU]
	if !found {
		// init key in map
		c.items[SKU] = 0
	}

	// add quantity for SKU
	c.items[SKU]++

	return nil
}

// GetTotalPrice returns the total price after summing up items in checkout, if there's an issue returns an error and 0 price
func (c *Checkout) GetTotalPrice() (totalPrice int, err error) {
	for sku := range c.items {
		item, ok := c.pricingModel[sku]
		if !ok {
			return 0, errNotFoundSKU(sku)
		}
		totalPrice += item.Price
	}
	return totalPrice, nil
}

func errNotFoundSKU(SKU string) error {
	return fmt.Errorf("not found SKU '%s' in pricing model", SKU)
}
