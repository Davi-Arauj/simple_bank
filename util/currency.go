package util

const (
	USD = "USD"
	EUR = "EUR"
	BR  = "BR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, BR:
		return true
	}
	return false
}
