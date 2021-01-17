package artspeople

import (
	"encoding/csv"
	"fmt"
	"io"
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

// A LineItemReconReport is a parsed and type-normalized version of the Line Item Reconciliation Report downloaded from Arts People.
type LineItemReconReport struct {
	rawLines [][]string
	Orders   map[int]*Order
}

// An Order represents all details of a single interaction with a customer, including all items purchased and the payment method.
// TODO: Consider converting from a struct to a type alias ([]*LineItem)
type Order struct {
	LineItems []*LineItem
}

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

// NewLineItemReconReport accepts an Arts People Line Item Reconciliation Report as an io.Reader and returns a fully parsed and type-normalized LineItemReconReport.
func NewLineItemReconReport(reportCsv io.Reader) (*LineItemReconReport, error) {
	lirReport := LineItemReconReport{}

	// TODO: Optimize report initialization by parsing lines and orders while we read the input.
	err := lirReport.readInput(reportCsv)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse input CSV: %v", err)
	}

	err = lirReport.parseRawLines()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse lines: %v", err)
	}

	return &lirReport, nil
}

// NewLineItem accepts an array of strings read in from the original CSV and returns an initiatlized LineItem.
// TODO: Move LineItem code to its own file
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

// GetRawLines returns the raw list of strings read in when creating the given LineItemReconReport.
// TODO: Remove this once LineItem parsing is implemented
func (l *LineItemReconReport) GetRawLines() [][]string {
	return l.rawLines
}

func (l *LineItemReconReport) readInput(reportCsv io.Reader) error {
	lines, err := csv.NewReader(reportCsv).ReadAll()
	if err != nil {
		return err
	}

	// Throw away the first line, which contains the field headers.
	l.rawLines = lines[1:]
	return nil
}

func (l *LineItemReconReport) parseRawLines() error {
	if l.Orders == nil {
		l.Orders = map[int]*Order{}
	}

	for _, rl := range l.rawLines {
		li, err := NewLineItem(rl)
		if err != nil {
			return fmt.Errorf("Failed to parse line [%s]: %v", rl, err)
		}

		if _, ok := l.Orders[li.OrderID]; !ok {
			l.Orders[li.OrderID] = &Order{}
		}
		l.Orders[li.OrderID].AddLineItem(li)
	}

	return nil
}

// AddLineItem appends a LineItem to the given Order.
func (o *Order) AddLineItem(li *LineItem) {
	o.LineItems = append(o.LineItems, li)
}
