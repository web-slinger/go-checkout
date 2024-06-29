package checkout

// PriceScheme contains a SKU including price and an optional special price
type PriceScheme struct {
	SKU          string
	Price        int
	SpecialPrice *SpecialPrice
}

// CalculateSpecialPrice returns the price for a given quantity of a known SKU
func (p *PriceScheme) CalculateSpecialPrice(quantity int) (price int) {
	// if no special price
	if p.SpecialPrice == nil {
		return p.Price * quantity
	}

	// quantity of SKU is multiple of special price quantity
	specialPriceCount := quantity / p.SpecialPrice.Quantity

	// remainder count that will be charged at normal price
	remainderCount := quantity % p.SpecialPrice.Quantity

	return specialPriceCount*p.SpecialPrice.Price + (remainderCount * p.Price)
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
			SpecialPrice: &SpecialPrice{
				Quantity: 3,
				Price:    130,
			},
		},
		"B": {
			Price: 30,
			SpecialPrice: &SpecialPrice{
				Quantity: 2,
				Price:    45,
			},
		},
		"C": {
			Price: 20,
		},
		"D": {
			Price: 15,
		},
	}
}
