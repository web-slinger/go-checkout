package main

// PriceScheme contains a SKU including price and an optional special price
type PriceScheme struct {
	SKU          string
	Price        int
	SpecialPrice *SpecialPrice
}

// SpecialPrice contains the SKU, quantity and price - used for multibuy discounts
type SpecialPrice struct {
	SKU      string
	Quantity int
	Price    int
}

// PricingModel is a map of SKU to PriceScheme
type PricingModel map[string]PriceScheme

// GetPricingModel returns the pricing model
func GetPricingModel() PricingModel {
	// this could be built up via 2 lists
	// list of prices keyed by SKU
	// list of special prices keyed by SKU
	return PricingModel{
		"A": {
			Price: 50,
		},
		"B": {
			Price: 30,
		},
		"C": {
			Price: 20,
		},
		"D": {
			Price: 15,
		},
	}
}
