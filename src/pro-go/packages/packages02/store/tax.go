package store

const defaultTaxRate float64 = 0.2
const minThreshold = 10

// taxRate for tax rate instances for calculating price
type taxRate struct {
    rate, threshold float64
}

// Unexported newTaxRate() creates the new *taxRate 
// inctanse for further price calculating
// Checking and setting for rate and minimal threshold
func newTaxRate(rate, threshold float64) *taxRate {
    if (rate == 0) {
        rate = defaultTaxRate
    } 
    if (threshold < minThreshold) {
        threshold = minThreshold
    }
    return &taxRate {rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) float64 {
    if (product.price > taxRate.threshold) {
        return product.price + (product.price * taxRate.rate)
    }
    return product.price
}
