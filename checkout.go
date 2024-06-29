package checkout

import "fmt"

// ICheckout is the interface for methods to scan and get total price
type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

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
	// check SKU exists in pricing model
	_, ok := c.pricingModel[SKU]
	if !ok {
		return errNotFoundSKU(SKU)
	}

	// check if SKU has already been scanned
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
	for sku, quantity := range c.items {
		// check SKU exists in pricing model
		item, ok := c.pricingModel[sku]
		if !ok {
			return 0, errNotFoundSKU(sku)
		}
		// calculate price with quantity of SKUs
		totalPrice += item.CalculateSpecialPrice(quantity)
	}
	return totalPrice, nil
}

func errNotFoundSKU(SKU string) error {
	return fmt.Errorf("not found SKU '%s' in pricing model", SKU)
}
