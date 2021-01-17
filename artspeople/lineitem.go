package artspeople

import (
	"fmt"
	"strconv"
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

// TODO: Import a currency package instead, just to be safe
type currency float64

// A LineItem represents a single piece of an order, such as a ticket, membership, donation, or payment.
type LineItem struct {
	rawLine       []string
	OrderID       int
	DateTime      time.Time
	ItemName      string
	Customer      string
	Price         currency
	Fees          currency
	PurchaseTotal currency
	PaymentMethod string
	GCUsed        bool
	Username      string
}

// NewLineItem accepts an array of strings read in from the original CSV and returns an initiatlized LineItem.
// TODO: Implement remainder of rawLine parsing
func NewLineItem(rawLine []string) (*LineItem, error) {
	orderID, err := strconv.Atoi(rawLine[lineIndexOrderID])
	if err != nil {
		return nil, fmt.Errorf("Could not parse order ID %s to string: %v", rawLine[lineIndexOrderID], err)
	}

	return &LineItem{
		OrderID: orderID,
	}, nil
}
