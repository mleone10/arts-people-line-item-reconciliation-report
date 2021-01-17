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
// TODO: Import a currency package instead, just to be safe
type Currency float64

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

	orderID, err := strconv.Atoi(rawLine[lineIndexOrderID])
	if err != nil {
		return nil, fmt.Errorf("could not parse order ID %s to string: %v", rawLine[lineIndexOrderID], err)
	}

	loc, err := time.LoadLocation(artsPeopleDateTimeZone)
	if err != nil {
		return nil, fmt.Errorf("could not create time.Location from zone %v", artsPeopleDateTimeZone)
	}
	dateTime, err := time.ParseInLocation(artsPeopleDateTimeFormat, rawLine[lineIndexDateTime], loc)
	if err != nil {
		return nil, fmt.Errorf("could not parse datetime: %v", err)
	}

	return &LineItem{
		OrderID:       orderID,
		DateTime:      dateTime,
		ItemName:      getStringIndex(rawLine, lineIndexItemName),
		Customer:      getStringIndex(rawLine, lineIndexCustomer),
		Price:         getCurrencyIndex(rawLine, lineIndexPrice),
		Fees:          getCurrencyIndex(rawLine, lineIndexFees),
		PurchaseTotal: getCurrencyIndex(rawLine, lineIndexPurchaseTotal),
		PaymentMethod: getStringIndex(rawLine, lineIndexPaymentMethod),
		GCUsed:        getStringIndex(rawLine, lineIndexGCUsed),
		Username:      getStringIndex(rawLine, lineIndexUsername),
	}, nil
}

func getStringIndex(rl []string, i int) string {
	return strings.TrimSpace(rl[i])
}

func getCurrencyIndex(rl []string, i int) Currency {
	// TODO: Implement string-to-currency parsing
	return Currency(0.0)
}
