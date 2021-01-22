package artspeople

import (
	"fmt"
	"strconv"
	"strings"
)

// A Currency is a representation of money within Arts People.  Since we're parsing these from strings, we have some flexibility around how we represent these values.
type Currency int64

// NewCurrencyFromString converts a string into a cents-only representation of a USD monetary amount.  This is apparently "Martin Fowler's method".
func NewCurrencyFromString(s string) (Currency, error) {
	if s == "" {
		return Currency(0), nil
	}

	// From https://stackoverflow.com/a/51660442
	// Split the raw currency string into, at most, 3 parts.  If, however, there aren't two parts or if the "cents" part isn't two digits, throw an error.
	n := strings.SplitN(s, ".", 3)
	if len(n) != 2 || len(n[1]) != 2 {
		return 0, fmt.Errorf("split currency field [%v]appears incorrect", s)
	}

	// Parse the "dollars" part into an integer of at most 56 bits.
	d, err := strconv.ParseInt(n[0], 10, 56)
	if err != nil {
		return 0, fmt.Errorf("failed to parse dollars part of currency [%v]: %v", s, err)
	}

	// Parse the "cents" part into an integer of at most 8 bits.
	c, err := strconv.ParseUint(n[1], 10, 8)
	if err != nil {
		return 0, fmt.Errorf("failed to parse cents part of currency [%v]: %v", s, err)
	}

	// If the dollars part is negative, also negate the cents part.
	if d < 0 {
		c = -c
	}

	return Currency(d*100 + int64(c)), nil
}
