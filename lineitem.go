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
	itemNamePayment          = "Payment"
)

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
	GCUsed        string // Maybe not used?
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
	price, err := NewCurrencyFromString(getStringIndex(rawLine, lineIndexPrice))
	if err != nil {
		return nil, fmt.Errorf("could not parse price: %v", err)
	}

	// Parse the item's fees
	fees, err := NewCurrencyFromString(getStringIndex(rawLine, lineIndexFees))
	if err != nil {
		return nil, fmt.Errorf("could not parse fees: %v", err)
	}

	// Parse the order's purchase total
	purchaseTotal, err := NewCurrencyFromString(getStringIndex(rawLine, lineIndexPurchaseTotal))
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

// IsPayment returns true if the given LineItem pertains to a payment.
func (li *LineItem) IsPayment() bool {
	return li.ItemName == itemNamePayment
}
