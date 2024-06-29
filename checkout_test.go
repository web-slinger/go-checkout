package main

import (
	"errors"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	tests := []struct {
		name     string
		SKUs     []string
		expPrice int
		expErr   error
	}{
		// simple cases
		{
			name:     "A cost 50",
			SKUs:     []string{"A"},
			expPrice: 50,
			expErr:   nil,
		},
		{
			name:     "B cost 30",
			SKUs:     []string{"B"},
			expPrice: 30,
			expErr:   nil,
		},
		{
			name:     "C cost 20",
			SKUs:     []string{"C"},
			expPrice: 20,
			expErr:   nil,
		},
		{
			name:     "D cost 15",
			SKUs:     []string{"A"},
			expPrice: 15,
			expErr:   nil,
		},

		// multi-priced checkout
		{
			name:     "3 x A cost 130",
			SKUs:     []string{"A", "A", "A"},
			expPrice: 130,
			expErr:   nil,
		},
		{
			name:     "2 x B cost 45",
			SKUs:     []string{"B", "B"},
			expPrice: 45,
			expErr:   nil,
		},

		// non multi-priced SKUs
		{
			name:     "2 x C cost 40",
			SKUs:     []string{"C", "C"},
			expPrice: 40,
			expErr:   nil,
		},
		{
			name:     "2 x D cost 30",
			SKUs:     []string{"D", "D"},
			expPrice: 30,
			expErr:   nil,
		},

		// many SKUs w/multi-priced SKUs - order doesn't matter for multi-price SKUs
		{
			name:     "2 x B + A cost 95",
			SKUs:     []string{"B", "A", "B"},
			expPrice: 95,
			expErr:   nil,
		},
		{
			name:     "3 x A + B cost 95",
			SKUs:     []string{"A", "B", "A", "A"},
			expPrice: 160,
			expErr:   nil,
		},

		// mix SKUs
		{
			name:     "A + B + C + D cost 95",
			SKUs:     []string{"A", "B", "C", "D"},
			expPrice: 115,
			expErr:   nil,
		},

		// unhappy scenarios - ? empty checkout returns error
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pricingModel := GetPricingModel()
			checkout := NewCheckout(pricingModel)

			for i := range tc.SKUs {
				err := checkout.Scan(tc.SKUs[i])
				if err != nil {
					t.Fatalf("error when scanning SKU %s, err: %v", tc.SKUs[i], err)
				}
			}

			price, err := checkout.GetTotalPrice()
			if err != tc.expErr {
				t.Errorf("expected err '%v', got '%v'", tc.expErr, err)
			}
			if price != tc.expPrice {
				t.Errorf("expected price '%v', got '%v'", tc.expPrice, price)
			}
		})
	}
}

func TestScan(t *testing.T) {
	tests := []struct {
		name   string
		SKUs   []string
		expErr error
	}{
		{
			name:   "A SKU exists no error",
			SKUs:   []string{"A"},
			expErr: nil,
		},
		{
			name:   "B SKU exists no error",
			SKUs:   []string{"B"},
			expErr: nil,
		},
		{
			name:   "C SKU exists no error",
			SKUs:   []string{"C"},
			expErr: nil,
		},
		{
			name:   "D SKU exists no error",
			SKUs:   []string{"D"},
			expErr: nil,
		},
		{
			name:   "Z SKU not exists has error",
			SKUs:   []string{"Z"},
			expErr: errors.New("some error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pricingModel := GetPricingModel()
			checkout := NewCheckout(pricingModel)

			for i := range tc.SKUs {
				err := checkout.Scan(tc.SKUs[i])
				if err != tc.expErr {
					t.Fatalf("expected error '%v', got '%v'", tc.expErr, err)
				}
			}
		})
	}
}
