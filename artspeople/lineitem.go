package artspeople

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	lineIndexOrderID = iota
	lineIndexDateTime
	lineIndexItemName
	lineIndexCustomer
	lineIndexPrice
	lineIndexFees
	lineIndexPurchaseTotal
	lineIndexPaymentMethod
	lineIndexGCUsed
	lineIndexUsername
)

const (
	numExpectedFields        = 10
	artsPeopleDateTimeFormat = "2006-01-02 03:04 PM"
	artsPeopleDateTimeZone   = "America/New_York"
)

// A Currency is a representation of money within Arts People.  Since we're parsing these from strings, we have some flexibility around how we represent these values.
type Currency int64

// A LineItem represents a single piece of an order, such as a ticket, membership, donation, or payment.
type LineItem struct {
	rawLine       []string
	OrderID       int
	DateTime      time.Time
	ItemName      string
	Customer      string
	Price         Currency
	Fees          Currency
	PurchaseTotal Currency
	PaymentMethod string
	GCUsed        string // This field does not appear to be used.  As such, we leave it as a string until test data exists.
	Username      string
}

// NewLineItem accepts an array of strings read in from the original CSV and returns an initiatlized LineItem.
func NewLineItem(rawLine []string) (*LineItem, error) {
	if len(rawLine) != numExpectedFields {
		return nil, fmt.Errorf("raw line contained %d fields, but only found %d", numExpectedFields, len(rawLine))
	}

	// Parse the order ID
	orderID, err := strconv.Atoi(rawLine[lineIndexOrderID])
	if err != nil {
		return nil, fmt.Errorf("could not parse order ID: %v", err)
	}

	// Parse the item datetime
	loc, err := time.LoadLocation(artsPeopleDateTimeZone)
	if err != nil {
		return nil, fmt.Errorf("could not create time.Location from zone %v", artsPeopleDateTimeZone)
	}
	dateTime, err := time.ParseInLocation(artsPeopleDateTimeFormat, getStringIndex(rawLine, lineIndexDateTime), loc)
	if err != nil {
		return nil, fmt.Errorf("could not parse datetime: %v", err)
	}

	// Parse the item price
	price, err := getCurrencyIndex(rawLine, lineIndexPrice)
	if err != nil {
		return nil, fmt.Errorf("could not parse price: %v", err)
	}

	// Parse the item's fees
	fees, err := getCurrencyIndex(rawLine, lineIndexFees)
	if err != nil {
		return nil, fmt.Errorf("could not parse fees: %v", err)
	}

	// Parse the order's purchase total
	purchaseTotal, err := getCurrencyIndex(rawLine, lineIndexPurchaseTotal)
	if err != nil {
		return nil, fmt.Errorf("could not parse purchase total: %v", err)
	}

	return &LineItem{
		OrderID:       orderID,
		DateTime:      dateTime,
		ItemName:      getStringIndex(rawLine, lineIndexItemName),
		Customer:      getStringIndex(rawLine, lineIndexCustomer),
		Price:         price,
		Fees:          fees,
		PurchaseTotal: purchaseTotal,
		PaymentMethod: getStringIndex(rawLine, lineIndexPaymentMethod),
		GCUsed:        getStringIndex(rawLine, lineIndexGCUsed),
		Username:      getStringIndex(rawLine, lineIndexUsername),
	}, nil
}

func getStringIndex(rl []string, i int) string {
	return strings.TrimSpace(rl[i])
}

func getCurrencyIndex(rl []string, i int) (Currency, error) {
	s := getStringIndex(rl, i)
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
